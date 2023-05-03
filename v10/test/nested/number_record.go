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

	"github.com/masmovil/gogen-avro/v10/compiler"
	"github.com/masmovil/gogen-avro/v10/vm"
	"github.com/masmovil/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NumberRecord struct {
	IntField int32 `json:"IntField"`

	LongField int64 `json:"LongField"`

	FloatField float32 `json:"FloatField"`

	DoubleField float64 `json:"DoubleField"`
}

const NumberRecordAvroCRC64Fingerprint = "\xf4Zu\xd5Nt'~"

func NewNumberRecord() NumberRecord {
	r := NumberRecord{}
	return r
}

func DeserializeNumberRecord(r io.Reader) (NumberRecord, error) {
	t := NewNumberRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNumberRecordFromSchema(r io.Reader, schema string) (NumberRecord, error) {
	t := NewNumberRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNumberRecord(r NumberRecord, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	return err
}

func (r NumberRecord) Serialize(w io.Writer) error {
	return writeNumberRecord(r, w)
}

func (r NumberRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"}],\"name\":\"NumberRecord\",\"type\":\"record\"}"
}

func (r NumberRecord) SchemaName() string {
	return "NumberRecord"
}

func (_ NumberRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NumberRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NumberRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NumberRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NumberRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NumberRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NumberRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ NumberRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NumberRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.IntField}

		return w

	case 1:
		w := types.Long{Target: &r.LongField}

		return w

	case 2:
		w := types.Float{Target: &r.FloatField}

		return w

	case 3:
		w := types.Double{Target: &r.DoubleField}

		return w

	}
	panic("Unknown field index")
}

func (r *NumberRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NumberRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NumberRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NumberRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NumberRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NumberRecord) Finalize()                        {}

func (_ NumberRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NumberRecordAvroCRC64Fingerprint)
}

func (r NumberRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	output["LongField"], err = json.Marshal(r.LongField)
	if err != nil {
		return nil, err
	}
	output["FloatField"], err = json.Marshal(r.FloatField)
	if err != nil {
		return nil, err
	}
	output["DoubleField"], err = json.Marshal(r.DoubleField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NumberRecord) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["LongField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LongField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LongField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FloatField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FloatField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FloatField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DoubleField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DoubleField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DoubleField")
	}
	return nil
}
