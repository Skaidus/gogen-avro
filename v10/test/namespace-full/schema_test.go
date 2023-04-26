package avro

import (
	"io"
	"testing"

	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &ComAvroTestSample{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeComAvroTestSample(r)
			return &record, err
		})
}
