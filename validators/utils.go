package validators

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isZeroValue[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}

func getAttribute[T, U any](object U, field string) (T, error) {
	var value T
	var ok bool

	reflectedObject := reflect.ValueOf(object)
	reflectedField := reflect.Indirect(reflectedObject).FieldByNameFunc(func(f string) bool {
		return strings.EqualFold(f, field)
	})

	if !reflectedField.IsValid() {
		return value, getAttributeError(field)
	}

	if value, ok = reflectedField.Interface().(T); ok {
		return value, nil
	}

	return value, fmt.Errorf("incorrect type for attribute %s", field)
}

func setAttribute[T, U any](object *U, field string, defaultValue T) error {
	reflectedObject := reflect.ValueOf(object).Elem()

	for reflectedObject.Kind() == reflect.Ptr || reflectedObject.Kind() == reflect.Interface {
		if reflectedObject.IsNil() {
			return errors.New("nil value encountered while dereferencing")
		}
		reflectedObject = reflectedObject.Elem()
	}

	if !reflectedObject.CanAddr() {
		return errors.New("please make sure that the structure you are setting attribute on is publicly exposed")
	}

	reflectedField := reflectedObject.FieldByNameFunc(func(f string) bool {
		return strings.EqualFold(f, field)
	})

	if !reflectedField.IsValid() {
		return getAttributeError(field)
	}

	reflectedField.Set(reflect.ValueOf(defaultValue))

	return nil
}

func getAttributeError(field string) error {
	return fmt.Errorf("impossible to get attribute %s", field)
}

func isAnEmail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return emailRegex.MatchString(e)
}
