package validators

import "slices"

type SliceValidator[T comparable, U any] struct {
	*sizableValidator[[]T, T, any, U]
}

func Slice[T comparable, U any](object *U, field string) *SliceValidator[T, U] {
	return &SliceValidator[T, U]{
		sizableValidator: newSizableValidator[[]T, T, any](object, field),
	}
}

func SliceFromValue[T comparable](field string, value []T) *SliceValidator[T, any] {
	return &SliceValidator[T, any]{
		sizableValidator: newSizableValueValidator[[]T, T, any](field, value),
	}
}

func (v *SliceValidator[T, U]) Contains(expected T) *SliceValidator[T, U] {
	v.chain(func(inner *fieldValidator[[]T, U]) ValidationError {
		if slices.Contains(inner.Value, expected) {
			return nil
		}

		return createValidationError("%s does not contain \"%v\"", v.Field, expected)
	})

	return v
}

func (v *SliceValidator[T, U]) DoesNotContain(notExpected T) *SliceValidator[T, U] {
	v.chain(func(inner *fieldValidator[[]T, U]) ValidationError {
		if slices.Contains(inner.Value, notExpected) {
			return createValidationError("%s contains \"%v\"", v.Field, notExpected)
		}

		return nil
	})

	return v
}
