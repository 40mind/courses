package repository

import (
    "context"
    "courses/domain/models"
    "database/sql"
    "fmt"
    "github.com/jimlawless/whereami"
    "log"
    "time"
)

const DBError = "db error"

type Repository struct {
    DB       *sql.DB
    DBConfig *models.DB
}

func New(db *sql.DB, dbConfig *models.DB) *Repository {
    return &Repository{
        DB:       db,
        DBConfig: dbConfig,
    }
}

func (rep *Repository) CreateDirection(ctx context.Context, name string) error {
    query := "SELECT * FROM graduate_work.create_direction($1)"

    _, err := rep.DB.ExecContext(ctx, query, name)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) GetDirections(ctx context.Context) ([]models.Directions, error) {
    query := "SELECT * FROM graduate_work.get_directions()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return []models.Directions{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Directions{}, nil
    }

    var directions []models.Directions
    for rows.Next() {
        var direction models.Directions
        err = rows.Scan(
            &direction.Name,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
            return []models.Directions{}, fmt.Errorf(DBError)
        }

        directions = append(directions, direction)
    }

    return directions, err
}

func (rep *Repository) GetDirection(ctx context.Context, id int) (models.Directions, error) {
    query := "SELECT * FROM graduate_work.get_direction($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var direction models.Directions
    err := row.Scan(
        &direction.Name,
    )
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return models.Directions{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Directions{}, nil
    }

    return direction, err
}

func (rep *Repository) GetCourses(ctx context.Context) ([]models.Courses, error) {
    query := "SELECT * FROM graduate_work.get_courses()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return []models.Courses{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Courses{}, nil
    }

    var courses []models.Courses
    for rows.Next() {
        var course models.Courses
        err = rows.Scan(
            &course.Name,
            &course.Direction,
            &course.NumOfClasses,
            &course.ClassTime,
            &course.WeekDays,
            &course.FirstClassDate,
            &course.LastClassDate,
            &course.Price,
            &course.Info,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
            return []models.Courses{}, fmt.Errorf(DBError)
        }

        courses = append(courses, course)
    }

    return courses, nil
}

func (rep *Repository) GetCourse(ctx context.Context, id int) (models.Courses, error) {
    query := "SELECT * FROM graduate_work.get_course($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var course models.Courses
    err := row.Scan(
        &course.Name,
        &course.Direction,
        &course.NumOfClasses,
        &course.ClassTime,
        &course.WeekDays,
        &course.FirstClassDate,
        &course.LastClassDate,
        &course.Price,
        &course.Info,
    )
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return models.Courses{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Courses{}, nil
    }

    return course, nil
}

func (rep *Repository) GetStudents(ctx context.Context) ([]models.Students, error) {
    query := "SELECT * FROM graduate_work.get_students()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return []models.Students{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Students{}, nil
    }

    var students []models.Students
    for rows.Next() {
        var student models.Students
        err = rows.Scan(
            &student.Name,
            &student.Surname,
            &student.Patronymic,
            &student.Email,
            &student.Phone,
            &student.Comment,
            &student.Payment,
            &student.DateOfPayment,
            &student.Course,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
            return []models.Students{}, fmt.Errorf(DBError)
        }

        students = append(students, student)
    }

    return students, err
}

func (rep *Repository) GetStudent(ctx context.Context, id int) (models.Students, error) {
    query := "SELECT * FROM graduate_work.get_student($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var student models.Students
    err := row.Scan(
        &student.Name,
        &student.Surname,
        &student.Patronymic,
        &student.Email,
        &student.Phone,
        &student.Comment,
        &student.Payment,
        &student.DateOfPayment,
        &student.Course,
    )
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return models.Students{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Students{}, nil
    }

    return student, nil
}

func (rep *Repository) CreateStudent(ctx context.Context, course int, name, surname, patronymic, email, phone, comment string) error {
    query := `SELECT graduate_work.create_student($1, $2, $3, $4, $5, $6, $7, $8, $9)`

    _, err := rep.DB.ExecContext(ctx, query, course, name, surname, patronymic, email, phone, comment, false, time.Unix(0, 0))
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) UpdateCourse(ctx context.Context, id, direction, numOfClasses, classTime int,
    name, weekDays, info string, firstClassDate, lastClassDate time.Time, price float64) error {

    query := `SELECT graduate_work.update_course($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

    _, err := rep.DB.ExecContext(ctx, query, id, name, direction, numOfClasses, classTime, weekDays, firstClassDate, lastClassDate, price, info)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) DeleteCourse(ctx context.Context, id int) error {
    query := `SELECT graduate_work.delete_course($1)`

    _, err := rep.DB.ExecContext(ctx, query, id)
    if err != nil {
        log.Printf("%s: %s: %s", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}
