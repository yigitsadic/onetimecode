# Onetimecodego

This project, designed for mapping random strings to given identifiers. For example with given identifier `c96e40ed-1a06-427b-9c2d-a85ebea7d9e8`
is mapped to random **x** length of string (*HPI05VM*) for **y** seconds. 

Application does not handle any authentication nor authorization and it's designed for single instance use only.

Generated codes stores in memory, because of that, any time service is down, your codes will go down.

## Configuration

Application configurable through environment variables.

```
CODE_EXP = 60
```

## Grpc

Example client can be found in **client.go**

```protobuf
syntax = "proto3";
package otcgo;

option go_package = "grpc;grpc";

message OneTimeCodeGen {
  string identifier = 1;
}

message ReadCodeReq {
  string value = 1;
}

message OneTimeCodeResponse {
  string identifier = 1;
  int64 expiresAt = 2;
  string value = 3;
}

service OneTimeCodeService {
  rpc CreateCode(OneTimeCodeGen) returns (OneTimeCodeResponse) {}
  rpc ReadCode(ReadCodeReq) returns (OneTimeCodeResponse) {}
}
```

## Docker

You can run ` docker run --rm -e CODE_EXP=5 -p 9000:9000 yigitsadic/onetimecodego:latest`

For client: `go run client.go`
