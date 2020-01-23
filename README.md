# Extensions Bootstrap

This project constitutes the foundation of an ocis extension.

# Folder Structure

```console
.
├── README.md
├── cmd
├── go.mod
└── pkg
    ├── proto
    ├── server
    └── service
```

# Creating your own extension 🤖

1. define your service on a `.proto` file (see pkg/proto/yeller.proto and the resources section for further reference)
2. run protobuf compiler: `protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. pkg/proto/yeller.proto`
3. implement go-micro interfaces. a.k.a add your business logic
4. now you've implemented some of your business logic let's hook it up as a go-micro service and add it to the registry®!
5. finally create your command

# Resources

- [protobuf language syntax](https://developers.google.com/protocol-buffers/docs/proto3)
- [as a cli app we use urfave/cli@v1](https://github.com/urfave/cli/blob/master/docs/v1/manual.md)
- [on go signal channels](https://gobyexample.com/signals)