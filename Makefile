dep:
	dep ensure -v
	
contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go

.PHONY: contracts dep
