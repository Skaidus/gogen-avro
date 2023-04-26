// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
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

// The test record
type Event struct {
	// Unique ID for this event.
	Id string `json:"id"`
	// Start IP of this observation's IP range.
	Start_ip IPAddress `json:"start_ip"`
	// End IP of this observation's IP range.
	End_ip IPAddress `json:"end_ip"`
}

const EventAvroCRC64Fingerprint = "W\xbc\xe5D\xb1l>j"

func NewEvent() Event {
	r := Event{}
	return r
}

func DeserializeEvent(r io.Reader) (Event, error) {
	t := NewEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventFromSchema(r io.Reader, schema string) (Event, error) {
	t := NewEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEvent(r Event, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeIPAddress(r.Start_ip, w)
	if err != nil {
		return err
	}
	err = writeIPAddress(r.End_ip, w)
	if err != nil {
		return err
	}
	return err
}

func (r Event) Serialize(w io.Writer) error {
	return writeEvent(r, w)
}

func (r Event) Schema() string {
	return "{\"doc\":\"The test record\",\"fields\":[{\"doc\":\"Unique ID for this event.\",\"name\":\"id\",\"type\":\"string\"},{\"doc\":\"Start IP of this observation's IP range.\",\"name\":\"start_ip\",\"type\":{\"aliases\":[\"ip_address\"],\"name\":\"IPAddress\",\"size\":16,\"type\":\"fixed\"}},{\"doc\":\"End IP of this observation's IP range.\",\"name\":\"end_ip\",\"type\":\"IPAddress\"}],\"name\":\"event\",\"subject\":\"event\",\"type\":\"record\"}"
}

func (r Event) SchemaName() string {
	return "event"
}

func (_ Event) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Event) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Event) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Event) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Event) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Event) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Event) SetString(v string)   { panic("Unsupported operation") }
func (_ Event) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Event) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := IPAddressWrapper{Target: &r.Start_ip}

		return w

	case 2:
		w := IPAddressWrapper{Target: &r.End_ip}

		return w

	}
	panic("Unknown field index")
}

func (r *Event) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Event) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Event) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Event) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Event) HintSize(int)                     { panic("Unsupported operation") }
func (_ Event) Finalize()                        {}

func (_ Event) AvroCRC64Fingerprint() []byte {
	return []byte(EventAvroCRC64Fingerprint)
}

func (r Event) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["start_ip"], err = json.Marshal(r.Start_ip)
	if err != nil {
		return nil, err
	}
	output["end_ip"], err = json.Marshal(r.End_ip)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Event) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["start_ip"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Start_ip); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for start_ip")
	}
	val = func() json.RawMessage {
		if v, ok := fields["end_ip"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.End_ip); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for end_ip")
	}
	return nil
}
