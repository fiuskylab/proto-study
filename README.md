# proto-study
Step-by-step on how to implement gRPC in a simple API.

## Proposition
- A simple user CRUD
- The database will be a JSON file

## References
- [gRPC.io](https://grpc.io/docs/languages/go/)
- [Protocol Buffers - Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [Protocol Buffers Guide - Proto3](https://developers.google.com/protocol-buffers/docs/proto3)

## Step-by-step

### 1. Initiating Go mod
- Go modules will allow us to easily add third-party/external packages
- > $ go mod init <your_url>
- In my case, the URL that I'll be using is the same of this repository so the command will be something like:
- > $ go mod init github.com/fiuskylab/proto-study

### 2. Creating Protos
- Now is needed to write out `proto` file
- So create a `protos/user.proto`
```
// Syntax version
syntax = "proto3";

// Package name
package userproto;

// Message can be interpreted as a struct
// An aggregate of fields with types
message UserRequest {
  // User identifier
  string uuid = 0;

  // Name field
  string name = 1;

  uint32 age = 2;
  // Age

  // User roles
  enum Role {
    COMMON = 0;
    AGENT  = 1;
    MOD    = 2;
    ADMIN  = 3;
  }

  // Role
  Role role = 3;

  // Email
  string email = 4;

  // Password
  string password = 5;
}

// The service will the struct that store all method of this proto
// This service is a simple one, but you can increase amount of methods and messages
// So, for example, if you create a service FindByID, you can create a message UserRequestID that will only have the field 'string uuid = 0;'
service CRUD {
  rpc CreateUser (UserRequest) returns (UserRequest) {}
}
```
