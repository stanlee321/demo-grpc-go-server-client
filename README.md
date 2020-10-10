# GRPC Demo server


## Setup

If windows:

```
choco install protoc --pre

protoc profiles.proto --go_out=plugins=grpc:.
protoc profiles.proto --go_out=plugins=grpc:.


go mod init github.com/stanlee321/demo-grpc-go-server-client/server/ProfileService

go mod tidy
go mod vendor

```

