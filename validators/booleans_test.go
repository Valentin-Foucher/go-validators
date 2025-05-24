package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolValidator(t *testing.T) {
	assert.Nil(t, getBoolZeroValidator().Equal(false).Validate())
	assert.Equal(t, getBoolZeroValidator().Equal(true).Validate().Error(), "bool is not equal to \"true\"")
	assert.Equal(t, getBoolZeroValidator().IsNotZeroValue().Validate().Error(), "bool has a zero value")
	assert.Nil(t, getBoolZeroValidator().NotEqual(true).Validate())
	assert.Equal(t, getBoolZeroValidator().NotEqual(false).Validate().Error(), "bool is equal to \"false\"")
	assert.Nil(t, getBoolZeroValidator().IsZeroValue().Validate())
	assert.Nil(t, getBoolZeroValidator().OneOf(false, false, true).Validate())
	assert.Equal(t, getBoolZeroValidator().OneOf(true).Validate().Error(), "bool has an invalid value")

	assert.Nil(t, getBoolNonZeroValidator().Equal(true).Validate())
	assert.Equal(t, getBoolNonZeroValidator().Equal(false).Validate().Error(), "bool is not equal to \"false\"")
	assert.Equal(t, getBoolNonZeroValidator().IsZeroValue().Validate().Error(), "bool has not a zero value")
	assert.Nil(t, getBoolNonZeroValidator().NotEqual(false).Validate())
	assert.Equal(t, getBoolNonZeroValidator().NotEqual(true).Validate().Error(), "bool is equal to \"true\"")
	assert.Nil(t, getBoolNonZeroValidator().IsNotZeroValue().Validate())
	assert.Nil(t, getBoolNonZeroValidator().OneOf(true, false, true).Validate())
	assert.Equal(t, getBoolNonZeroValidator().OneOf(false).Validate().Error(), "bool has an invalid value")
}

func TestDefaultBool(t *testing.T) {
	bool1 := true
	bool2 := false
	obj := TestObject{Bool: bool1}
	validator := Bool(&obj, "bool")

	assert.Equal(t, bool1, obj.Bool)
	validator.Default(bool2)
	assert.Equal(t, bool1, obj.Bool)

	obj = TestObject{Bool: false}
	validator = Bool(&obj, "bool")

	validator.Default(bool1)
	assert.Equal(t, bool1, obj.Bool)
}

func getBoolZeroValidator() *BoolValidator[TestObject] {
	return Bool(&TestObject{Bool: false}, "bool")
}

func getBoolNonZeroValidator() *BoolValidator[TestObject] {
	return Bool(&TestObject{Bool: true}, "bool")
}
