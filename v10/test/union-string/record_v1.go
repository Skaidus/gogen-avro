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

type Record_v1 struct {
	Foo *UnionNullString `json:"foo"`

	Bar *UnionStringNull `json:"bar"`
}

const Record_v1AvroCRC64Fingerprint = "\x85|\xeb\x7f9\xf7J\x15"

func NewRecord_v1() Record_v1 {
	r := Record_v1{}
	r.Bar = NewUnionStringNull()

	return r
}

func DeserializeRecord_v1(r io.Reader) (Record_v1, error) {
	t := NewRecord_v1()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecord_v1FromSchema(r io.Reader, schema string) (Record_v1, error) {
	t := NewRecord_v1()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecord_v1(r Record_v1, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Foo, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Bar, w)
	if err != nil {
		return err
	}
	return err
}

func (r Record_v1) Serialize(w io.Writer) error {
	return writeRecord_v1(r, w)
}

func (r Record_v1) Schema() string {
	return "{\"fields\":[{\"name\":\"foo\",\"type\":[\"null\",\"string\"]},{\"name\":\"bar\",\"type\":[\"string\",\"null\"]}],\"name\":\"record_v1\",\"type\":\"record\"}"
}

func (r Record_v1) SchemaName() string {
	return "record_v1"
}

func (_ Record_v1) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Record_v1) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Record_v1) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Record_v1) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Record_v1) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Record_v1) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Record_v1) SetString(v string)   { panic("Unsupported operation") }
func (_ Record_v1) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record_v1) Get(i int) types.Field {
	switch i {
	case 0:
		r.Foo = NewUnionNullString()

		return r.Foo
	case 1:
		r.Bar = NewUnionStringNull()

		return r.Bar
	}
	panic("Unknown field index")
}

func (r *Record_v1) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Record_v1) NullField(i int) {
	switch i {
	case 0:
		r.Foo = nil
		return
	case 1:
		r.Bar = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Record_v1) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Record_v1) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Record_v1) HintSize(int)                     { panic("Unsupported operation") }
func (_ Record_v1) Finalize()                        {}

func (_ Record_v1) AvroCRC64Fingerprint() []byte {
	return []byte(Record_v1AvroCRC64Fingerprint)
}

func (r Record_v1) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["foo"], err = json.Marshal(r.Foo)
	if err != nil {
		return nil, err
	}
	output["bar"], err = json.Marshal(r.Bar)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Record_v1) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["foo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Foo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for foo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["bar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bar); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for bar")
	}
	return nil
}
