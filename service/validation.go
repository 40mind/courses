package service

import (
	"courses/domain/models"
	"fmt"
	"github.com/jimlawless/whereami"
	"log"
	"reflect"
	"regexp"
	"time"
)


func validateCourse(course models.Course) error {
	err := validateField(course.Name, "name"); if err != nil { return err }
	err = validateField(course.NumOfClasses, "num_of_classes"); if err != nil { return err }
	err = validateField(course.ClassTime, "class_time"); if err != nil { return err }
	err = validateField(course.WeekDays, "week_days"); if err != nil { return err }
	err = validateField(course.FirstClassDate, "first_class_date"); if err != nil { return err }
	err = validateField(course.LastClassDate, "last_class_date"); if err != nil { return err }
	err = validateField(course.Price, "price"); if err != nil { return err }
	err = validateField(course.Info, "info"); if err != nil { return err }

	return nil
}

func validateStudent(student models.Student) error {
	err := validateField(student.Name, "name"); if err != nil { return err }
	err = validateField(student.Surname, "surname"); if err != nil { return err }
	err = validateField(student.Patronymic, "patronymic"); if err != nil { return err }
	err = validateField(student.Email, "email"); if err != nil { return err }
	err = validateField(student.Phone, "phone"); if err != nil { return err }
	err = validateField(student.Comment, "comment"); if err != nil { return err }

	err = validateEmail(student.Email); if err != nil { return err }
	err = validatePhone(student.Phone); if err != nil { return err }

	return nil
}

func validateField(value any, field string) error {
	switch reflect.TypeOf(value).String() {
	case "string":
		strValue, ok := value.(string); if !ok {
			return fmt.Errorf("can't convert field %s to string", field)
		}

		if strValue == "" {
			log.Printf("empty field %s\n", field)
			return fmt.Errorf("empty field %s", field)
		}

		return nil

	case "int":
		intValue, ok := value.(int); if !ok {
			return fmt.Errorf("can't convert field %s to int", field)
		}

		if intValue <= 0 {
			return fmt.Errorf("field %s should be positive", field)
		}

		return nil

	case "float64":
		floatValue, ok := value.(float64); if !ok {
			return fmt.Errorf("can't convert field %s to float64", field)
		}

		if floatValue <= 0 {
			return fmt.Errorf("field %s should be positive", field)
		}

		return nil

	case "time.Time":
		_, ok := value.(time.Time); if !ok {
			return fmt.Errorf("can't convert field %s to time.Time", field)
		}

		return nil

	default:
		return fmt.Errorf("unknown type")
	}
}

func validateEmail(email string) error {
	pattern := `^\w+@\w+\.\w+$`
	match, err := regexp.Match(pattern, []byte(email))
	if err != nil {
		log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	if !match {
		log.Printf("validation error: email is invalid: %s\n", whereami.WhereAmI())
		return fmt.Errorf("validation error: email is invalid")
	}

	return nil
}

func validatePhone(phone string) error {
	pattern1 := `^\d+$`
	pattern2 := `^\+\d+$`

	match1, err := regexp.Match(pattern1, []byte(phone))
	if err != nil {
		log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	match2, err := regexp.Match(pattern2, []byte(phone))
	if err != nil {
		log.Printf("%s: %s: %s\n", controllerError, err.Error(), whereami.WhereAmI())
		return fmt.Errorf(controllerError)
	}

	if !match1 && !match2 {
		log.Printf("validation error: phone is invalid: %s\n", whereami.WhereAmI())
		return fmt.Errorf("validation error: phone is invalid")
	}

	return nil
}
