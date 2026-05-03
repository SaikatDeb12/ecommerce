package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/SaikatDeb12/ecommerce/internal/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate  = validator.New()
	SecretKey string
)

func ParseBody(body io.Reader, out interface{}) error {
	return json.NewDecoder(body).Decode(out)
}

func EncodeBody(body http.ResponseWriter, out interface{}) error {
	return json.NewEncoder(body).Encode(out)
}

func ValidateStruct(payload interface{}) error {
	return ValidateStruct(payload)
}

func GetEnvVariables(key string) string {
	return os.Getenv(key)
}

func ResponseError(w http.ResponseWriter, statusCode int, err error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	newError := models.ErrorModel{
		Message:    message,
		Error:      err.Error(),
		StatusCode: statusCode,
	}

	if err := EncodeBody(w, newError); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		err := EncodeBody(w, body)
		if err != nil {
			fmt.Printf("%+v", err)
		}
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
