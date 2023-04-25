// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"io"

	"github.com/actgardner/gogen-avro/v10/util"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeTestFixedType(r TestFixedType, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type TestFixedTypeWrapper struct {
	Target *TestFixedType
}

type TestFixedType [12]byte

func (b *TestFixedType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b TestFixedType) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

func (_ TestFixedTypeWrapper) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) SetInt(v int32)      { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) SetLong(v int64)     { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r TestFixedTypeWrapper) SetBytes(v []byte) {
	copy((*r.Target)[:], v)
}
func (_ TestFixedTypeWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) NullField(int)                    { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) HintSize(int)                     { panic("Unsupported operation") }
func (_ TestFixedTypeWrapper) Finalize()                        {}
func (_ TestFixedTypeWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
