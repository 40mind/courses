package models

import (
    "gopkg.in/guregu/null.v4"
)

type Direction struct {
    Id   null.Int    `json:"id" db:"id"`
    Name null.String `json:"name" db:"name"`
}

type Course struct {
    Id             null.Int       `json:"id" db:"id"`
    Name           null.String    `json:"name" db:"name"`
    NumOfClasses   null.Int       `json:"num_of_classes" db:"num_of_classes"`
    ClassTime      null.Int       `json:"class_time" db:"class_time"`
    WeekDays       null.String    `json:"week_days" db:"week_days"`
    FirstClassDate null.Time      `json:"first_class_date" db:"first_class_date"`
    LastClassDate  null.Time      `json:"last_class_date" db:"last_class_date"`
    Price          null.Float     `json:"price" db:"price"`
    Info           null.String    `json:"info" db:"info"`
    DirectionId    null.Int       `json:"direction_id" db:"direction_id"`
    DirectionName  null.String    `json:"direction_name" db:"direction_name"`
}

type Student struct {
    Id            null.Int       `json:"id" db:"id"`
    Name          null.String    `json:"name" db:"name"`
    Surname       null.String    `json:"surname" db:"surname"`
    Patronymic    null.String    `json:"patronymic" db:"patronymic"`
    Email         null.String    `json:"email" db:"email"`
    Phone         null.String    `json:"phone" db:"phone"`
    Comment       null.String    `json:"comment" db:"comment"`
    Payment       null.Bool      `json:"payment" db:"payment"`
    DateOfPayment null.Time      `json:"date_of_payment" db:"date_of_payment"`
    CourseId      null.Int       `json:"course_id" db:"course_id"`
    CourseName    null.String    `json:"course_name" db:"course_name"`
}
