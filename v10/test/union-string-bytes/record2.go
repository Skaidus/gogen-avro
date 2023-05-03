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

	"github.com/Skaidus/gogen-avro/v10/compiler"
	"github.com/Skaidus/gogen-avro/v10/vm"
	"github.com/Skaidus/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Record2 struct {
	Intfield int32 `json:"intfield"`
}

const Record2AvroCRC64Fingerprint = "S\n_\xab\xc7-\xf9\x9c"

func NewRecord2() Record2 {
	r := Record2{}
	return r
}

func DeserializeRecord2(r io.Reader) (Record2, error) {
	t := NewRecord2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecord2FromSchema(r io.Reader, schema string) (Record2, error) {
	t := NewRecord2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecord2(r Record2, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Intfield, w)
	if err != nil {
		return err
	}
	return err
}

func (r Record2) Serialize(w io.Writer) error {
	return writeRecord2(r, w)
}

func (r Record2) Schema() string {
	return "{\"fields\":[{\"name\":\"intfield\",\"type\":\"int\"}],\"name\":\"record2\",\"type\":\"record\"}"
}

func (r Record2) SchemaName() string {
	return "record2"
}

func (_ Record2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Record2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Record2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Record2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Record2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Record2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Record2) SetString(v string)   { panic("Unsupported operation") }
func (_ Record2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record2) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Intfield}

		return w

	}
	panic("Unknown field index")
}

func (r *Record2) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Record2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Record2) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Record2) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Record2) HintSize(int)                     { panic("Unsupported operation") }
func (_ Record2) Finalize()                        {}

func (_ Record2) AvroCRC64Fingerprint() []byte {
	return []byte(Record2AvroCRC64Fingerprint)
}

func (r Record2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["intfield"], err = json.Marshal(r.Intfield)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Record2) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["intfield"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Intfield); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for intfield")
	}
	return nil
}
