package public

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
	"github.com/pkg/errors"
)

func GetCurrentStore(c *controllers.BaseController) (models.Store, error) {
	manager := c.GetSession("user")
	id := manager.(models.User).StoreID
	store := models.FindStoreByID(uint(id))
	if !store.Exists() {
		return store, errors.New("User has not store assigned")
	}

	return store, nil
}
