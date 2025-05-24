package validators

import (
	"fmt"
	"regexp"
	"strings"
)

type StringValidator[T any] struct {
	*sizableValidator[string, string, any, T]
	*comparableValidator[string, T]
}

func String[T any](object *T, field string) *StringValidator[T] {
	return &StringValidator[T]{
		sizableValidator:    newSizableValidator[string, string, any](object, field),
		comparableValidator: newComparableValidator[string](object, field),
	}
}

func StringFromValue(field string, value string) *StringValidator[any] {
	return &StringValidator[any]{
		comparableValidator: newComparableValueValidator(field, value),
	}
}

func (v StringValidator[T]) Validate() ValidationError {
	if v.comparableValidator.priorExecutorSetter != nil {
		v.comparableValidator.priorExecutorSetter()
	}

	if err := v.comparableValidator.Validate(); err != nil {
		return err
	}

	if err := v.sizableValidator.Validate(); err != nil {
		return err
	}

	return nil
}

func (v *StringValidator[T]) Default(defaultValue string) *StringValidator[T] {
	v.comparableValidator.Default(defaultValue)
	v.sizableValidator.Value = v.comparableValidator.Value
	return v
}

func (v *StringValidator[T]) UpdateBeforeValidation(f func(*T) string) *StringValidator[T] {
	v.comparableValidator.UpdateBeforeValidation(f)
	return v
}

// MatchRegex checks if the string matches the provided regex pattern.
func (v *StringValidator[T]) MatchRegex(regex string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		match, err := regexp.Compile(fmt.Sprintf("\\b%s\\b", regex))
		if err != nil {
			return createValidationError("Invalid regex \"%s\" for field %s", regex, inner.Field)
		}

		if !match.MatchString(inner.Value) {
			return createValidationError("%s does not match \"%s\"", inner.Field, regex)
		}

		return nil
	})

	return v
}

// Contains checks if the string contains the specified substring.
func (v *StringValidator[T]) Contains(substring string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if !strings.Contains(inner.Value, substring) {
			return createValidationError("%s does not contain \"%s\"", inner.Field, substring)
		}

		return nil
	})

	return v
}

// DoesNotContain checks if the string does not contain the specified substring.
func (v *StringValidator[T]) DoesNotContain(substring string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if strings.Contains(inner.Value, substring) {
			return createValidationError("%s contains \"%s\"", inner.Field, substring)
		}

		return nil
	})

	return v
}

// StartsWith checks if the string starts with the specified prefix.
func (v *StringValidator[T]) StartsWith(prefix string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if !strings.HasPrefix(inner.Value, prefix) {
			return createValidationError("%s does not start with \"%s\"", inner.Field, prefix)
		}

		return nil
	})

	return v
}

// DoesNotStartWith checks if the string does not start with the specified prefix.
func (v *StringValidator[T]) DoesNotStartWith(prefix string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if strings.HasPrefix(inner.Value, prefix) {
			return createValidationError("%s starts with \"%s\"", inner.Field, prefix)
		}

		return nil
	})

	return v
}

// EndsWith checks if the string ends with the specified suffix.
func (v *StringValidator[T]) EndsWith(suffix string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if !strings.HasSuffix(inner.Value, suffix) {
			return createValidationError("%s does not end with \"%s\"", inner.Field, suffix)
		}

		return nil
	})

	return v
}

// DoesNotEndWith checks if the string does not end with the specified suffix.
func (v *StringValidator[T]) DoesNotEndWith(suffix string) *StringValidator[T] {
	v.comparableValidator.chain(func(inner *fieldValidator[string, T]) ValidationError {
		if strings.HasSuffix(inner.Value, suffix) {
			return createValidationError("%s ends with \"%s\"", inner.Field, suffix)
		}

		return nil
	})

	return v
}
