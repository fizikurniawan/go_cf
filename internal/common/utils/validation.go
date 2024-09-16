package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors memformat validasi error dari validator
func FormatValidationErrors(err error, dto interface{}) map[string][]string {
	fieldErrors := make(map[string][]string)

	// Periksa apakah error adalah error JSON kosong
	if strings.Contains(err.Error(), "EOF") {
		fieldErrors["non_field_error"] = []string{"Request body is empty or invalid JSON"}
		return fieldErrors
	}

	// Dapatkan nama field dari struct DTO
	fieldNames := getFieldNames(dto)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())       // Nama field dalam camelCase
			snakeCaseField := fieldNames[field]       // Konversi ke snake_case
			msg := humanizeValidationMessage(e.Tag()) // Buat pesan lebih human readable
			if _, exists := fieldErrors[snakeCaseField]; !exists {
				fieldErrors[snakeCaseField] = []string{}
			}
			fieldErrors[snakeCaseField] = append(fieldErrors[snakeCaseField], msg)
		}
	} else {
		fieldErrors["non_field_error"] = []string{err.Error()}
	}

	return fieldErrors
}

// getFieldNames mengembalikan peta nama field dari camelCase ke snake_case
func getFieldNames(dto interface{}) map[string]string {
	fieldNames := make(map[string]string)
	v := reflect.ValueOf(dto).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			jsonTag = strings.Split(jsonTag, ",")[0] // Ambil tag JSON, abaikan opsi binding
			fieldNames[strings.ToLower(field.Name)] = jsonTag
		}
	}
	return fieldNames
}

// HumanizeValidationMessage mengubah tag validasi menjadi pesan yang lebih human readable
func humanizeValidationMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "min":
		return "This field must be at least the specified length"
	case "max":
		return "This field must not exceed the specified length"
	// Tambahkan case lain sesuai dengan kebutuhan
	default:
		return "Invalid input"
	}
}
