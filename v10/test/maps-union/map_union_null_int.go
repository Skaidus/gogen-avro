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

func writeMapUnionNullInt(r map[string]*UnionNullInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeUnionNullInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapUnionNullIntWrapper struct {
	Target *map[string]*UnionNullInt
	keys   []string
	values []*UnionNullInt
}

func (_ *MapUnionNullIntWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapUnionNullIntWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]*UnionNullInt, 0, s)
	}
}

func (r *MapUnionNullIntWrapper) NullField(_ int) {
	r.values[len(r.values)-1] = nil
}

func (r *MapUnionNullIntWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapUnionNullIntWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v *UnionNullInt
	v = NewUnionNullInt()

	r.values = append(r.values, v)
	return r.values[len(r.values)-1]
}

func (_ *MapUnionNullIntWrapper) AppendArray() types.Field { panic("Unsupported operation") }
