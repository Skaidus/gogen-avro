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

func NewComCompanySharedTypeTwoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewComCompanySharedTypeTwo()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ComCompanySharedTypeTwoReader struct {
	r io.Reader
	p *vm.Program
}

func NewComCompanySharedTypeTwoReader(r io.Reader) (*ComCompanySharedTypeTwoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewComCompanySharedTypeTwo()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ComCompanySharedTypeTwoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ComCompanySharedTypeTwoReader) Read() (ComCompanySharedTypeTwo, error) {
	t := NewComCompanySharedTypeTwo()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
