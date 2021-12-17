package everything_test

import (
	"testing"

	"github.com/foxclore/go-utils/v1/everything"
)

type absTestTable struct {
	variable    interface{}
	expected    float64
	description string
}

func TestAbs(t *testing.T) {
	var table []absTestTable
	var (
		intP int = 10
		intN int = -10

		int8P int8 = 10
		int8N int8 = -10

		int16P int16 = 10
		int16N int16 = -10

		int32P int32 = 10
		int32N int32 = -10

		int64P int64 = 10
		int64N int64 = -10

		uintV uint = 15

		uint8V uint8 = 15

		uint16V uint16 = 15

		uint32V uint32 = 15

		uint64V uint64 = 15

		float32P float32 = 5.4
		float32N float32 = -5.4

		float64P float64 = 5.4
		float64N float64 = -5.4
	)

	table = append(table, absTestTable{intP, 10, "int positive"})
	table = append(table, absTestTable{intN, 10, "int negative"})

	table = append(table, absTestTable{int8P, 10, "int8 positive"})
	table = append(table, absTestTable{int8N, 10, "int8 negative"})

	table = append(table, absTestTable{int16P, 10, "int16 positive"})
	table = append(table, absTestTable{int16N, 10, "int16 negative"})

	table = append(table, absTestTable{int32P, 10, "int32 positive"})
	table = append(table, absTestTable{int32N, 10, "int32 negative"})

	table = append(table, absTestTable{int64P, 10, "int64 positive"})
	table = append(table, absTestTable{int64N, 10, "int64 negative"})

	table = append(table, absTestTable{uintV, 15, "uint"})

	table = append(table, absTestTable{uint8V, 15, "uint8"})

	table = append(table, absTestTable{uint16V, 15, "uint16"})

	table = append(table, absTestTable{uint32V, 15, "uint32"})

	table = append(table, absTestTable{uint64V, 15, "uint64"})

	var f32Res float32 = 5.4

	table = append(table, absTestTable{float32P, float64(f32Res), "float32 positive"})
	table = append(table, absTestTable{float32N, float64(f32Res), "float32 negative"})

	table = append(table, absTestTable{float64P, 5.4, "float64 positive"})
	table = append(table, absTestTable{float64N, 5.4, "float64 negative"})

	for _, tx := range table {
		absoluteVal := everything.Abs(tx.variable)
		if absoluteVal != tx.expected {
			t.Errorf("Failed test for %s, expected %f, got %f", tx.description, tx.expected, absoluteVal)
		}
	}
}
