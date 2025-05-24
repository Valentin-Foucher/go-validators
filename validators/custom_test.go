package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidator(t *testing.T) {
	assert.Equal(t, Email(&TestObject{Email: ""}, "email").IsValid().Validate().Error(), "Invalid email for field email")
	assert.Equal(t, Email(&TestObject{Email: "test@test.tsetestsetestte16541654154154154145s165165165sq1d65qs1ds56q1ds6q51d56sq1d56qs1d65qs16d51sq651dqs651dqs65d1sq651dqs651s"}, "email").IsValid().Validate().Error(), "Invalid email for field email")
	assert.Equal(t, Email(&TestObject{Email: "t@"}, "email").IsValid().Validate().Error(), "Invalid email for field email")
	assert.Nil(t, Email(&TestObject{Email: "test@test.fr"}, "email").IsValid().Validate())
}

func TestDefaultEmail(t *testing.T) {
	email1 := "aaa@aaa.aaa"
	email2 := "bbb@bbb.bbb"
	obj := TestObject{Email: email1}
	validator := Email(&obj, "email")

	assert.Equal(t, email1, obj.Email)
	validator.Default(email2)
	assert.Equal(t, email1, obj.Email)

	obj = TestObject{Email: ""}
	validator = Email(&obj, "email")
	validator.Default(email2)
	assert.Equal(t, email2, obj.Email)
}
