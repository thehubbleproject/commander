dep:
	dep ensure -v
	rm -rf vendor/github.com/ethereum
	mkdir -p vendor/github.com/ethereum
	git clone -b v1.9.0 --single-branch --depth 1 https://github.com/ethereum/go-ethereum vendor/github.com/ethereum/go-ethereum
	
contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go
	abigen --abi=contracts/merkleTree/merkleTree.abi --pkg=merkleTree --out=contracts/merkleTree/merkleTree.go
		abigen --abi=contracts/trial/trial.abi --pkg=trial --out=contracts/trial/trial.go


clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/bopr ./cmd

init:
	./build/bopr init

reset:
	

start:
	./build/bopr start
	 
.PHONY: contracts dep build clean start
