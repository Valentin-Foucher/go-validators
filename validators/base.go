package validators

import (
	"fmt"

	"errors"
)

type Validator interface {
	Validate() ValidationError
}

type fieldValidator[T any, U any] struct {
	executors           []func(*fieldValidator[T, U]) ValidationError
	priorExecutorSetter func()

	Field  string
	Value  T
	Object *U
}

func (v *fieldValidator[T, U]) Validate() ValidationError {
	if v.priorExecutorSetter != nil {
		v.priorExecutorSetter()
	}

	for _, executor := range v.executors {
		if err := executor(v); err != nil {
			return err
		}
	}

	return nil
}

func (v *fieldValidator[T, U]) chain(f func(*fieldValidator[T, U]) ValidationError) {
	v.executors = append(v.executors, f)
}

// Default requires Object property type to be exposed publicly (capitalized)
func (v *fieldValidator[T, U]) Default(defaultValue T) *fieldValidator[T, U] {
	if v.Object == nil {
		panic(errors.New("unsupported default call"))
	}

	if !isZeroValue(v.Value) {
		return v
	}

	if err := setAttribute(v.Object, v.Field, defaultValue); err != nil {
		panic(err)
	}

	v.Value = defaultValue

	return v
}

// UpdateBeforeValidation allows setting a value before validation occurs.
func (v *fieldValidator[T, U]) UpdateBeforeValidation(f func(*U) T) *fieldValidator[T, U] {
	v.priorExecutorSetter = func() {
		if v.Object == nil {
			panic(errors.New("unsupported default call"))
		}

		newValue := f(v.Object)

		if err := setAttribute(v.Object, v.Field, newValue); err != nil {
			panic(err)
		}

		v.Value = newValue
	}

	return v
}

// This function is the main entry point for validation.
// It accepts a variable number of Validator instances and returns a ValidationError
// if any of the validators fail. If all validators pass, it returns nil.
func Validate(validators ...Validator) ValidationError {
	errorMessage := ""
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			if errorMessage != "" {
				errorMessage += ", "
			}

			errorMessage += err.Error()
		}
	}

	if errorMessage != "" {
		return ValidationError(errors.New(errorMessage))
	}

	return nil
}

func createValidationError(format string, args ...any) ValidationError {
	return ValidationError(fmt.Errorf(format, args...))
}
