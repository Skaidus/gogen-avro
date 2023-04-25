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

// Common information related to the event which must be included in any clean event
type HeaderworksData struct {
	// Unique identifier for the event used for de-duplication and tracing.
	Uuid *UnionNullDatatypeUUID `json:"uuid"`
	// Fully qualified name of the host that generated the event that generated the data.
	Hostname *UnionNullString `json:"hostname"`
	// Trace information not redundant with this object
	Trace *UnionNullHeaderworksTrace `json:"trace"`
}

const HeaderworksDataAvroCRC64Fingerprint = "6<\xf6?EE\xcd\v"

func NewHeaderworksData() HeaderworksData {
	r := HeaderworksData{}
	r.Uuid = nil
	r.Hostname = nil
	r.Trace = nil
	return r
}

func DeserializeHeaderworksData(r io.Reader) (HeaderworksData, error) {
	t := NewHeaderworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeHeaderworksDataFromSchema(r io.Reader, schema string) (HeaderworksData, error) {
	t := NewHeaderworksData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeHeaderworksData(r HeaderworksData, w io.Writer) error {
	var err error
	err = writeUnionNullDatatypeUUID(r.Uuid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Hostname, w)
	if err != nil {
		return err
	}
	err = writeUnionNullHeaderworksTrace(r.Trace, w)
	if err != nil {
		return err
	}
	return err
}

func (r HeaderworksData) Serialize(w io.Writer) error {
	return writeHeaderworksData(r, w)
}

func (r HeaderworksData) Schema() string {
	return "{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"headerworks.Data\",\"type\":\"record\"}"
}

func (r HeaderworksData) SchemaName() string {
	return "headerworks.Data"
}

func (_ HeaderworksData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ HeaderworksData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ HeaderworksData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ HeaderworksData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ HeaderworksData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ HeaderworksData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ HeaderworksData) SetString(v string)   { panic("Unsupported operation") }
func (_ HeaderworksData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksData) Get(i int) types.Field {
	switch i {
	case 0:
		r.Uuid = NewUnionNullDatatypeUUID()

		return r.Uuid
	case 1:
		r.Hostname = NewUnionNullString()

		return r.Hostname
	case 2:
		r.Trace = NewUnionNullHeaderworksTrace()

		return r.Trace
	}
	panic("Unknown field index")
}

func (r *HeaderworksData) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = nil
		return
	case 1:
		r.Hostname = nil
		return
	case 2:
		r.Trace = nil
		return
	}
	panic("Unknown field index")
}

func (r *HeaderworksData) NullField(i int) {
	switch i {
	case 0:
		r.Uuid = nil
		return
	case 1:
		r.Hostname = nil
		return
	case 2:
		r.Trace = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ HeaderworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ HeaderworksData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ HeaderworksData) HintSize(int)                     { panic("Unsupported operation") }
func (_ HeaderworksData) Finalize()                        {}

func (_ HeaderworksData) AvroCRC64Fingerprint() []byte {
	return []byte(HeaderworksDataAvroCRC64Fingerprint)
}

func (r HeaderworksData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	output["hostname"], err = json.Marshal(r.Hostname)
	if err != nil {
		return nil, err
	}
	output["trace"], err = json.Marshal(r.Trace)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *HeaderworksData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["uuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Uuid); err != nil {
			return err
		}
	} else {
		r.Uuid = NewUnionNullDatatypeUUID()

		r.Uuid = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["hostname"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hostname); err != nil {
			return err
		}
	} else {
		r.Hostname = NewUnionNullString()

		r.Hostname = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["trace"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Trace); err != nil {
			return err
		}
	} else {
		r.Trace = NewUnionNullHeaderworksTrace()

		r.Trace = nil
	}
	return nil
}
