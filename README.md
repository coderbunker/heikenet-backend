# heikenet-backend
Welcome..

### compile the contracts:

abigen -sol contracts/erc20/ERC20.sol -pkg erc20_contract -out contracts/erc20/ERC20.go
abigen -sol contracts/retainer/retainerHeike.sol -pkg retainer_contract -out contracts/retainer/retainerHeike.go

### heroku

heroku config
heroku config:set SECRET=secret
heroku pg:psql
heroku logs --tail
