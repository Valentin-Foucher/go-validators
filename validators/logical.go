package validators

import "errors"

type logicalValidator struct {
	executors *[]func() ValidationError
}

func (v logicalValidator) Validate() ValidationError {
	for _, executor := range *v.executors {
		if err := executor(); err != nil {
			return err
		}
	}

	return nil
}

func Or(validators ...Validator) *logicalValidator {
	executors := []func() ValidationError{
		func() ValidationError {
			errorMessage := ""

			for _, validator := range validators {
				err := validator.Validate()
				if err == nil {
					return nil
				}

				if errorMessage != "" {
					errorMessage += ", "
				}

				errorMessage += err.Error()
			}

			return errors.New(errorMessage)
		},
	}

	return &logicalValidator{executors: &executors}
}

func And(validators ...Validator) *logicalValidator {
	executors := []func() ValidationError{
		func() ValidationError {
			for _, validator := range validators {
				if err := validator.Validate(); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return &logicalValidator{executors: &executors}
}
