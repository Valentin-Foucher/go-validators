package validators

import (
	"golang.org/x/exp/constraints"
)

type orderedValidator[T constraints.Ordered, U any] struct {
	*comparableValidator[T, U]
}

func (v *orderedValidator[T, U]) Gt(bound T) *orderedValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value <= bound {
			return createValidationError("%s is not greater than %v", inner.Field, bound)
		}

		return nil
	})

	return v
}

func (v *orderedValidator[T, U]) Lt(bound T) *orderedValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value >= bound {
			return createValidationError("%s is not lower than %v", inner.Field, bound)
		}

		return nil
	})

	return v
}

func (v *orderedValidator[T, U]) Gte(bound T) *orderedValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value < bound {
			return createValidationError("%s is lower than %v", inner.Field, bound)
		}

		return nil
	})

	return v
}

func (v *orderedValidator[T, U]) Lte(bound T) *orderedValidator[T, U] {
	v.chain(func(inner *fieldValidator[T, U]) ValidationError {
		if inner.Value > bound {
			return createValidationError("%s is greater than %v", inner.Field, bound)
		}

		return nil
	})

	return v
}

func newOrderedValidator[T constraints.Ordered, U any](object *U, field string) *orderedValidator[T, U] {
	return &orderedValidator[T, U]{
		comparableValidator: newComparableValidator[T](object, field),
	}
}

func newOrderedValueValidator[T constraints.Ordered](field string, value T) *orderedValidator[T, any] {
	return &orderedValidator[T, any]{
		comparableValidator: newComparableValueValidator(field, value),
	}
}
