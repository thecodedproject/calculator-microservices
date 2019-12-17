package ops_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thecodedproject/calculator_microservices/add/ops"
)

func TestAdd(t *testing.T) {
	tests := []struct{
		name string
		inputs []float64
		expected float64
	}{
		{
			name: "Zero inputs returns zero",
			expected: 0.0,
		},
		{
			name: "One input returns value",
			inputs: []float64{2.0},
			expected: 2.0,
		},
		{
			name: "Multiple inputs returns sum of all",
			inputs: []float64{2.0, 3.4, 5.21, 2.01},
			expected: 12.62,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ops.Add(test.inputs)
			assert.Equal(t, test.expected, actual)
		})
	}

}
