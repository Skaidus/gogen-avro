// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/masmovil/gogen-avro/v10/compiler"
	"github.com/masmovil/gogen-avro/v10/container"
	"github.com/masmovil/gogen-avro/v10/vm"
)

func NewNumberRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNumberRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NumberRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewNumberRecordReader(r io.Reader) (*NumberRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNumberRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NumberRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NumberRecordReader) Read() (NumberRecord, error) {
	t := NewNumberRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
