package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValidator(t *testing.T) {
	assert.Nil(t, getStringZeroValidator().Equal("").Validate())
	assert.Equal(t, getStringZeroValidator().Equal("test").Validate().Error(), "string is not equal to \"test\"")
	assert.Nil(t, getStringZeroValidator().NotEqual("test").Validate())
	assert.Equal(t, getStringZeroValidator().NotEqual("").Validate().Error(), "string is equal to \"\"")
	assert.Equal(t, getStringZeroValidator().IsNotZeroValue().Validate().Error(), "string has a zero value")
	assert.Nil(t, getStringZeroValidator().IsZeroValue().Validate())
	assert.Nil(t, getStringZeroValidator().OneOf("", "456", "test").Validate())
	assert.Equal(t, getStringZeroValidator().OneOf("123", "456", "tset").Validate().Error(), "string has an invalid value")

	assert.Nil(t, getStringNonZeroValidator().Equal("test").Validate())
	assert.Equal(t, getStringNonZeroValidator().Equal("123456").Validate().Error(), "string is not equal to \"123456\"")
	assert.Nil(t, getStringNonZeroValidator().NotEqual("").Validate())
	assert.Equal(t, getStringNonZeroValidator().NotEqual("test").Validate().Error(), "string is equal to \"test\"")
	assert.Equal(t, getStringNonZeroValidator().IsZeroValue().Validate().Error(), "string has not a zero value")
	assert.Nil(t, getStringNonZeroValidator().IsNotZeroValue().Validate())
	assert.Nil(t, getStringNonZeroValidator().OneOf("123", "456", "test").Validate())
	assert.Equal(t, getStringNonZeroValidator().OneOf("123", "456", "tset").Validate().Error(), "string has an invalid value")
}

func TestDefaultString(t *testing.T) {
	string1 := "abc"
	string2 := "def"
	obj := TestObject{String: string1}
	validator := String(&obj, "string")

	assert.Equal(t, string1, obj.String)
	validator.Default(string2)
	assert.Equal(t, string1, obj.String)

	obj = TestObject{String: ""}
	validator = String(&obj, "string")

	validator.Default(string2)
	assert.Equal(t, string2, obj.String)
}

func TestSubstring(t *testing.T) {
	assert.Nil(t, getStringNonZeroValidator().Contains("es").Validate())
	assert.Equal(t, getStringNonZeroValidator().Contains("123456").Validate().Error(), "string does not contain \"123456\"")
	assert.Nil(t, getStringNonZeroValidator().DoesNotContain("123456").Validate())
	assert.Equal(t, getStringNonZeroValidator().DoesNotContain("es").Validate().Error(), "string contains \"es\"")
	assert.Nil(t, getStringNonZeroValidator().StartsWith("tes").Validate())
	assert.Equal(t, getStringNonZeroValidator().StartsWith("123456").Validate().Error(), "string does not start with \"123456\"")
	assert.Nil(t, getStringNonZeroValidator().DoesNotStartWith("123456").Validate())
	assert.Equal(t, getStringNonZeroValidator().DoesNotStartWith("tes").Validate().Error(), "string starts with \"tes\"")
	assert.Nil(t, getStringNonZeroValidator().EndsWith("est").Validate())
	assert.Equal(t, getStringNonZeroValidator().EndsWith("123456").Validate().Error(), "string does not end with \"123456\"")
	assert.Nil(t, getStringNonZeroValidator().DoesNotEndWith("123456").Validate())
	assert.Equal(t, getStringNonZeroValidator().DoesNotEndWith("est").Validate().Error(), "string ends with \"est\"")
}

func TestRegexMatching(t *testing.T) {
	obj := TestObject{String: "11223za"}

	assert.Equal(t, String(&obj, "string").MatchRegex(`1\x2*3{abcdefg}[a-z]{2}`).Validate().Error(), "Invalid regex \"1\\x2*3{abcdefg}[a-z]{2}\" for field string")
	assert.Equal(t, String(&obj, "string").MatchRegex("12*3[a-z]{2}").Validate().Error(), "string does not match \"12*3[a-z]{2}\"")
	assert.Nil(t, String(&obj, "string").MatchRegex("1{1,4}2*3[a-z]{2}").Validate())
}

func TestSetPostValidation(t *testing.T) {
	obj := TestObject{String: "10987654321"}
	validator := String(&obj, "string")

	validator.UpdateBeforeValidation(func(o *TestObject) string {
		if o.Int == 2 {
			return "123456"
		} else {
			return "456789"
		}
	}).Validate()

	assert.Equal(t, obj.String, "456789")

	obj.Int = 2

	validator.UpdateBeforeValidation(func(o *TestObject) string {
		if o.Int == 2 {
			return "123456"
		} else {
			return "456789"
		}
	}).Validate()

	assert.Equal(t, obj.String, "123456")
}

func getStringZeroValidator() *StringValidator[TestObject] {
	return String(&TestObject{String: ""}, "string")
}

func getStringNonZeroValidator() *StringValidator[TestObject] {
	return String(&TestObject{String: "test"}, "string")
}
