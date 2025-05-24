package validators

type PointerValidator[T, U any] struct {
	innerValidator Validator
	*fieldValidator[T, U]

	isDefined bool
}

func Pointer[T, U any](object *U, field string, getValidators ...func(string, T) Validator) *PointerValidator[*T, U] {
	var defined bool
	var innerValidator Validator

	value, err := getAttribute[*T](object, field)
	if err != nil {
		panic(err)
	}

	if value == nil {
		defined = false
	} else {
		defined = true
		for _, validator := range getValidators {
			innerValidator = validator(field, *value)
		}
	}

	return &PointerValidator[*T, U]{
		innerValidator: innerValidator,
		fieldValidator: &fieldValidator[*T, U]{
			executors: make([]func(*fieldValidator[*T, U]) ValidationError, 0),
			Field:     field,
			Value:     value,
			Object:    object,
		},
		isDefined: defined,
	}
}

// IsDefined checks if the pointer field is not nil.
func (v PointerValidator[T, U]) IsDefined() PointerValidator[T, U] {
	v.chain(func(*fieldValidator[T, U]) ValidationError {
		if !v.isDefined {
			return createValidationError("%s is not defined", v.Field)
		}

		return nil
	})

	return v
}

// IsNotDefined checks if the pointer field is nil.
func (v PointerValidator[T, U]) IsNotDefined() PointerValidator[T, U] {
	v.chain(func(*fieldValidator[T, U]) ValidationError {
		if v.isDefined {
			return createValidationError("%s is defined", v.Field)
		}

		return nil
	})

	return v
}

func (v PointerValidator[T, U]) Validate() ValidationError {
	if err := v.fieldValidator.Validate(); err != nil {
		return err
	}

	if v.innerValidator != nil {
		return v.innerValidator.Validate()
	}

	return nil
}
