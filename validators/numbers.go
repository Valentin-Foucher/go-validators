package validators

type IntValidator[T any] struct {
	*orderedValidator[int, T]
}

type FloatValidator[T float64 | float32, U any] struct {
	*orderedValidator[T, U]
}

func Int[T any](object *T, field string) *IntValidator[T] {
	return &IntValidator[T]{
		orderedValidator: newOrderedValidator[int](object, field),
	}
}

func IntFromValue(field string, value int) *IntValidator[any] {
	return &IntValidator[any]{
		orderedValidator: newOrderedValueValidator(field, value),
	}
}

func Float[T any](object *T, field string) *FloatValidator[float64, T] {
	return &FloatValidator[float64, T]{
		orderedValidator: newOrderedValidator[float64](object, field),
	}
}

func Float32[T any](object *T, field string) *FloatValidator[float32, T] {
	return &FloatValidator[float32, T]{
		orderedValidator: newOrderedValidator[float32](object, field),
	}
}

func FloatFromValue(field string, value float64) *FloatValidator[float64, any] {
	return &FloatValidator[float64, any]{
		orderedValidator: newOrderedValueValidator(field, value),
	}
}

func Float32FromValue(field string, value float32) *FloatValidator[float32, any] {
	return &FloatValidator[float32, any]{
		orderedValidator: newOrderedValueValidator(field, value),
	}
}
