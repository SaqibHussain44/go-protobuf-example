
You need the following library to run:

- `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
- `go get -u google.golang.org/grpc`
- `protoc --go_out=plugins=grpc:. *.proto`

protoc -I greeter/ greeter/greeter.proto --go_out=plugins=grpc:greeter

```bash
Issue encountered:
protoc -I greeter/ greeter/greeter.proto --go_out=plugins=grpc:greeter
protoc-gen-go: program not found or is not executable
--go_out: protoc-gen-go: Plugin failed with status code 1.
```

protoc-gen-go needs to be in your shell path, i.e. one of the directories listed in the PATH environment variable, which is different from the Go path. You can test this by simply typing protoc-gen-go at the command line: If it says "command not found" (or similar) then it's not in your PATH.

```bash
$ export PATH=$PATH:$GOPATH/bin
```

Run the server first:

```bash
$ go run greeter_server/main.go
```

The run the client:

```bash
$ go run greeter_client/main.go
```
