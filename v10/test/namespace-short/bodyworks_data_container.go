// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/Skaidus/gogen-avro/v10/compiler"
	"github.com/Skaidus/gogen-avro/v10/container"
	"github.com/Skaidus/gogen-avro/v10/vm"
)

func NewBodyworksDataWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewBodyworksData()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type BodyworksDataReader struct {
	r io.Reader
	p *vm.Program
}

func NewBodyworksDataReader(r io.Reader) (*BodyworksDataReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewBodyworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &BodyworksDataReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r BodyworksDataReader) Read() (BodyworksData, error) {
	t := NewBodyworksData()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
