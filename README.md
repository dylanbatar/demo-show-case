### Protoc
```
protoc -Igrpcs/proto --go_out=. --go-grpc_out=. --go_opt=module=github.com/dylanbatar/demo-show-case --go-grpc_opt=module=github.com/dylanbatar/demo-show-case grpcs/proto/person.proto
```

