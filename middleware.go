package main

import (
    "github.com/phroggyy/go-api-exploration/persistence"
)

type Controller struct {
    session    *persistence.DB
}

func NewController(db *persistence.DB) *Controller {
    return &Controller{session:db}
}