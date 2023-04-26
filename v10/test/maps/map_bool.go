// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapBool(r map[string]bool, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteBool(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapBoolWrapper struct {
	Target *map[string]bool
	keys   []string
	values []bool
}

func (_ *MapBoolWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapBoolWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapBoolWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapBoolWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]bool, 0, s)
	}
}

func (r *MapBoolWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapBoolWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapBoolWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v bool
	r.values = append(r.values, v)
	return &types.Boolean{Target: &r.values[len(r.values)-1]}
}

func (_ *MapBoolWrapper) AppendArray() types.Field { panic("Unsupported operation") }
