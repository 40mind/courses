package repository

import (
    "context"
    "courses/domain/models"
    "database/sql"
    "fmt"
    "github.com/jimlawless/whereami"
    "log"
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
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) GetDirections(ctx context.Context) ([]models.Direction, error) {
    query := "SELECT * FROM graduate_work.get_directions()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return []models.Direction{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Direction{}, nil
    }

    var directions []models.Direction
    for rows.Next() {
        var direction models.Direction
        err = rows.Scan(
            &direction.Id,
            &direction.Name,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
            return []models.Direction{}, fmt.Errorf(DBError)
        }

        directions = append(directions, direction)
    }

    return directions, err
}

func (rep *Repository) GetDirection(ctx context.Context, id int) (models.Direction, error) {
    query := "SELECT * FROM graduate_work.get_direction_by_id($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var direction models.Direction
    err := row.Scan(
        &direction.Id,
        &direction.Name,
    )
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return models.Direction{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Direction{}, nil
    }

    return direction, err
}

func (rep *Repository) UpdateDirection(ctx context.Context, direction models.Direction) error {
    query := "SELECT * FROM graduate_work.update_direction($1, $2)"

    _, err := rep.DB.ExecContext(ctx, query, direction.Id, direction.Name)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) DeleteDirection(ctx context.Context, id int) error {
    query := "SELECT * FROM graduate_work.delete_direction($1)"

    _, err := rep.DB.ExecContext(ctx, query, id)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) CreateCourse(ctx context.Context, course models.Course) error {
    query := "SELECT * FROM graduate_work.create_course($1, $2, $3, $4, $5, $6, $7, $8, $9)"

    _, err := rep.DB.ExecContext(ctx, query, course.Name, course.NumOfClasses, course.ClassTime, course.WeekDays,
        course.FirstClassDate, course.LastClassDate, course.Price, course.Info, course.DirectionId)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) GetCourses(ctx context.Context) ([]models.Course, error) {
    query := "SELECT * FROM graduate_work.get_courses()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return []models.Course{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Course{}, nil
    }

    var courses []models.Course
    for rows.Next() {
        var course models.Course
        err = rows.Scan(
            &course.Id,
            &course.Name,
            &course.NumOfClasses,
            &course.ClassTime,
            &course.WeekDays,
            &course.FirstClassDate,
            &course.LastClassDate,
            &course.Price,
            &course.Info,
            &course.DirectionId,
            &course.DirectionName,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
            return []models.Course{}, fmt.Errorf(DBError)
        }

        courses = append(courses, course)
    }

    return courses, nil
}

func (rep *Repository) GetCourse(ctx context.Context, id int) (models.Course, error) {
    query := "SELECT * FROM graduate_work.get_course_by_id($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var course models.Course
    err := row.Scan(
        &course.Id,
        &course.Name,
        &course.NumOfClasses,
        &course.ClassTime,
        &course.WeekDays,
        &course.FirstClassDate,
        &course.LastClassDate,
        &course.Price,
        &course.Info,
        &course.DirectionId,
        &course.DirectionName,
    )
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return models.Course{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Course{}, nil
    }

    return course, nil
}

func (rep *Repository) UpdateCourse(ctx context.Context, course models.Course) error {
    query := "SELECT * FROM graduate_work.update_course($1, $2, $3, $4, $5, $6, $7, $8, $9)"

    _, err := rep.DB.ExecContext(ctx, query, course.Id, course.Name, course.NumOfClasses, course.ClassTime, course.WeekDays,
        course.FirstClassDate, course.LastClassDate, course.Price, course.Info)

    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) DeleteCourse(ctx context.Context, id int) error {
    query := `SELECT graduate_work.delete_course($1)`

    _, err := rep.DB.ExecContext(ctx, query, id)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) CreateStudent(ctx context.Context, student models.Student) error {
    query := "SELECT * FROM graduate_work.create_student($1, $2, $3, $4, $5, $6, $7, $8, $9)"

    _, err := rep.DB.ExecContext(ctx, query, student.Name, student.Surname, student.Patronymic, student.Email,
        student.Phone, student.Comment, false, nil, student.CourseId)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) GetStudents(ctx context.Context) ([]models.Student, error) {
    query := "SELECT * FROM graduate_work.get_students()"

    rows, err := rep.DB.QueryContext(ctx, query)
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return []models.Student{}, fmt.Errorf(DBError)
    }
    defer rows.Close()

    if err == sql.ErrNoRows {
        return []models.Student{}, nil
    }

    var students []models.Student
    for rows.Next() {
        var student models.Student
        err = rows.Scan(
            &student.Id,
            &student.Name,
            &student.Surname,
            &student.Patronymic,
            &student.Email,
            &student.Phone,
            &student.Comment,
            &student.Payment,
            &student.DateOfPayment,
            &student.CourseId,
            &student.CourseName,
        )
        if err != nil && err != sql.ErrNoRows {
            log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
            return []models.Student{}, fmt.Errorf(DBError)
        }

        students = append(students, student)
    }

    return students, err
}

func (rep *Repository) GetStudent(ctx context.Context, id int) (models.Student, error) {
    query := "SELECT * FROM graduate_work.get_student($1)"

    row := rep.DB.QueryRowContext(ctx, query, id)

    var student models.Student
    err := row.Scan(
        &student.Id,
        &student.Name,
        &student.Surname,
        &student.Patronymic,
        &student.Email,
        &student.Phone,
        &student.Comment,
        &student.Payment,
        &student.DateOfPayment,
        &student.CourseId,
        &student.CourseName,
    )
    if err != nil && err != sql.ErrNoRows {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return models.Student{}, fmt.Errorf(DBError)
    }

    if err == sql.ErrNoRows {
        return models.Student{}, nil
    }

    return student, nil
}

func (rep *Repository) UpdateStudent(ctx context.Context, student models.Student) error {
    query := "SELECT * FROM graduate_work.update_student($1, $2, $3, $4, $5, $6, $7, $8, $9)"

    _, err := rep.DB.ExecContext(ctx, query, student.Id, student.Name, student.Surname, student.Patronymic, student.Email,
        student.Phone, student.Comment, student.Payment, student.DateOfPayment)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}

func (rep *Repository) DeleteStudent(ctx context.Context, id int) error {
    query := `SELECT graduate_work.delete_student($1)`

    _, err := rep.DB.ExecContext(ctx, query, id)
    if err != nil {
        log.Printf("%s: %s: %s\n", DBError, err.Error(), whereami.WhereAmI())
        return fmt.Errorf(DBError)
    }

    return nil
}
