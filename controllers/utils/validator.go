package controllers

import (
	"github.com/asaskevich/govalidator"
	"fmt"

	"github.com/Qiaorui/zooli/models"
	"github.com/pkg/errors"
)

func Validate(obj interface{}, validatePK ...bool) error {
	var err error
	switch obj := obj.(type) {

	default:
		fmt.Println(obj, "unexpected type %T")
	case models.User:
		if obj.ExistsUsername() {
			err = errors.New("username already exists")
		}
	case models.Role:
		if obj.ExistsName() {
			err = errors.New("role already exists")
		}

	}
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(obj)
	//prepare the error message
	/*if err != nil {
		for _, e := range err.(govalidator.Errors) {
			field := e.(govalidator.Error).Name
			errMsg += "Invalid " + field //+ "<br>"
		}
	}*/


	return err
}