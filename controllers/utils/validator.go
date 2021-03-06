package controllers

import (
	"github.com/asaskevich/govalidator"
	"fmt"

	"github.com/Qiaorui/zooli/models"
	"github.com/pkg/errors"
	"regexp"
)

func Validate(obj interface{}, validatePK ...bool) error {
	var err error
	switch obj := obj.(type) {

	default:
		fmt.Println(obj, "unexpected type %T")
	case models.User:
		if obj.ExistsUsername() {
			err = errors.New("username already exists")
			break
		}
		role := new(models.Role)
		role.ID = obj.RoleID
		if !role.Exists() {
			err = errors.New("Invalid Role")
		}
	case models.Role:
		if obj.ExistsName() {
			err = errors.New("role already exists")
		}
	case models.Company:
		if obj.ExistsName() {
			err = errors.New("company name already exists")
		}
	case models.Store:
		for _, s := range obj.Services {
			if !s.Exists() {
				err = errors.New("Invalid Service")
			}
		}
		for _, m := range obj.Managers {
			if !m.Exists() {
				err = errors.New("Invalid Manager")
			}
		}

	case models.Product:
		found := false
		availableServices := models.FindStoreByID(obj.StoreID).Services
		for _, s := range availableServices {
			if s.ID == obj.ServiceID {
				found = true
				break
			}
		}
		if !found {
			err = errors.New("product provides a service that is not available for its store")
		}
	case models.Order:
		//
	case models.Customer:
		//
	case models.Vehicle:
		if obj.ExistsPlate() {
			err = errors.New("vehicle plate already been registered")
		}

	}

	if err != nil {
		return err
	}
	// TODO: a temporary fix, should be changed later
	govalidator.TagMap["phone"] = govalidator.Validator(func(str string) bool {
		match, _ := regexp.MatchString("^[\\d-+]+$", str)
		return match
	})

	_, err = govalidator.ValidateStruct(obj)

	return err
}