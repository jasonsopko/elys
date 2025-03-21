package transferhook

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"

	"github.com/elys-network/elys/x/transferhook/keeper"
	"github.com/elys-network/elys/x/transferhook/types"

	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	keeper keeper.Keeper
	app    porttypes.IBCModule
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper, app porttypes.IBCModule) IBCModule {
	return IBCModule{
		keeper: k,
		app:    app,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return im.app.OnChanOpenInit(
		ctx,
		order,
		connectionHops,
		portID,
		channelID,
		channelCap,
		counterparty,
		version,
	)
}

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelId string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelId, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// doCustomLogic()
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// doCustomLogic()
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is successfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnRecvPacket (transferhook): Sequence: %d, Source: %s, %s; Destination: %s, %s",
		packet.Sequence, packet.SourcePort, packet.SourceChannel, packet.DestinationPort, packet.DestinationChannel))

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// ibc-go v5 has a Memo field that can store forwarding info
	// For older version of ibc-go, the data must be stored in the receiver field
	var metadata string
	if data.Memo != "" { // ibc-go v5+
		metadata = data.Memo
	} else { // ibc-go v5- (e.g. Osmosis)
		metadata = data.Receiver
	}

	// If a valid receiver address has been provided and no memo,
	// this is clearly just an normal IBC transfer
	// Pass down the stack immediately instead of parsing
	_, err := sdk.AccAddressFromBech32(data.Receiver)
	if err == nil && data.Memo == "" {
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	// parse out any forwarding info
	packetForwardMetadata, err := types.ParsePacketMetadata(metadata)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// If the parsed metadata is nil, that means there is no forwarding logic
	// Pass the packet down to the next middleware
	if packetForwardMetadata == nil {
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	// Modify the packet data by replacing the JSON metadata field with a receiver address
	// to allow the packet to continue down the stack
	data.Receiver = packetForwardMetadata.Receiver
	bz, err := transfertypes.ModuleCdc.MarshalJSON(&data)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}
	packet.Data = bz

	// Pass the new packet down the middleware stack first
	ack := im.app.OnRecvPacket(ctx, packet, relayer)
	if !ack.Success() {
		return ack
	}

	transferhookParams := im.keeper.GetParams(ctx)

	// If the transfer was successful, then route to the corresponding module, if applicable
	switch routingInfo := packetForwardMetadata.RoutingInfo.(type) {
	case types.AmmPacketMetadata:
		// If amm routing is inactive (but the packet had routing info in the memo) return an ack error
		if !transferhookParams.AmmActive {
			im.keeper.Logger(ctx).Error(fmt.Sprintf("Packet from %s had amm routing info but transferhook amm routing is disabled", data.Sender))
			return channeltypes.NewErrorAcknowledgement(types.ErrPacketForwardingInactive)
		}
		im.keeper.Logger(ctx).Info(fmt.Sprintf("Forwarding packet from %s to amm", data.Sender))

		// swap operation
		if err := im.keeper.Swap(ctx, packet, data, routingInfo); err != nil {
			im.keeper.Logger(ctx).Error(fmt.Sprintf("Error swap packet from transferhook for %s: %s", data.Sender, err.Error()))
			return channeltypes.NewErrorAcknowledgement(err)
		}

		return ack

	default:
		return channeltypes.NewErrorAcknowledgement(errorsmod.Wrapf(types.ErrUnsupportedTransferhookRoute, "%T", routingInfo))
	}
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("[IBC-TRANSFER] OnAcknowledgementPacket  %v", packet))
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger(ctx).Error(fmt.Sprintf("[IBC-TRANSFER] OnTimeoutPacket  %v", packet))
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// This is implemented by ICS4 and all middleware that are wrapping base application.
// The base application will call `sendPacket` or `writeAcknowledgement` of the middleware directly above them
// which will call the next middleware until it reaches the core IBC handler.
// SendPacket implements the ICS4 Wrapper interface
func (im IBCModule) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	return nil
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCModule) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return nil
}

// GetAppVersion returns the interchain accounts metadata.
func (im IBCModule) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return ibctransfertypes.Version, true // im.keeper.GetAppVersion(ctx, portID, channelID)
}
