all: build

init-dev: init-chain init-validator

start-dev:
	go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go start --home ./build/.auth-chaind

init-chain:
	go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go init --chain-id=auth auth --home ./build/.auth-chaind
	echo "y" | go run cmd/auth-chaincli/main.go keys add auth-chain1 --home ./build/.auth-chaind

init-validator:
	go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go add-genesis-account $(shell go run cmd/auth-chaincli/main.go keys show auth-chain1 -a --home ./build/.auth-chaind) 1000000000stake --home ./build/.auth-chaind
	go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go gentx --name auth-chain1 --home ./build/.auth-chaind --moniker authchain --website test.com --identity test --security-contact test@test.com --details atest
	go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go collect-gentxs --home ./build/.auth-chaind

install-debug: go.sum
	go build -mod=readonly -gcflags="all=-N -l" ./cmd/auth-chaind
	go build -mod=readonly -gcflags="all=-N -l" ./cmd/auth-chaincli

build:
	@mkdir -p build/
	@go build -mod=mod -o build/clayd ./cmd/clayd
	@go build -mod=mod -o build/claycli ./cmd/claycli

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

localnet-consensus:
	./scripts/add-validators-to-validator-set.sh

localnet-start: init-dev export-key
	NODE0ADDRESS=$(shell go run cmd/auth-chaind/main.go cmd/auth-chaind/genaccounts.go tendermint show-node-id --home ./build/.auth-chaind)@192.16.10.2:26656 docker-compose up

export-key:
	echo "password1234\npassword1234" | go run cmd/auth-chaincli/main.go keys export auth-chain1 2> ./build/validator

clean:
	sudo rm -r ./build
	docker-compose down
