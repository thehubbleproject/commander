dep:
	dep ensure -v
	
contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go

clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/bopr cmd/main.go

.PHONY: contracts dep build clean
