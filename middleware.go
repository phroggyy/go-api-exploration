package main

import (
    "github.com/phroggyy/apitest/persistence"
)

type Controller struct {
    session    *persistence.DB
}

func NewController(db *persistence.DB) *Controller {
    return &Controller{session:db}
}