# Go version of gRPC server for [Remote Piano](https://github.com/kaboc/flutter_remote_piano)

## Usage

```sh
$ go run ./server -p {port number}
```

Port 50051 is used as the default if `-p` option is omitted.

## Generating gRPC code

Run the command below at the root of this project.

```sh
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/piano.proto
```

## Links

* [Remote Piano](https://github.com/kaboc/flutter_remote_piano)
* [Piano Server written in Dart](https://github.com/kaboc/piano_server)
