run:
	echo "starting server locally"
	go run .
	cd webapp && yarn start

run-only-web:
	echo "starting webapp locally"
	cd webapp && yarn start

run-only-server:
	echo "starting server locally"
	go run .

build-grpc:
	echo "building employee connectors..."
	export PATH="$PATH:$(go env GOPATH)/bin"
	cd ./proto
	ls
	cd proto && protoc --go_out=../servers/services/grpc --go_opt=paths=source_relative --go-grpc_out=../servers/services/grpc --go-grpc_opt=paths=source_relative customer.proto


test:
	cd ..
	$(MAKE) -C ../Server main.go
	ls
	echo "ass"