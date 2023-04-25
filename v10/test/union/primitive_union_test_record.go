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

type PrimitiveUnionTestRecord struct {
	UnionField *UnionNullStringLongIntFloatDoubleBool `json:"UnionField"`
}

const PrimitiveUnionTestRecordAvroCRC64Fingerprint = "k\xbb\x01A\xbdb\b\xd4"

func NewPrimitiveUnionTestRecord() PrimitiveUnionTestRecord {
	r := PrimitiveUnionTestRecord{}
	r.UnionField = nil
	return r
}

func DeserializePrimitiveUnionTestRecord(r io.Reader) (PrimitiveUnionTestRecord, error) {
	t := NewPrimitiveUnionTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePrimitiveUnionTestRecordFromSchema(r io.Reader, schema string) (PrimitiveUnionTestRecord, error) {
	t := NewPrimitiveUnionTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePrimitiveUnionTestRecord(r PrimitiveUnionTestRecord, w io.Writer) error {
	var err error
	err = writeUnionNullStringLongIntFloatDoubleBool(r.UnionField, w)
	if err != nil {
		return err
	}
	return err
}

func (r PrimitiveUnionTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveUnionTestRecord(r, w)
}

func (r PrimitiveUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"UnionField\",\"type\":[\"null\",\"string\",\"long\",\"int\",\"float\",\"double\",\"boolean\"]}],\"name\":\"PrimitiveUnionTestRecord\",\"type\":\"record\"}"
}

func (r PrimitiveUnionTestRecord) SchemaName() string {
	return "PrimitiveUnionTestRecord"
}

func (_ PrimitiveUnionTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveUnionTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.UnionField = NewUnionNullStringLongIntFloatDoubleBool()

		return r.UnionField
	}
	panic("Unknown field index")
}

func (r *PrimitiveUnionTestRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.UnionField = nil
		return
	}
	panic("Unknown field index")
}

func (r *PrimitiveUnionTestRecord) NullField(i int) {
	switch i {
	case 0:
		r.UnionField = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ PrimitiveUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ PrimitiveUnionTestRecord) Finalize()                        {}

func (_ PrimitiveUnionTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(PrimitiveUnionTestRecordAvroCRC64Fingerprint)
}

func (r PrimitiveUnionTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["UnionField"], err = json.Marshal(r.UnionField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PrimitiveUnionTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["UnionField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UnionField); err != nil {
			return err
		}
	} else {
		r.UnionField = NewUnionNullStringLongIntFloatDoubleBool()

		r.UnionField = nil
	}
	return nil
}
