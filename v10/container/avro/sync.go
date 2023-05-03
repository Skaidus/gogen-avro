// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */
package avro

import (
	"encoding/json"
	"io"

	"github.com/masmovil/gogen-avro/v10/util"
	"github.com/masmovil/gogen-avro/v10/vm/types"
)

func writeSync(r Sync, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type SyncWrapper struct {
	Target *Sync
}

type Sync [16]byte

func (b *Sync) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b Sync) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

func (_ SyncWrapper) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ SyncWrapper) SetInt(v int32)      { panic("Unsupported operation") }
func (_ SyncWrapper) SetLong(v int64)     { panic("Unsupported operation") }
func (_ SyncWrapper) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ SyncWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r SyncWrapper) SetBytes(v []byte) {
	copy((*r.Target)[:], v)
}
func (_ SyncWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ SyncWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ SyncWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ SyncWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SyncWrapper) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SyncWrapper) NullField(int)                    { panic("Unsupported operation") }
func (_ SyncWrapper) HintSize(int)                     { panic("Unsupported operation") }
func (_ SyncWrapper) Finalize()                        {}
func (_ SyncWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
