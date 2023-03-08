package service

import (
    "context"
    "courses/domain/models"
    "courses/domain/repository"
    "net/http"
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

func (s *Service) GetDirections(ctx context.Context) ([]models.Direction, error) {
    return s.Repository.GetDirections(ctx)
}

func (s *Service) GetDirection(ctx context.Context, id int) (models.Direction, error) {
    return s.Repository.GetDirection(ctx, id)
}

func (s *Service) GetCourses(ctx context.Context) ([]models.Course, error) {
    return s.Repository.GetCourses(ctx)
}

func (s *Service) GetCourse(ctx context.Context, id int) (models.Course, error) {
    return s.Repository.GetCourse(ctx, id)
}

func (s *Service) GetStudents(ctx context.Context) ([]models.Student, error) {
    return s.Repository.GetStudents(ctx)
}

func (s *Service) GetStudent(ctx context.Context, id int) (models.Student, error) {
    return s.Repository.GetStudent(ctx, id)
}

func (s *Service) CreateDirection(ctx context.Context, direction models.Direction) (error, int) {
    err := validateField(direction.Name, "name"); if err != nil { return err, http.StatusBadRequest}

    err = s.Repository.CreateDirection(ctx, direction.Name)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusCreated
}

func (s *Service) CreateCourse(ctx context.Context, course models.Course) (error, int) {
    err := validateCourse(course); if err != nil { return err, http.StatusBadRequest }
    err = validateField(course.DirectionId, "direction_id"); if err != nil { return err, http.StatusBadRequest }

    err = s.Repository.CreateCourse(ctx, course)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusCreated
}

func (s *Service) CreateStudent(ctx context.Context, student models.Student) (error, int) {
    err := validateStudent(student); if err != nil { return err, http.StatusBadRequest }
    err = validateField(student.CourseId, "course_id"); if err != nil { return err, http.StatusBadRequest }

    err = s.Repository.CreateStudent(ctx, student)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusCreated
}

func (s *Service) UpdateDirection(ctx context.Context, direction models.Direction) (error, int) {
    err := validateField(direction.Name, "name"); if err != nil { return err, http.StatusBadRequest}
    err = validateField(direction.Id, "id"); if err != nil { return err, http.StatusBadRequest}

    err = s.Repository.UpdateDirection(ctx, direction)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusOK
}

func (s *Service) UpdateCourse(ctx context.Context, course models.Course) (error, int) {
    err := validateCourse(course); if err != nil { return err, http.StatusBadRequest }
    err = validateField(course.Id, "id"); if err != nil { return err, http.StatusBadRequest}

    err = s.Repository.UpdateCourse(ctx, course)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusOK
}

func (s *Service) UpdateStudent(ctx context.Context, student models.Student) (error, int) {
    err := validateStudent(student); if err != nil { return err, http.StatusBadRequest }
    err = validateField(student.Id, "id"); if err != nil { return err, http.StatusBadRequest}

    err = s.Repository.UpdateStudent(ctx, student)
    if err != nil {
        return err, http.StatusInternalServerError
    }

    return nil, http.StatusOK
}

func (s *Service) DeleteDirection(ctx context.Context, id int) error {
    return s.Repository.DeleteDirection(ctx, id)
}

func (s *Service) DeleteCourse(ctx context.Context, id int) error {
    return s.Repository.DeleteCourse(ctx, id)
}

func (s *Service) DeleteStudent(ctx context.Context, id int) error {
    return s.Repository.DeleteStudent(ctx, id)
}
