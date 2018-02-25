package business

import (
	"github.com/Qiaorui/zooli/controllers"
	"github.com/Qiaorui/zooli/models"
)

func GetCurrentStore(c *controllers.BaseController) (models.Store, error) {
	manager := c.GetSession("user")
	id := manager.(models.User).StoreID
	var store models.Store
	store = models.FindStoreByID(uint(id))

	return store, nil
}
