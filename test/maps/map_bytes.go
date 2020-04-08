// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
package avro

import (
	"fmt"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

func writeMapBytes(r map[string][]byte, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteBytes(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapBytesWrapper struct {
	Target *map[string][]byte
	keys   []string
	values [][]byte
}

func (_ *MapBytesWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapBytesWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapBytesWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (_ *MapBytesWrapper) NullField(i int) { panic("Unsupported operation") }

func (r *MapBytesWrapper) Finalize() {
	fmt.Printf("Finalizing!\n")
	for i := range r.keys {
		fmt.Printf("%v\n", r.keys[i])
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapBytesWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v []byte
	r.values = append(r.values, v)
	return &types.Bytes{Target: &r.values[len(r.values)-1]}
}

func (_ *MapBytesWrapper) AppendArray() types.Field { panic("Unsupported operation") }
