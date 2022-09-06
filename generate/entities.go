package generate

import (
	"reflect"
)

var protoMapTypes = map[string]reflect.Kind{
	"string":   reflect.String,
	"*string":  reflect.String,
	"int":      reflect.Int32,
	"*int":     reflect.Int32,
	"int8":     reflect.Int32,
	"*int8":    reflect.Int32,
	"uint8":    reflect.Int32,
	"*uint8":   reflect.Int32,
	"int16":    reflect.Int32,
	"*int16":   reflect.Int32,
	"uint16":   reflect.Int32,
	"*uint16":  reflect.Int32,
	"int32":    reflect.Int32,
	"*int32":   reflect.Int32,
	"uint32":   reflect.Uint32,
	"*uint32":  reflect.Uint32,
	"uint64":   reflect.Uint64,
	"*uint64":  reflect.Uint64,
	"int64":    reflect.Int64,
	"*int64":   reflect.Int64,
	"[]string": reflect.Slice,
	"[]int":    reflect.Slice,
	"[]int64":  reflect.Slice,
	"[]int32":  reflect.Slice,
	"[]uint32": reflect.Slice,
	"[]uint64": reflect.Slice,
	"bool":     reflect.Bool,
	"*bool":    reflect.Bool,
	"struct":   reflect.Struct,
	"*struct":  reflect.Struct,
	"float32":  reflect.Float32,
	"*float32": reflect.Float32,
	"float64":  reflect.Float64,
	"*float64": reflect.Float64,
}

type protoMessage struct {
	Name    string          `json:"name"`
	Comment string          `json:"comment"`
	Item    []messageMember `json:"item"`
}

type messageMember struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type protoService struct {
	Name    string          `json:"name"`
	Comment string          `json:"comment"`
	Item    []serviceMember `json:"item"`
}

type serviceMember struct {
	FuncName     string `json:"func_name"`
	RequestName  string `json:"request_name"`
	ResponseName string `json:"response_name"`
	Comment      string `json:"comment"`
}
