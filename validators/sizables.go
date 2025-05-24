package validators

type sizable[T, U comparable] interface {
	~string | ~[]T | ~map[T]U
}

type sizableValidator[T sizable[U, V], U, V comparable, W any] struct {
	*fieldValidator[T, W]
}

func (v *sizableValidator[T, U, V, W]) MaxSize(size int) *sizableValidator[T, U, V, W] {
	v.chain(func(inner *fieldValidator[T, W]) ValidationError {
		if len(inner.Value) > size {
			return createValidationError("%s too long (max: %d)", inner.Field, size)
		}

		return nil
	})

	return v
}

func (v *sizableValidator[T, U, V, W]) MinSize(size int) *sizableValidator[T, U, V, W] {
	v.chain(func(inner *fieldValidator[T, W]) ValidationError {
		if len(inner.Value) < size {
			return createValidationError("%s too short (min: %d)", inner.Field, size)
		}

		return nil
	})

	return v
}

func (v *sizableValidator[T, U, V, W]) MinMaxSize(min, max int) *sizableValidator[T, U, V, W] {
	return v.MinSize(min).MaxSize(max)
}

func newSizableValidator[T sizable[U, V], U, V comparable, W any](object *W, field string) *sizableValidator[T, U, V, W] {
	value, err := getAttribute[T](object, field)
	if err != nil {
		panic(err)
	}

	return &sizableValidator[T, U, V, W]{
		fieldValidator: &fieldValidator[T, W]{
			executors: make([]func(*fieldValidator[T, W]) ValidationError, 0),
			Field:     field,
			Value:     value,
			Object:    object,
		},
	}
}

func newSizableValueValidator[T sizable[U, V], U, V comparable](field string, value T) *sizableValidator[T, U, V, any] {
	return &sizableValidator[T, U, V, any]{
		fieldValidator: &fieldValidator[T, any]{
			executors: make([]func(*fieldValidator[T, any]) ValidationError, 0),
			Field:     field,
			Value:     value,
		},
	}
}
