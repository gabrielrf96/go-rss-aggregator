CMD_SERVER=rssagg-server
CMD_DEVCLI=rssagg-cli
ALIAS_DEVCLI=cli

-include .env

init:
	cp -n .env.example .env
	go mod download
	make all

all:
	make devcli
	make build

devcli:
	go build -o ./bin/ ./cmd/${CMD_DEVCLI}
	ln -sf ./bin/${CMD_DEVCLI} ./${ALIAS_DEVCLI}

build:
	go build -o ./bin/ ./cmd/${CMD_SERVER}

run:
	make build
	./bin/${CMD_SERVER}
