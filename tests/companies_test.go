package test

import (
	"testing"
	"zooli/controllers/admin"

	"github.com/Qiaorui/zooli/models"
)

func TestGet(t *testing.T) {
	companiesCtrl := &admin.CompaniesController{}
	companiesCtrl.get()
	if c.Data["companies"] != models.FindCompanies() {
		t.Errorf("Get failed")
	}
}
