package validators

type BoolValidator[T any] struct {
	*comparableValidator[bool, T]
}

func Bool[T any](object *T, field string) *BoolValidator[T] {
	return &BoolValidator[T]{
		comparableValidator: newComparableValidator[bool](object, field),
	}
}

func BoolFromValue(field string, value bool) *BoolValidator[any] {
	return &BoolValidator[any]{
		comparableValidator: newComparableValueValidator(field, value),
	}
}
