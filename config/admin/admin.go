package admin

import (
	_ "errors"
	_ "fmt"
	_ "strconv"
	_ "strings"
	"github.com/qor/qor"
	"github.com/qor/admin"
	"qor-test/app/models"
	"qor-test/db"
)

var Admin *admin.Admin

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB.Set("publish:draft_mode", true)})
	Admin.SetSiteName("Qor Demo")

	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

	// Add Product
	Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})
}

