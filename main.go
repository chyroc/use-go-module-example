package main

import (
	"fmt"
	packageV1 "github.com/chyroc/go-module-example"
	packageV2 "github.com/chyroc/go-module-example/v2"
)

func main() {
	fmt.Println(packageV1.FuncA())
	fmt.Println(packageV2.FuncB())
}
