// Code generated by Skaidus. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewNestedUnionRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNestedUnionRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NestedUnionRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewNestedUnionRecordReader(r io.Reader) (*NestedUnionRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNestedUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NestedUnionRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NestedUnionRecordReader) Read() (NestedUnionRecord, error) {
	t := NewNestedUnionRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
