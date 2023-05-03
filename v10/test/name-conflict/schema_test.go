package avro

import (
	"io"
	"testing"

	"github.com/Skaidus/gogen-avro/v10/container"
	"github.com/Skaidus/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTrip(t,
		func() container.AvroRecord { return &NameConflictTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeNameConflictTestRecord(r)
			return &record, err
		})
}
