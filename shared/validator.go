package shared

import (
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Checks if val is a validated
func Validate(val interface{}) error {
	validate := validator.New()

	validators := map[string]validator.Func{
		"register":  isRegNo,
		"phone":     isPhoneNo,
		"birthdate": isBirthDate,
	}
	for key, value := range validators {
		if err := validate.RegisterValidation(key, value); err != nil {
			return err
		}
	}

	if err := validate.Struct(val); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

// Checks if val is a email
func IsEmail(val string) bool {
	_, err := mail.ParseAddress(val)
	return err == nil
}

// Checks if val is a vehicle's plate no
func IsPlateNo(val string) bool {
	ok, _ := regexp.MatchString("^[0-9]{4}[а-яА-ЯөӨүҮ]{3}$|^[0-9]{4}[а-яА-ЯөӨүҮ]{2}$", val)
	return ok
}

// Checks if val is a numeric
func IsNumeric(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}

// Checks if val is a citizen's register number
func isRegNo(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.String {
		return false
	}

	ok, _ := regexp.MatchString("^[а-яА-ЯөӨүҮ]{2}[0-9]{8}$", field.String())
	return ok
}

// Checks if val is a phone number
func isPhoneNo(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.String {
		return false
	}

	ok, _ := regexp.MatchString(`^\d{8}$`, field.String())
	return ok
}

// Checks if val is a birthdate
func isBirthDate(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.String {
		return false
	}

	_, err := time.Parse(time.DateOnly, field.String())
	return err == nil
}
