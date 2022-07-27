package main

import (
	"petshop-grpc-go/pkg/cli"
)

//go:generate make grpc-gen

func main() {
	cli.Run()
}
