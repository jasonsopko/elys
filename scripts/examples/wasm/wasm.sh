elysd tx wasm store cw20_base.wasm --from=treasury --keyring-backend=test --chain-id=elystestnet-1 --gas=auto --gas-adjustment=1.3 -y
elysd tx wasm instantiate 1 '{"name":"Cw20","symbol":"cwww", "decimals":6, "initial_balances":[]}' --from=treasury --label "Cw20" --chain-id=elystestnet-1 --gas=auto --gas-adjustment=1.3 -b=sync --keyring-backend=test --no-admin -y