package model

import validation "github.com/go-ozzo/ozzo-validation"

// Обязательное поле, если выполняется условие
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
