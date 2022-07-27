# Petshop gRPC

This is a fun little project that implements the petshop api example via gRPC/Go.

Currently, only the UserService is implemented.

The gRPC server transport security is disabled for this demo.

# Building

Just run `make`, which will build the static binary and put it in the bin/ directory.

# Usage

```
Usage of bin/server:
  -address string
        The address to bind the gRPC Server to. (default "0.0.0.0")
  -log-level int
        The log level to set. (-1 .. 5) (default 1)
  -port uint
        The port to bind the gRPC Server to. (default 50051)
```