# https://github.com/pjchender/besg-grpc
go := /usr/local/go/bin/go
go_root := `$(go) env GOROOT`

init:
	${go} get github.com/golang/protobuf/protoc-gen-go
	${go} get github.com/golang/protobuf/protoc-gen-go
	cd proto/calculator ; protoc calculator.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
	sudo cp -n proto/calculator ${go_root}/src/proto -r
	${go} mod init simple
	${go} mod tidy

run:
	${go} run server.go &
	sleep 1
	${go} run client.go 
	- killall go

all:
