package service

import (
    "context"
    "courses/domain/models"
    "courses/domain/repository"
    "net/http"
    "strconv"
    "strings"
)

const (
    controllerError = "controller error"
    dateLayout      = "dd-mm-yyyy hh-mm-ss"
)

type Service struct {
    Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
    return &Service{
        Repository: repository,
    }
}

func (s *Service) GetDirections(ctx context.Context) ([]models.Directions, error) {
    return s.Repository.GetDirections(ctx)
}

func (s *Service) GetDirection(ctx context.Context, id int) (models.Directions, error) {
    return s.Repository.GetDirection(ctx, id)
}

func (s *Service) GetCourses(ctx context.Context) ([]models.Courses, error) {
    return s.Repository.GetCourses(ctx)
}

func (s *Service) GetCourse(ctx context.Context, id int) (models.Courses, error) {
    return s.Repository.GetCourse(ctx, id)
}

func (s *Service) GetStudents(ctx context.Context) ([]models.Students, error) {
    return s.Repository.GetStudents(ctx)
}

func (s *Service) GetStudent(ctx context.Context, id int) (models.Students, error) {
    return s.Repository.GetStudent(ctx, id)
}

func (s *Service) CreateStudent(ctx context.Context, r *http.Request) (error, int) {
    splitURL := strings.Split(r.URL.Path, "/")
    course, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        return err, 400
    }

    name, err := validateStringForm(r, "name")
    if err != nil {
        return err, 400
    }
    surname, err := validateStringForm(r, "surname")
    if err != nil {
        return err, 400
    }
    patronymic, err := validateStringForm(r, "patronymic")
    if err != nil {
        return err, 400
    }
    email, err := validateStringForm(r, "email")
    if err != nil {
        return err, 400
    }
    phone, err := validateStringForm(r, "phone")
    if err != nil {
        return err, 400
    }
    comment, err := validateStringForm(r, "comment")
    if err != nil {
        return err, 400
    }

    err = validateEmail(email)
    if err != nil {
        return err, 400
    }
    err = validatePhone(phone)
    if err != nil {
        return err, 400
    }

    err = s.Repository.CreateStudent(ctx, course, name, surname, patronymic, email, phone, comment)
    if err != nil {
        return err, 500
    }

    return nil, 201
}

func (s *Service) UpdateCourse(ctx context.Context, r *http.Request) (error, int) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        return err, 400
    }

    name, err := validateStringForm(r, "name")
    if err != nil {
        return err, 400
    }
    numOfClasses, err := validateIntForm(r, "num_of_classes")
    if err != nil {
        return err, 400
    }
    classTime, err := validateIntForm(r, "class_time")
    if err != nil {
        return err, 400
    }
    weekDays, err := validateStringForm(r, "week_days")
    if err != nil {
        return err, 400
    }
    firstClassDate, err := validateDateForm(r, "first_class_date")
    if err != nil {
        return err, 400
    }
    lastClassDate, err := validateDateForm(r, "last_class_date")
    if err != nil {
        return err, 400
    }
    price, err := validateFloatForm(r, "price")
    if err != nil {
        return err, 400
    }
    info, err := validateStringForm(r, "info")
    if err != nil {
        return err, 400
    }
    direction, err := validateIntForm(r, "direction")
    if err != nil {
        return err, 400
    }

    err = s.Repository.UpdateCourse(ctx, id, direction, numOfClasses, classTime, name, weekDays, info, firstClassDate, lastClassDate, price)
    if err != nil {
        return err, 500
    }

    return nil, 200
}

func (s *Service) DeleteCourse(ctx context.Context, r *http.Request) (error, int) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL)-1])
    if err != nil {
        return err, 400
    }

    err = s.Repository.DeleteCourse(ctx, id)
    if err != nil {
        return err, 500
    }

    return nil, 200
}
