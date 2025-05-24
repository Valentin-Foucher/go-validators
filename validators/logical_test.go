package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogicalValidation(t *testing.T) {
	assert.Nil(t,
		Validate(
			Or(
				getStringZeroValidator().IsZeroValue(),
				getStringZeroValidator().IsNotZeroValue(),
			),
		),
	)

	assert.NotNil(t,
		Validate(
			And(
				Or(
					getStringZeroValidator().IsZeroValue(),
					getStringZeroValidator().IsNotZeroValue(),
				),
				Or(
					getStringZeroValidator().IsNotZeroValue(),
					getStringZeroValidator().NotEqual(""),
				),
			),
		),
	)
}
