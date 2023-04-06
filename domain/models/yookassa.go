package models

import "gopkg.in/guregu/null.v4"

type CreatePayment struct {
    Amount          Amount              `json:"amount"`
    Capture         null.Bool           `json:"capture"`
    Confirmation    Confirmation        `json:"confirmation"`
    Description     null.String         `json:"description"`
    Metadata        map[string]string   `json:"metadata"`
}

type Amount struct {
    Value           null.Float          `json:"value"`
    Currency        null.String         `json:"currency"`
}

type Confirmation struct {
    Type            null.String         `json:"type"`
    ReturnUrl       null.String         `json:"return_url"`
    ConfirmationUrl null.String         `json:"confirmation_url"`
}

type GetPayment struct {
    Id              null.String         `json:"id"`
    Status          null.String         `json:"status"`
    Confirmation    Confirmation        `json:"confirmation"`
    Metadata        map[string]string   `json:"metadata"`
}
