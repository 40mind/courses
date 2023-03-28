package models

import (
    "gopkg.in/guregu/null.v4"
)

type AdminInfo struct {
    Direction []Direction `json:"directions"`
    Course    []Course    `json:"courses"`
    Student   []Student   `json:"students"`
    Admins    []Admin     `json:"admins"`
}

type Admin struct {
    Id              null.Int    `json:"id" db:"id"`
    Login           null.String `json:"login" db:"login"`
    Password        null.String `json:"password" db:"password"`
}
