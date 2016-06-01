package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wayt/go-starter/controllers"
	"log"
)

// Serve command
var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run sql migration",
	Long:  "Run sql migration",
	Run: func(cmd *cobra.Command, args []string) {

		typ := "mysql"
		dsn := "root:toto42@tcp(192.168.99.100:3306)/starter?charset=utf8&parseTime=True"

		controllers.InitDB(typ, dsn)
		if err := controllers.AutoMigrate(); err != nil {
			log.Fatalln(err)
		}
	},
}
