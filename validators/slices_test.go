package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceValidator(t *testing.T) {
	assert.Nil(t, getSliceValidator().MinSize(2).Validate())
	assert.Equal(t, getSliceValidator().MinSize(8).Validate().Error(), "slice too short (min: 8)")
	assert.Nil(t, getSliceValidator().MaxSize(8).Validate())
	assert.Equal(t, getSliceValidator().MaxSize(2).Validate().Error(), "slice too long (max: 2)")

	assert.Equal(t, getSliceValidator().MinMaxSize(2, 2).Validate().Error(), "slice too long (max: 2)")
	assert.Equal(t, getSliceValidator().MinMaxSize(8, 2).Validate().Error(), "slice too short (min: 8)")
	assert.Nil(t, getSliceValidator().MinMaxSize(2, 8).Validate())
	assert.Equal(t, getSliceValidator().MinMaxSize(8, 8).Validate().Error(), "slice too short (min: 8)")

	assert.Nil(t, getSliceValidator().Contains(2).Validate())
	assert.Nil(t, getSliceValidator().Contains(6).Validate())
	assert.Equal(t, getSliceValidator().Contains(123456789).Validate().Error(), "slice does not contain \"123456789\"")
	assert.Nil(t, getSliceValidator().DoesNotContain(0).Validate())
	assert.Nil(t, getSliceValidator().DoesNotContain(7).Validate())
	assert.Equal(t, getSliceValidator().DoesNotContain(3).Validate().Error(), "slice contains \"3\"")
}

func TestDefaultSlice(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	obj := TestObject{Slice: slice1}
	validator := Slice[int](&obj, "slice")

	assert.Equal(t, slice1, obj.Slice)
	validator.Default(slice2)
	assert.Equal(t, slice1, obj.Slice)

	obj = TestObject{Slice: nil}
	validator = Slice[int](&obj, "slice")

	validator.Default(slice2)
	assert.Equal(t, slice2, obj.Slice)
}

func getSliceValidator() *SliceValidator[int, TestObject] {
	return Slice[int](&TestObject{Slice: []int{1, 2, 3, 4, 5, 6}}, "slice")
}
