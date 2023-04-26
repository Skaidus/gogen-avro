// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NestedUnionRecord struct {
	IntField int32 `json:"IntField"`
}

const NestedUnionRecordAvroCRC64Fingerprint = "#\xb6\xed\xa0\x87F=\xef"

func NewNestedUnionRecord() NestedUnionRecord {
	r := NestedUnionRecord{}
	return r
}

func DeserializeNestedUnionRecord(r io.Reader) (NestedUnionRecord, error) {
	t := NewNestedUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNestedUnionRecordFromSchema(r io.Reader, schema string) (NestedUnionRecord, error) {
	t := NewNestedUnionRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNestedUnionRecord(r NestedUnionRecord, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	return err
}

func (r NestedUnionRecord) Serialize(w io.Writer) error {
	return writeNestedUnionRecord(r, w)
}

func (r NestedUnionRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"}],\"name\":\"NestedUnionRecord\",\"type\":\"record\"}"
}

func (r NestedUnionRecord) SchemaName() string {
	return "NestedUnionRecord"
}

func (_ NestedUnionRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ NestedUnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedUnionRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.IntField}

		return w

	}
	panic("Unknown field index")
}

func (r *NestedUnionRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NestedUnionRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NestedUnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NestedUnionRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NestedUnionRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NestedUnionRecord) Finalize()                        {}

func (_ NestedUnionRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NestedUnionRecordAvroCRC64Fingerprint)
}

func (r NestedUnionRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NestedUnionRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IntField")
	}
	return nil
}
