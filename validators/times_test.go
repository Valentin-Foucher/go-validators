package validators

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeValidator(t *testing.T) {
	var zero time.Time

	assert.Nil(t, getTimeZeroValidator().Equal(zero).Validate())
	assert.Contains(t, getTimeZeroValidator().Equal(time.Now()).Validate().Error(), "time is not equal to ")
	assert.Equal(t, getTimeZeroValidator().IsNotZeroValue().Validate().Error(), "time has a zero value")
	assert.Nil(t, getTimeZeroValidator().NotEqual(time.Now()).Validate())
	assert.Contains(t, getTimeZeroValidator().NotEqual(zero).Validate().Error(), "time is equal to ")
	assert.Nil(t, getTimeZeroValidator().IsZeroValue().Validate())
	assert.Nil(t, getTimeZeroValidator().OneOf(zero, time.Now(), time.Now()).Validate())
	assert.Equal(t, getTimeZeroValidator().OneOf(time.Now(), time.Now()).Validate().Error(), "time has an invalid value")

	now := time.Now()
	assert.Nil(t, getTimeNonZeroValidator(now).Equal(now).Validate())
	assert.Contains(t, getTimeNonZeroValidator(now).Equal(now.Add(1)).Validate().Error(), "time is not equal to")
	assert.Equal(t, getTimeNonZeroValidator(now).IsZeroValue().Validate().Error(), "time has not a zero value")
	assert.Nil(t, getTimeNonZeroValidator(now).NotEqual(time.Now()).Validate())
	assert.Contains(t, getTimeNonZeroValidator(now).NotEqual(now).Validate().Error(), "time is equal to ")
	assert.Nil(t, getTimeNonZeroValidator(now).IsNotZeroValue().Validate())
	assert.Nil(t, getTimeNonZeroValidator(now).OneOf(now, time.Now(), time.Now()).Validate())
	assert.Equal(t, getTimeNonZeroValidator(now).OneOf(time.Now(), time.Now()).Validate().Error(), "time has an invalid value")

	assert.Nil(t, getTimeNonZeroValidator(time.Now()).After(now).Validate())
	assert.Contains(t, getTimeNonZeroValidator(now).After(now).Validate().Error(), "time is not after ")
	assert.Nil(t, getTimeNonZeroValidator(now).Before(time.Now()).Validate())
	assert.Contains(t, getTimeNonZeroValidator(now).Before(now).Validate().Error(), "time is not before ")
}

func TestDefaultTime(t *testing.T) {
	var zeroTime time.Time
	time1 := time.Now()
	time2 := time.Now()
	obj := TestObject{Time: time1}
	validator := Time(&obj, "time")

	assert.Equal(t, time1, obj.Time)
	validator.Default(time2)
	assert.Equal(t, time1, obj.Time)

	obj = TestObject{Time: zeroTime}
	validator = Time(&obj, "time")

	validator.Default(time2)
	assert.Equal(t, time2, obj.Time)
}

func getTimeZeroValidator() *TimeValidator[TestObject] {
	var zero time.Time
	return Time(&TestObject{Time: zero}, "time")
}

func getTimeNonZeroValidator(now time.Time) *TimeValidator[TestObject] {
	return Time(&TestObject{Time: now}, "time")
}
