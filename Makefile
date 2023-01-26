BINARY=gatus

all:
	go build

# Because there's a folder called "test", we need to make the target "test" phony
.PHONY: test

install:
	go build -mod vendor -o $(BINARY) .

run:
	GATUS_CONFIG_PATH=./config.yaml ./$(BINARY)

clean:
	rm $(BINARY)

test:
	go test ./... -cover


##########
# Docker #
##########

docker-build:
	docker build -t twinproduction/gatus:latest .

docker-run:
	docker run -p 8080:8080 --name gatus twinproduction/gatus:latest

docker-build-and-run: docker-build docker-run


#############
# Front end #
#############

frontend-build:
	npm --prefix web/app run build

frontend-run:
	npm --prefix web/app run serve


#--------------------------------------------------------------------------------------------------------------------


linux:
	go vet
	GOOS=linux GOARCH=amd64 go build -o gatus_linux

deploy_tcs: linux
	scp gatus_linux philip@tcs.com:/home/philip/tmp
	tar -czf deploy-gatus.tar.gz -C web/app .
	scp deploy-gatus.tar.gz philip@tcs.com:/home/philip/tmp

deploy: linux frontend-build
	echo scp gatus_linux pschlump@nfc-auth.com:/tmp
	echo tar -czf deploy-gatus.tar.gz -C web/static .
	echo scp deploy-gatus.tar.gz pschlump@nfc-auth.com:/tmp
	scp ./config/config.yaml  pschlump@nfc-auth.com:/tmp

