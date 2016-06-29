package models

import (
	_ "fmt"
	_ "log"
	_ "strings"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name string
}
