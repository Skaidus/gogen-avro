// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
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

type AvroContainerHeader struct {
	Magic Magic `json:"magic"`

	Meta map[string]Bytes `json:"meta"`

	Sync Sync `json:"sync"`
}

const AvroContainerHeaderAvroCRC64Fingerprint = "\xc0\x12\x03\xc0wi\xf96"

func NewAvroContainerHeader() AvroContainerHeader {
	r := AvroContainerHeader{}
	r.Meta = make(map[string]Bytes)

	return r
}

func DeserializeAvroContainerHeader(r io.Reader) (AvroContainerHeader, error) {
	t := NewAvroContainerHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAvroContainerHeaderFromSchema(r io.Reader, schema string) (AvroContainerHeader, error) {
	t := NewAvroContainerHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAvroContainerHeader(r AvroContainerHeader, w io.Writer) error {
	var err error
	err = writeMagic(r.Magic, w)
	if err != nil {
		return err
	}
	err = writeMapBytes(r.Meta, w)
	if err != nil {
		return err
	}
	err = writeSync(r.Sync, w)
	if err != nil {
		return err
	}
	return err
}

func (r AvroContainerHeader) Serialize(w io.Writer) error {
	return writeAvroContainerHeader(r, w)
}

func (r AvroContainerHeader) Schema() string {
	return "{\"fields\":[{\"name\":\"magic\",\"type\":{\"name\":\"Magic\",\"size\":4,\"type\":\"fixed\"}},{\"name\":\"meta\",\"type\":{\"type\":\"map\",\"values\":\"bytes\"}},{\"name\":\"sync\",\"type\":{\"name\":\"Sync\",\"size\":16,\"type\":\"fixed\"}}],\"name\":\"AvroContainerHeader\",\"type\":\"record\"}"
}

func (r AvroContainerHeader) SchemaName() string {
	return "AvroContainerHeader"
}

func (_ AvroContainerHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ AvroContainerHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AvroContainerHeader) Get(i int) types.Field {
	switch i {
	case 0:
		w := MagicWrapper{Target: &r.Magic}

		return w

	case 1:
		r.Meta = make(map[string]Bytes)

		w := MapBytesWrapper{Target: &r.Meta}

		return &w

	case 2:
		w := SyncWrapper{Target: &r.Sync}

		return w

	}
	panic("Unknown field index")
}

func (r *AvroContainerHeader) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AvroContainerHeader) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AvroContainerHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AvroContainerHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AvroContainerHeader) HintSize(int)                     { panic("Unsupported operation") }
func (_ AvroContainerHeader) Finalize()                        {}

func (_ AvroContainerHeader) AvroCRC64Fingerprint() []byte {
	return []byte(AvroContainerHeaderAvroCRC64Fingerprint)
}

func (r AvroContainerHeader) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["magic"], err = json.Marshal(r.Magic)
	if err != nil {
		return nil, err
	}
	output["meta"], err = json.Marshal(r.Meta)
	if err != nil {
		return nil, err
	}
	output["sync"], err = json.Marshal(r.Sync)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AvroContainerHeader) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["magic"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Magic); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for magic")
	}
	val = func() json.RawMessage {
		if v, ok := fields["meta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Meta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for meta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sync"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sync); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sync")
	}
	return nil
}
