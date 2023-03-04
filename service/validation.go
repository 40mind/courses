package service

import (
	"fmt"
	"github.com/jimlawless/whereami"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func validateStringForm(r *http.Request, key string) (string, error) {
	value := r.FormValue(key)
	if value == "" {
		log.Printf("empty field %s", key)
		return "", fmt.Errorf("empty field %s", key)
	}

	return value, nil
}

func validateIntForm(r *http.Request, key string) (int, error) {
	value := r.FormValue(key)
	if value == "" {
		log.Printf("empty field %s", key)
		return -1, fmt.Errorf("empty field %s", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("can't convert field %s to int", key)
		return -1, fmt.Errorf("can't convert field %s to int", key)
	}

	return intValue, nil
}

func validateFloatForm(r *http.Request, key string) (float64, error) {
	value := r.FormValue(key)
	if value == "" {
		log.Printf("empty field %s", key)
		return -1, fmt.Errorf("empty field %s", key)
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("can't convert field %s to float", key)
		return -1, fmt.Errorf("can't convert field %s to float", key)
	}

	return floatValue, nil
}

func validateDateForm(r *http.Request, key string) (time.Time, error) {
	value := r.FormValue(key)
	if value == "" {
		log.Printf("empty field %s", key)
		return time.Time{}, fmt.Errorf("empty field %s", key)
	}

	dateValue, err := time.Parse(dateLayout, value)
	if err != nil {
		log.Printf("can't convert field %s to date", key)
		return time.Time{}, fmt.Errorf("can't convert field %s to date", key)
	}

	return dateValue, nil
}

func validateEmail(email string) error {
	pattern := `^\w+@\w+\.\w+$`
	match, err := regexp.Match(pattern, []byte(email))
	if err != nil {
		log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	if !match {
		log.Printf("validation error: email is invalid: %s", whereami.WhereAmI())
		return fmt.Errorf("validation error: email is invalid")
	}

	return nil
}

func validatePhone(phone string) error {
	pattern1 := `^\d+$`
	pattern2 := `^\+\d+$`

	match1, err := regexp.Match(pattern1, []byte(phone))
	if err != nil {
		log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	match2, err := regexp.Match(pattern2, []byte(phone))
	if err != nil {
		log.Printf("%s: %s: %s", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	if !match1 && !match2 {
		log.Printf("validation error: phone is invalid: %s", whereami.WhereAmI())
		return fmt.Errorf("validation error: phone is invalid")
	}

	return nil
}
