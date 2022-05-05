package controller

import (
	"github.com/ralugr/filter-service/pkg/stores"
	"github.com/ralugr/filter-service/pkg/validators"
)

var Validators []validators.Base
var Store stores.Base

func AddValidator(v validators.Base) {
	Validators = append(Validators, v)
}

func RemoveValidator(index int) {
	Validators = append(Validators[:index], Validators[index+1:]...)
}

func GetValidators() []validators.Base {
	return Validators
}

func GetStore() stores.Base {
	return Store
}

func UpdateStore(newStore stores.Base) {
	Store = newStore
}
