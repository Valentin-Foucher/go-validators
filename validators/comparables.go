package validators

import "slices"

type comparableValidator[T comparable, U any] struct {
	*fieldValidator[T, U]
}

// Equal checks if the value of the field is equal to the expected value.
func (v *comparableValidator[T, U]) Equal(expected T) *comparableValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value != expected {
			return createValidationError("%s is not equal to \"%v\"", v.Field, expected)
		}

		return nil
	})

	return v
}

// NotEqual checks if the value of the field is not equal to the notExpected value.
func (v *comparableValidator[T, U]) NotEqual(notExpected T) *comparableValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value == notExpected {
			return createValidationError("%s is equal to \"%v\"", v.Field, notExpected)
		}

		return nil
	})

	return v
}

// IsNotZeroValue checks if the value of the field is not left empty.
func (v *comparableValidator[T, U]) IsNotZeroValue() *comparableValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if isZeroValue(inner.Value) {
			return createValidationError("%s has a zero value", v.Field)
		}

		return nil
	})

	return v
}

// IsZeroValue checks if the value of the field is left empty.
func (v *comparableValidator[T, U]) IsZeroValue() *comparableValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if !isZeroValue(inner.Value) {
			return createValidationError("%s has not a zero value", v.Field)
		}

		return nil
	})

	return v
}

// OneOf checks if the value of the field is one of the provided options.
func (v *comparableValidator[T, U]) OneOf(options ...T) *comparableValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if slices.Contains(options, inner.Value) {
			return nil
		}

		return createValidationError("%s has an invalid value", v.Field)
	})

	return v
}

func newComparableValidator[T comparable, U any](object *U, field string) *comparableValidator[T, U] {
	value, err := getAttribute[T](object, field)
	if err != nil {
		panic(err)
	}

	return &comparableValidator[T, U]{
		fieldValidator: &fieldValidator[T, U]{
			executors: make([]func(*fieldValidator[T, U]) ValidationError, 0),
			Field:     field,
			Value:     value,
			Object:    object,
		},
	}
}

func newComparableValueValidator[T comparable](field string, value T) *comparableValidator[T, any] {
	return &comparableValidator[T, any]{
		fieldValidator: &fieldValidator[T, any]{
			executors: make([]func(*fieldValidator[T, any]) ValidationError, 0),
			Field:     field,
			Value:     value,
		},
	}
}
