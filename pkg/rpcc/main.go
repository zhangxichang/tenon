package rpcc

import (
	"fmt"
	"reflect"

	"github.com/gorilla/websocket"
)

type MethodCall struct {
	Id     uint   `json:"id"`
	Method string `json:"method"`
	Params []any  `json:"params"`
}

type MethodResult struct {
	Id  uint  `json:"id"`
	Ok  any   `json:"ok"`
	Err error `json:"err"`
}

type RPConnection struct {
	conn *websocket.Conn
	fns  map[string]reflect.Value
}

func New(conn *websocket.Conn) *RPConnection {
	return &RPConnection{
		conn: conn,
		fns:  make(map[string]reflect.Value),
	}
}
func (self *RPConnection) RegisterFn(method string, fn any) error {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	if fnValue.Kind() != reflect.Func {
		return fmt.Errorf("[%s] 必须是一个函数", method)
	}
	numOut := fnType.NumOut()
	if numOut > 2 {
		return fmt.Errorf("[%s] 返回值数量最多2个, 实际有%v个", method, numOut)
	}
	if numOut == 2 {
		if fnType.Out(1) != reflect.TypeFor[error]() {
			return fmt.Errorf("[%s] 两个返回值时, 第二个必须是error", method)
		}
	}
	self.fns[method] = fnValue
	return nil
}
func (self *RPConnection) Call(req *MethodCall) (*MethodResult, error) {
	fnValue, exist := self.fns[req.Method]
	if !exist {
		return nil, fmt.Errorf("没有找到[%v]方法, 调用ID: [%v]", req.Method, req.Id)
	}
	fnType := fnValue.Type()
	numIn := fnType.NumIn()
	numParams := len(req.Params)
	if numParams != numIn {
		return nil, fmt.Errorf("[%s] 参数数量不匹配: 期望 %d, 实际 %d", req.Method, numIn, numParams)
	}
	args := make([]reflect.Value, numIn)
	for i := range numIn {
		paramType := fnType.In(i)
		argValue := reflect.ValueOf(req.Params[i])
		if !argValue.Type().ConvertibleTo(paramType) {
			return nil, fmt.Errorf("[%s] 参数 %d 类型不匹配: 期望 %v, 实际 %v", req.Method, i, paramType, argValue.Type())
		}
		args[i] = argValue.Convert(paramType)
	}
	results := fnValue.Call(args)
	var ok any = nil
	var err error = nil
	if len(results) == 1 {
		val := results[0].Interface()
		if e, isErr := val.(error); isErr {
			err = e
		} else {
			ok = val
		}
	} else if len(results) == 2 {
		ok = results[0].Interface()
		if errVal := results[1].Interface(); errVal != nil {
			err = errVal.(error)
		}
	}
	return &MethodResult{Id: req.Id, Ok: ok, Err: err}, nil
}
func (self *RPConnection) Serve() error {
	for {
		req := &MethodCall{}
		if err := self.conn.ReadJSON(req); err != nil {
			return fmt.Errorf("读取方法调用请求错误: %v", err)
		}
		go func() {
			result, err := self.Call(req)
			if err != nil {
				fmt.Printf("调用方法错误: %v\n", err)
				return
			}
			self.conn.WriteJSON(result)
		}()
	}
}
