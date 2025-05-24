package validators

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestObject struct {
	Int     int
	String  string
	Float64 float64
	Bool    bool
	Time    time.Time
	Email   string
	Slice   []int
	Map     map[int]int
	Pointer *int
}

func TestValidation(t *testing.T) {
	assert.NotNil(t,
		Validate(
			getStringZeroValidator().IsZeroValue(),
			getStringZeroValidator().IsNotZeroValue(),
		),
	)

	assert.Nil(t,
		Validate(
			getStringZeroValidator().IsZeroValue(),
			getStringZeroValidator().IsZeroValue(),
			getStringZeroValidator().Equal(""),
			Email(&TestObject{Email: "test@email.fr"}, "Email").IsValid(),
		),
	)

	assert.NotNil(t,
		Validate(
			Map[int, int](&TestObject{Map: map[int]int{1: 2}}, "Map").ContainsKey(2),
		),
	)
}

func TestValidationWithMultipleConditions(t *testing.T) {
	assert.Nil(t,
		Validate(
			Slice[int](&TestObject{Slice: []int{1, 2, 3}}, "Slice").
				DoesNotContain(0).
				Contains(3).
				MinSize(3),
			Int(&TestObject{Int: 1}, "Int").
				Equal(1).
				IsNotZeroValue(),
		),
	)

	assert.NotNil(t,
		Validate(
			Slice[int](&TestObject{Slice: []int{1, 2, 3}}, "Slice").
				DoesNotContain(0).
				Contains(3).
				MinSize(7),
		),
	)
}

func TestValidationWithLowercaseField(t *testing.T) {
	assert.NotPanics(t,
		func() {
			String(&TestObject{String: ""}, "string")
		},
	)

	assert.NotPanics(t,
		func() {
			String(&TestObject{String: ""}, "String")
		},
	)

	assert.Panics(t,
		func() {
			String(&TestObject{String: ""}, "Int")
		},
	)

	assert.Panics(t,
		func() {
			String(&TestObject{String: ""}, "Anotherfield")
		},
	)
}
