package main

import (
    "fmt"
)

type Api struct {  }

func (a *Api) Get(min float64, max float64) {
    fmt.Printf("GET called with min: %f and max: %f\n", min, max)
}

func (a *Api) Post(user string, body string) {
    fmt.Printf("POST: new pweet %s:%s\n", user, body)
}
