package main

import (
	"errors"

	"github.com/hauntedness/panics"
)

func main() {
	panics.Ignore(errors.New("err"))
}
