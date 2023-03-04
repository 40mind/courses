package models

import "time"

type AdminInfo struct {
    Directions []Directions `json:"directions"`
    Courses    []Courses    `json:"courses"`
    Students   []Students   `json:"students"`
}

type Directions struct {
    Name string `json:"name" db:"name"`
}

type Courses struct {
    Name           string    `json:"name" db:"name"`
    NumOfClasses   int       `json:"num_of_classes" db:"num_of_classes"`
    ClassTime      int       `json:"class_time" db:"class_time"`
    WeekDays       string    `json:"week_days" db:"week_days"`
    FirstClassDate time.Time `json:"first_class_date" db:"first_class_date"`
    LastClassDate  time.Time `json:"last_class_date" db:"last_class_date"`
    Price          float64   `json:"price" db:"price"`
    Info           string    `json:"info" db:"info"`
    Direction      string    `json:"direction" db:"direction"`
}

type Students struct {
    Name          string    `json:"name" db:"name"`
    Surname       string    `json:"surname" db:"surname"`
    Patronymic    string    `json:"patronymic" db:"patronymic"`
    Email         string    `json:"email" db:"email"`
    Phone         string    `json:"phone" db:"phone"`
    Comment       string    `json:"comment" db:"comment"`
    Payment       bool      `json:"payment" db:"payment"`
    DateOfPayment time.Time `json:"date_of_payment" db:"date_of_payment"`
    Course        string    `json:"course" db:"course"`
}
