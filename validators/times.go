package validators

import "time"

type TimeValidator[T any] struct {
	*comparableValidator[time.Time, T]
}

func Time[T any](object *T, field string) *TimeValidator[T] {
	return &TimeValidator[T]{
		comparableValidator: newComparableValidator[time.Time](object, field),
	}
}

func TimeFromValue(field string, value time.Time) *TimeValidator[any] {
	return &TimeValidator[any]{
		comparableValidator: newComparableValueValidator(field, value),
	}
}

func (v *TimeValidator[T]) After(date time.Time) *TimeValidator[T] {
	v.chain(func(inner *fieldValidator[time.Time, T]) ValidationError {
		if !inner.Value.After(date) {
			return createValidationError("%s is not after than \"%v\"", v.Field, date)
		}

		return nil
	})

	return v
}

func (v *TimeValidator[T]) Before(date time.Time) *TimeValidator[T] {
	v.chain(func(inner *fieldValidator[time.Time, T]) ValidationError {
		if !inner.Value.Before(date) {
			return createValidationError("%s is not before \"%v\"", v.Field, date)
		}

		return nil
	})

	return v
}
