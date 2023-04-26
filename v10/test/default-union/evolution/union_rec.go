// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Skaidus/gogen-avro/v10/compiler"
	"github.com/Skaidus/gogen-avro/v10/vm"
	"github.com/Skaidus/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type UnionRec struct {
	A int32 `json:"a"`
}

const UnionRecAvroCRC64Fingerprint = "1\xf9\xae\xb7W\x80#\xf9"

func NewUnionRec() UnionRec {
	r := UnionRec{}
	return r
}

func DeserializeUnionRec(r io.Reader) (UnionRec, error) {
	t := NewUnionRec()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUnionRecFromSchema(r io.Reader, schema string) (UnionRec, error) {
	t := NewUnionRec()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUnionRec(r UnionRec, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.A, w)
	if err != nil {
		return err
	}
	return err
}

func (r UnionRec) Serialize(w io.Writer) error {
	return writeUnionRec(r, w)
}

func (r UnionRec) Schema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":\"int\"}],\"name\":\"unionRec\",\"type\":\"record\"}"
}

func (r UnionRec) SchemaName() string {
	return "unionRec"
}

func (_ UnionRec) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ UnionRec) SetInt(v int32)       { panic("Unsupported operation") }
func (_ UnionRec) SetLong(v int64)      { panic("Unsupported operation") }
func (_ UnionRec) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ UnionRec) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ UnionRec) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ UnionRec) SetString(v string)   { panic("Unsupported operation") }
func (_ UnionRec) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRec) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.A}

		return w

	}
	panic("Unknown field index")
}

func (r *UnionRec) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *UnionRec) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ UnionRec) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionRec) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionRec) HintSize(int)                     { panic("Unsupported operation") }
func (_ UnionRec) Finalize()                        {}

func (_ UnionRec) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecAvroCRC64Fingerprint)
}

func (r UnionRec) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["a"], err = json.Marshal(r.A)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *UnionRec) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["a"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.A); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for a")
	}
	return nil
}
