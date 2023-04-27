package avro

import (
	"io"
	"testing"

	"github.com/Skaidus/gogen-avro/v10/container"
	"github.com/Skaidus/gogen-avro/v10/test"
	evolution "github.com/Skaidus/gogen-avro/v10/test/alias-field/evolution"
)

func TestEvolution(t *testing.T) {
	test.RoundTripEvolution(t,
		func() container.AvroRecord { return &AliasRecord{} },
		func() container.AvroRecord { return &evolution.AliasRecord{} },
		func(r io.Reader, schema string) (container.AvroRecord, error) {
			record, err := evolution.DeserializeAliasRecordFromSchema(r, schema)
			return &record, err
		})
}
