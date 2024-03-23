LOCAL_BIN:=$(CURDIR)/bin

install-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install golang.org/x/vuln/cmd/govulncheck@latest

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

generate:
	mkdir -p pkg/swagger
	make generate-user-api
	$(LOCAL_BIN)/statik -f -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'
	make generate-auth-api
	make generate-access-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--go_out=pkg/user_v1 \
	--go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/user_v1/user.proto 

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 \
	--go_out=pkg/auth_v1 \
	--go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 \
	--go_out=pkg/access_v1 \
	--go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/access_v1 \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/access_v1/access.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

generate-mocks:
	GOBIN=$(LOCAL_BIN) go generate ./...

test:
	go clean -testcache
	go test ./... -covermode count  -count 5

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count  -count 5
	grep -v 'mock\|_test' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";

gen-ca-cert:
	openssl ecparam -name prime256v1 -genkey -out test/certs/ca.key
	openssl req -new -key test/certs/ca.key -out test/certs/ca.csr -subj "/C=RU/ST=SPb/O=BestCAever"
	openssl x509 -req -sha256 -days 1825 -in test/certs/ca.csr -signkey test/certs/ca.key -out test/certs/ca.crt

gen-service-cert:
	openssl ecparam -name prime256v1 -genkey -out test/certs/service.key
	openssl req -new -key test/certs/service.key -out test/certs/service.csr -config config/certificate.conf
	openssl x509 -req -sha256 -days 365 -in test/certs/service.csr -CA test/certs/ca.crt -CAkey test/certs/ca.key -CAcreateserial -out test/certs/service.crt -extfile config/certificate.conf -extensions req_ext

gen-certs: gen-ca-cert gen-service-cert

gnz-install:
	go install github.com/bojand/ghz/cmd/ghz@latest

grpc-load-test:
	ghz \
		--call user_v1.UserV1.Get \
		--data '{"id": 1}' \
		--rps 100 \
		--total 3000 \
		--skipTLS \
		localhost:50051

grpc-error-load-test:
	ghz \
		--call user_v1.UserV1.Get \
		--data '{"id": 110}' \
		--rps 100 \
		--total 3000 \
		--skipTLS \
		localhost:50051