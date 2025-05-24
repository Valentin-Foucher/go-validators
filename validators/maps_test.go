package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapValidator(t *testing.T) {

	assert.Nil(t, getMapValidator().MinSize(2).Validate())
	assert.Equal(t, getMapValidator().MinSize(8).Validate().Error(), "Map too short (min: 8)")
	assert.Nil(t, getMapValidator().MaxSize(8).Validate())
	assert.Equal(t, getMapValidator().MaxSize(2).Validate().Error(), "Map too long (max: 2)")

	assert.Equal(t, getMapValidator().MinMaxSize(2, 2).Validate().Error(), "Map too long (max: 2)")
	assert.Equal(t, getMapValidator().MinMaxSize(8, 2).Validate().Error(), "Map too short (min: 8)")
	assert.Nil(t, getMapValidator().MinMaxSize(2, 8).Validate())
	assert.Equal(t, getMapValidator().MinMaxSize(8, 8).Validate().Error(), "Map too short (min: 8)")

	assert.Nil(t, getMapValidator().ContainsKey(1).Validate())
	assert.Nil(t, getMapValidator().ContainsValue(2).Validate())
	assert.Equal(t, getMapValidator().ContainsValue(1).Validate().Error(), "Map does not contain value \"1\"")
	assert.Equal(t, getMapValidator().ContainsKey(2).Validate().Error(), "Map does not contain key \"2\"")

	assert.Nil(t, getMapValidator().DoesNotContainKey(2).Validate())
	assert.Nil(t, getMapValidator().DoesNotContainValue(1).Validate())
	assert.Equal(t, getMapValidator().DoesNotContainValue(2).Validate().Error(), "Map contains value \"2\"")
	assert.Equal(t, getMapValidator().DoesNotContainKey(1).Validate().Error(), "Map contains key \"1\"")
}

func TestDefaultMap(t *testing.T) {
	map1 := map[int]int{1: 2}
	map2 := map[int]int{3: 4, 5: 6}

	obj := TestObject{Map: map1}
	validator := Map[int, int](&obj, "map")

	assert.Equal(t, map1, obj.Map)
	validator.Default(map2)
	assert.Equal(t, map1, obj.Map)

	obj = TestObject{Map: nil}
	validator = Map[int, int](&obj, "map")

	validator.Default(map2)
	assert.Equal(t, map2, obj.Map)
}

func getMapValidator() *MapValidator[int, int, TestObject] {
	return Map[int, int](&TestObject{Map: map[int]int{1: 2, 3: 4, 5: 6}}, "Map")
}
