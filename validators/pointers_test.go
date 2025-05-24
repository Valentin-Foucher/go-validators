package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](value T) *T {
	return &value
}

func TestPointerValidator(t *testing.T) {
	assert.Nil(t, getPointerValidator().IsDefined().Validate())
	assert.Equal(t, getPointerValidator().IsNotDefined().Validate().Error(), "pointer is defined")

	assert.Nil(t, getNilPointerValidator().IsNotDefined().Validate())
	assert.Equal(t, getNilPointerValidator().IsDefined().Validate().Error(), "pointer is not defined")

	assert.Nil(t, getPointerValidatorWithValidInnerValidator().Validate())
	assert.Equal(t, getPointerValidatorWithInvalidInnerValidator().Validate().Error(), "pointer has not a zero value")
}

func TestDefaultPointer(t *testing.T) {
	pointer1 := ptr(123)
	pointer2 := ptr(456)
	obj := TestObject{Pointer: pointer1}
	validator := Pointer[int](&obj, "pointer")

	assert.Equal(t, pointer1, obj.Pointer)
	validator.Default(pointer2)
	assert.Equal(t, pointer1, obj.Pointer)

	obj = TestObject{Pointer: nil}
	validator = Pointer[int](&obj, "pointer")

	validator.Default(pointer2)
	assert.Equal(t, pointer2, obj.Pointer)

	// unsupported
	assert.Panics(t, func() {
		Pointer(&TestObject{Pointer: ptr(2)}, "pointer", func(field string, i int) Validator {
			return IntFromValue(field, i).IsNotZeroValue().Default(123456)
		})
	})
}

func getPointerValidator() *PointerValidator[*int, TestObject] {
	return Pointer[int](&TestObject{Pointer: ptr(2)}, "pointer")
}

func getNilPointerValidator() *PointerValidator[*int, TestObject] {
	var nilPointer *int
	return Pointer[int](&TestObject{Pointer: nilPointer}, "pointer")
}

func getPointerValidatorWithValidInnerValidator() *PointerValidator[*int, TestObject] {
	return Pointer(&TestObject{Pointer: ptr(2)}, "pointer", func(field string, i int) Validator {
		return IntFromValue(field, i).IsNotZeroValue()
	})
}

func getPointerValidatorWithInvalidInnerValidator() *PointerValidator[*int, TestObject] {
	return Pointer(&TestObject{Pointer: ptr(2)}, "pointer", func(field string, i int) Validator {
		return IntFromValue(field, i).IsZeroValue()
	})
}
