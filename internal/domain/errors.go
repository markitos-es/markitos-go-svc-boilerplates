package domain

import "errors"

var ErrBoilerplateNotFound error = errors.New("boilerplate not found")
var ErrBoilerplateBadRequest error = errors.New("bad request")
