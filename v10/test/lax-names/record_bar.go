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

type RecordBar struct {
	Id string `json:"id"`
}

const RecordBarAvroCRC64Fingerprint = "0\x99\x1e\"\x92\xe1\x82{"

func NewRecordBar() RecordBar {
	r := RecordBar{}
	return r
}

func DeserializeRecordBar(r io.Reader) (RecordBar, error) {
	t := NewRecordBar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecordBarFromSchema(r io.Reader, schema string) (RecordBar, error) {
	t := NewRecordBar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecordBar(r RecordBar, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	return err
}

func (r RecordBar) Serialize(w io.Writer) error {
	return writeRecordBar(r, w)
}

func (r RecordBar) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RecordBar\",\"type\":\"record\"}"
}

func (r RecordBar) SchemaName() string {
	return "RecordBar"
}

func (_ RecordBar) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RecordBar) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RecordBar) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RecordBar) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RecordBar) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RecordBar) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RecordBar) SetString(v string)   { panic("Unsupported operation") }
func (_ RecordBar) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecordBar) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	}
	panic("Unknown field index")
}

func (r *RecordBar) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RecordBar) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RecordBar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RecordBar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RecordBar) HintSize(int)                     { panic("Unsupported operation") }
func (_ RecordBar) Finalize()                        {}

func (_ RecordBar) AvroCRC64Fingerprint() []byte {
	return []byte(RecordBarAvroCRC64Fingerprint)
}

func (r RecordBar) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RecordBar) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	return nil
}
