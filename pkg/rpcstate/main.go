package rpcstate

import (
	"tenon/pkg/rpcc"
)

type RPCState struct {
	str string
}

func New() *RPCState {
	return &RPCState{
		str: "echo: ",
	}
}
func (self *RPCState) Register(rpcc *rpcc.RPConnection) {
	rpcc.RegisterFn("echo", func(str string) (*string, error) {
		echoStr := self.str + str
		return &echoStr, nil
	})
}
