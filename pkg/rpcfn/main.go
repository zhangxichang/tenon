package rpcfn

import (
	"fmt"
	"tenon/pkg/rpcc"
)

func RPCFunction(rpcc *rpcc.RPConnection) {
	rpcc.RegisterFn("hello", func() {
		fmt.Println("你好，世界")
	})
}
