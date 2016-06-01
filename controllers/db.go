package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import mysql driver
)

// DB object, should not be used outside controllers !
var db *gorm.DB

// InitDB create your DB object
// Support mysql by default, you can change it by importing another dialect
func InitDB(typ, dsn string) error {

	// Modify this if needed
	if typ != "mysql" {
		panic(fmt.Errorf("invalid db type `%s`, you MUST change dialect import in controllers/db.go", typ))
	}

	fmt.Printf("Opening database...\n")
	var err error
	db, err = gorm.Open(typ, dsn)
	if err != nil {
		return fmt.Errorf("db: open: %s", err)
	}

	return nil
}

func AutoMigrate() error {

	fmt.Printf("%d controller(s)\n", len(controllers))

	for _, ctl := range controllers {

		fmt.Printf("\tmigrate %s ... ", ctl.Model.TableName())
		if err := db.AutoMigrate(ctl.Model).Error; err != nil {
			fmt.Printf("error\n")
			return err
		}
		fmt.Printf("done\n")
	}

	return nil
}
