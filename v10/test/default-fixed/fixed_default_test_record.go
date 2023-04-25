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

type FixedDefaultTestRecord struct {
	FixedField TestFixedDefaultType `json:"FixedField"`
}

const FixedDefaultTestRecordAvroCRC64Fingerprint = "\xac\xb5kw\x1eWK\xcf"

func NewFixedDefaultTestRecord() FixedDefaultTestRecord {
	r := FixedDefaultTestRecord{}
	copy(r.FixedField[:], []byte("\x00\x01\x12\x00\x13C\x00\x01\x12\x00\x13S"))
	return r
}

func DeserializeFixedDefaultTestRecord(r io.Reader) (FixedDefaultTestRecord, error) {
	t := NewFixedDefaultTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFixedDefaultTestRecordFromSchema(r io.Reader, schema string) (FixedDefaultTestRecord, error) {
	t := NewFixedDefaultTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFixedDefaultTestRecord(r FixedDefaultTestRecord, w io.Writer) error {
	var err error
	err = writeTestFixedDefaultType(r.FixedField, w)
	if err != nil {
		return err
	}
	return err
}

func (r FixedDefaultTestRecord) Serialize(w io.Writer) error {
	return writeFixedDefaultTestRecord(r, w)
}

func (r FixedDefaultTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":\"\\u0000\\u0001\\u0012\\u0000\\u0013C\\u0000\\u0001\\u0012\\u0000\\u0013S\",\"name\":\"FixedField\",\"type\":{\"name\":\"TestFixedDefaultType\",\"size\":12,\"type\":\"fixed\"}}],\"name\":\"FixedDefaultTestRecord\",\"type\":\"record\"}"
}

func (r FixedDefaultTestRecord) SchemaName() string {
	return "FixedDefaultTestRecord"
}

func (_ FixedDefaultTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FixedDefaultTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := TestFixedDefaultTypeWrapper{Target: &r.FixedField}

		return w

	}
	panic("Unknown field index")
}

func (r *FixedDefaultTestRecord) SetDefault(i int) {
	switch i {
	case 0:
		copy(r.FixedField[:], []byte("\x00\x01\x12\x00\x13C\x00\x01\x12\x00\x13S"))
		return
	}
	panic("Unknown field index")
}

func (r *FixedDefaultTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FixedDefaultTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ FixedDefaultTestRecord) Finalize()                        {}

func (_ FixedDefaultTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(FixedDefaultTestRecordAvroCRC64Fingerprint)
}

func (r FixedDefaultTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["FixedField"], err = json.Marshal(r.FixedField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FixedDefaultTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["FixedField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FixedField); err != nil {
			return err
		}
	} else {
		copy(r.FixedField[:], []byte("\x00\x01\x12\x00\x13C\x00\x01\x12\x00\x13S"))
	}
	return nil
}
