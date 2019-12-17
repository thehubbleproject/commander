dep:
	dep ensure -v
	rm -rf vendor/github.com/ethereum
	mkdir -p vendor/github.com/ethereum
	git clone -b v1.9.0 --single-branch --depth 1 https://github.com/ethereum/go-ethereum vendor/github.com/ethereum/go-ethereum
	
contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go

clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/bopr cmd/main.go

init:
	./build/bopr init

reset:
	

start:
	./build/bopr start
	 
.PHONY: contracts dep build clean start
