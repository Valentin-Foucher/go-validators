package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntValidator(t *testing.T) {
	assert.Nil(t, getIntZeroValidator().Equal(0).Validate())
	assert.Equal(t, getIntZeroValidator().Equal(1).Validate().Error(), "int is not equal to \"1\"")
	assert.Equal(t, getIntZeroValidator().IsNotZeroValue().Validate().Error(), "int has a zero value")
	assert.Nil(t, getIntZeroValidator().NotEqual(1).Validate())
	assert.Equal(t, getIntZeroValidator().NotEqual(0).Validate().Error(), "int is equal to \"0\"")
	assert.Nil(t, getIntZeroValidator().IsZeroValue().Validate())
	assert.Nil(t, getIntZeroValidator().OneOf(0, 123456, 654321).Validate())
	assert.Equal(t, getIntZeroValidator().OneOf(123, 456, 789).Validate().Error(), "int has an invalid value")

	assert.Nil(t, getIntNonZeroValidator().Equal(1).Validate())
	assert.Equal(t, getIntNonZeroValidator().Equal(123456).Validate().Error(), "int is not equal to \"123456\"")
	assert.Equal(t, getIntNonZeroValidator().IsZeroValue().Validate().Error(), "int has not a zero value")
	assert.Nil(t, getIntNonZeroValidator().NotEqual(123456).Validate())
	assert.Equal(t, getIntNonZeroValidator().NotEqual(1).Validate().Error(), "int is equal to \"1\"")
	assert.Nil(t, getIntNonZeroValidator().IsNotZeroValue().Validate())
	assert.Nil(t, getIntNonZeroValidator().OneOf(0, 1, 23456).Validate())
	assert.Equal(t, getIntNonZeroValidator().OneOf(123, 456, 789).Validate().Error(), "int has an invalid value")

	assert.Nil(t, getIntZeroValidator().Gt(-1).Validate())
	assert.Equal(t, getIntZeroValidator().Gt(0).Validate().Error(), "int is not greater than 0")
	assert.Nil(t, getIntZeroValidator().Lt(1).Validate())
	assert.Equal(t, getIntZeroValidator().Lt(0).Validate().Error(), "int is not lower than 0")
	assert.Nil(t, getIntZeroValidator().Gte(0).Validate())
	assert.Equal(t, getIntZeroValidator().Gte(1).Validate().Error(), "int is lower than 1")
	assert.Nil(t, getIntZeroValidator().Lte(0).Validate())
	assert.Equal(t, getIntZeroValidator().Lte(-1).Validate().Error(), "int is greater than -1")
}

func TestDefaultInt(t *testing.T) {
	int1 := 123
	int2 := 456
	obj := TestObject{Int: int1}
	validator := Int(&obj, "int")

	assert.Equal(t, int1, obj.Int)
	validator.Default(int2)
	assert.Equal(t, int1, obj.Int)

	obj = TestObject{Int: 0}
	validator = Int(&obj, "int")

	validator.Default(int2)
	assert.Equal(t, int2, obj.Int)
}

func getIntZeroValidator() *IntValidator[TestObject] {
	return Int(&TestObject{Int: 0}, "int")
}

func getIntNonZeroValidator() *IntValidator[TestObject] {
	return Int(&TestObject{Int: 1}, "int")
}

func TestFloatValidator(t *testing.T) {
	assert.Nil(t, getFloatZeroValidator().Equal(0.0).Validate())
	assert.Equal(t, getFloatZeroValidator().Equal(1.0).Validate().Error(), "float64 is not equal to \"1\"")
	assert.Equal(t, getFloatZeroValidator().IsNotZeroValue().Validate().Error(), "float64 has a zero value")
	assert.Nil(t, getFloatZeroValidator().NotEqual(1.0).Validate())
	assert.Equal(t, getFloatZeroValidator().NotEqual(0.0).Validate().Error(), "float64 is equal to \"0\"")
	assert.Nil(t, getFloatZeroValidator().IsZeroValue().Validate())
	assert.Nil(t, getFloatZeroValidator().OneOf(0.0, 0.123456, 0.654321).Validate())
	assert.Equal(t, getFloatZeroValidator().OneOf(123, 456, 789).Validate().Error(), "float64 has an invalid value")

	assert.Nil(t, getFloatNonZeroValidator().Equal(1.123456789).Validate())
	assert.Equal(t, getFloatNonZeroValidator().Equal(123456.123456789).Validate().Error(), "float64 is not equal to \"123456.123456789\"")
	assert.Equal(t, getFloatNonZeroValidator().IsZeroValue().Validate().Error(), "float64 has not a zero value")
	assert.Nil(t, getFloatNonZeroValidator().NotEqual(1.12345678910112123126485949849819498489198581).Validate())
	assert.Equal(t, getFloatNonZeroValidator().NotEqual(1.123456789).Validate().Error(), "float64 is equal to \"1.123456789\"")
	assert.Nil(t, getFloatNonZeroValidator().IsNotZeroValue().Validate())
	assert.Nil(t, getFloatNonZeroValidator().OneOf(1.123456789, 0.0).Validate())
	assert.Equal(t, getFloatNonZeroValidator().OneOf(123, 456, 789).Validate().Error(), "float64 has an invalid value")

	assert.Nil(t, getFloatZeroValidator().Gt(-0.000001).Validate())
	assert.Equal(t, getFloatZeroValidator().Gt(0).Validate().Error(), "float64 is not greater than 0")
	assert.Nil(t, getFloatZeroValidator().Lt(0.000001).Validate())
	assert.Equal(t, getFloatZeroValidator().Lt(0).Validate().Error(), "float64 is not lower than 0")
	assert.Nil(t, getFloatZeroValidator().Gte(0).Validate())
	assert.Equal(t, getFloatZeroValidator().Gte(1).Validate().Error(), "float64 is lower than 1")
	assert.Nil(t, getFloatZeroValidator().Lte(0).Validate())
	assert.Equal(t, getFloatZeroValidator().Lte(-1).Validate().Error(), "float64 is greater than -1")
}

func TestDefaultFloat(t *testing.T) {
	float1 := 1.23
	float2 := 4.56
	obj := TestObject{Float64: float1}
	validator := Float(&obj, "float64")

	assert.Equal(t, float1, obj.Float64)
	validator.Default(float2)
	assert.Equal(t, float1, obj.Float64)

	obj = TestObject{Float64: 0}
	validator = Float(&obj, "float64")

	validator.Default(float2)
	assert.Equal(t, float2, obj.Float64)
}

func getFloatZeroValidator() *FloatValidator[float64, TestObject] {
	return Float(&TestObject{Float64: 0.0}, "float64")
}

func getFloatNonZeroValidator() *FloatValidator[float64, TestObject] {
	return Float(&TestObject{Float64: 1.123456789}, "float64")
}
