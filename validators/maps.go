package validators

type MapValidator[T, U comparable, V any] struct {
	*sizableValidator[map[T]U, T, U, V]
	keys   []T
	values []U
}

func Map[T, U comparable, V any](object *V, field string) *MapValidator[T, U, V] {
	sizableValidator := newSizableValidator[map[T]U, T, U](object, field)

	keys := make([]T, 0)
	values := make([]U, 0)
	for key, item := range sizableValidator.Value {
		keys = append(keys, key)
		values = append(values, item)
	}

	return &MapValidator[T, U, V]{
		sizableValidator: sizableValidator,
		keys:             keys,
		values:           values,
	}
}

func MapFromValue[T, U comparable](field string, value map[T]U) *MapValidator[T, U, any] {
	return &MapValidator[T, U, any]{
		sizableValidator: newSizableValueValidator[map[T]U, T, U](field, value),
	}
}

// ContainsKey checks if the map contains the expected key.
func (v *MapValidator[T, U, V]) ContainsKey(expected T) *MapValidator[T, U, V] {
	v.chain(func(inner *fieldValidator[map[T]U, V]) ValidationError {
		for _, key := range v.keys {
			if key == expected {
				return nil
			}
		}

		return createValidationError("%s does not contain key \"%v\"", v.Field, expected)
	})

	return v
}

// ContainsValue checks if the map contains the expected value.
func (v *MapValidator[T, U, V]) ContainsValue(expected U) *MapValidator[T, U, V] {
	v.chain(func(inner *fieldValidator[map[T]U, V]) ValidationError {
		for _, value := range v.values {
			if value == expected {
				return nil
			}
		}

		return createValidationError("%s does not contain value \"%v\"", v.Field, expected)
	})

	return v
}

// DoesNotContainKey checks if the map does not contain the expected key.
func (v *MapValidator[T, U, V]) DoesNotContainKey(expected T) *MapValidator[T, U, V] {
	v.chain(func(inner *fieldValidator[map[T]U, V]) ValidationError {
		for _, key := range v.keys {
			if key == expected {
				return createValidationError("%s contains key \"%v\"", v.Field, expected)
			}
		}

		return nil
	})

	return v
}

// DoesNotContainValue checks if the map does not contain the expected value.
func (v *MapValidator[T, U, V]) DoesNotContainValue(expected U) *MapValidator[T, U, V] {
	v.chain(func(inner *fieldValidator[map[T]U, V]) ValidationError {
		for _, value := range v.values {
			if value == expected {
				return createValidationError("%s contains value \"%v\"", v.Field, expected)
			}
		}

		return nil
	})

	return v
}
