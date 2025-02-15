package utils

import (
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	translator ut.Translator
)

var ErrGetTranslator = errors.New("error getting translations of errors")

func InitTranslate(locale string) (err error) {
	// Modify the validator engine properties in the gin framework to achieve translation effects
	engine, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		fmt.Println("Error getting gin framework's validator")
	}

	idT := id.New()
	enT := en.New()

	uni := ut.New(enT, idT, enT)

	translator, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}

	switch locale {
	case "en":
		en_translations.RegisterDefaultTranslations(engine, translator)
	case "id":
		id_translations.RegisterDefaultTranslations(engine, translator)
	}

	return nil
}

func ValidateError(c *gin.Context, err error) ([]ErrorField, error) {

	if errors.Is(err, io.EOF) {
		return nil, errors.New("request body is empty")
	}

	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, err
	}

	errorFields := make([]ErrorField, len(errors))
	for i, v := range errors {
		errorFields[i] = ErrorField{
			Field:   v.Field(),
			Message: v.Translate(translator),
		}
	}

	return errorFields, nil
}
