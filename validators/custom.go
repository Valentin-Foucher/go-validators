package validators

type EmailValidator[T any] struct {
	*StringValidator[T]
}

func Email[T any](object *T, field string) *EmailValidator[T] {
	return &EmailValidator[T]{
		StringValidator: String(object, field),
	}
}

func EmailFromValue(field string, value string) *EmailValidator[any] {
	return &EmailValidator[any]{
		StringValidator: StringFromValue(field, value),
	}
}

func (v *EmailValidator[T]) IsValid() *EmailValidator[T] {
	v.comparableValidator.chain(func(validator *fieldValidator[string, T]) ValidationError {
		if !isAnEmail(validator.Value) {
			return createValidationError("Invalid email for field %s", v.comparableValidator.Field)
		}

		return nil
	})

	return v
}
