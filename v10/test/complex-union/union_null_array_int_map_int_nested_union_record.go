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

type UnionNullArrayIntMapIntNestedUnionRecordTypeEnum int

const (
	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumArrayInt UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 1

	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumMapInt UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 2

	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNestedUnionRecord UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 3
)

type UnionNullArrayIntMapIntNestedUnionRecord struct {
	Null              *types.NullVal
	ArrayInt          []int32
	MapInt            map[string]int32
	NestedUnionRecord NestedUnionRecord
	UnionType         UnionNullArrayIntMapIntNestedUnionRecordTypeEnum
}

func writeUnionNullArrayIntMapIntNestedUnionRecord(r *UnionNullArrayIntMapIntNestedUnionRecord, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumArrayInt:
		return writeArrayInt(r.ArrayInt, w)
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumMapInt:
		return writeMapInt(r.MapInt, w)
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNestedUnionRecord:
		return writeNestedUnionRecord(r.NestedUnionRecord, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayIntMapIntNestedUnionRecord")
}

func NewUnionNullArrayIntMapIntNestedUnionRecord() *UnionNullArrayIntMapIntNestedUnionRecord {
	return &UnionNullArrayIntMapIntNestedUnionRecord{}
}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) Serialize(w io.Writer) error {
	return writeUnionNullArrayIntMapIntNestedUnionRecord(r, w)
}

func DeserializeUnionNullArrayIntMapIntNestedUnionRecord(r io.Reader) (*UnionNullArrayIntMapIntNestedUnionRecord, error) {
	t := NewUnionNullArrayIntMapIntNestedUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullArrayIntMapIntNestedUnionRecordFromSchema(r io.Reader, schema string) (*UnionNullArrayIntMapIntNestedUnionRecord, error) {
	t := NewUnionNullArrayIntMapIntNestedUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) Schema() string {
	return "[\"null\",{\"items\":\"int\",\"type\":\"array\"},{\"type\":\"map\",\"values\":\"int\"},{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"}],\"name\":\"NestedUnionRecord\",\"type\":\"record\"}]"
}

func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetInt(v int32)    { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) SetLong(v int64) {

	r.UnionType = (UnionNullArrayIntMapIntNestedUnionRecordTypeEnum)(v)
}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayInt = make([]int32, 0)
		return &ArrayIntWrapper{Target: (&r.ArrayInt)}
	case 2:
		r.MapInt = make(map[string]int32)
		return &MapIntWrapper{Target: (&r.MapInt)}
	case 3:
		r.NestedUnionRecord = NewNestedUnionRecord()
		return &types.Record{Target: (&r.NestedUnionRecord)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) Finalize() {}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumArrayInt:
		return json.Marshal(map[string]interface{}{"array": r.ArrayInt})
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumMapInt:
		return json.Marshal(map[string]interface{}{"map": r.MapInt})
	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNestedUnionRecord:
		return json.Marshal(map[string]interface{}{"NestedUnionRecord": r.NestedUnionRecord})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayIntMapIntNestedUnionRecord")
}

func (r *UnionNullArrayIntMapIntNestedUnionRecord) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayInt)
	}
	if value, ok := fields["map"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.MapInt)
	}
	if value, ok := fields["NestedUnionRecord"]; ok {
		r.UnionType = 3
		return json.Unmarshal([]byte(value), &r.NestedUnionRecord)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayIntMapIntNestedUnionRecord")
}
