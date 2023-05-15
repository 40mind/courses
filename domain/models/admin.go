package models

import (
    "gopkg.in/guregu/null.v4"
)

type Admin struct {
    Id              null.Int    `json:"id" db:"id"`
    Login           null.String `json:"login" db:"login"`
    Password        null.String `json:"password" db:"password"`
}

type Editor struct {
    Id              null.Int    `json:"id" db:"id"`
    Login           null.String `json:"login" db:"login"`
    Password        null.String `json:"password" db:"password"`
    Courses         []int64     `json:"courses" db:"courses"`
}
