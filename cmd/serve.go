package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/wayt/go-starter/api/v1"
	"github.com/wayt/go-starter/controllers"
)

// Serve command
var Serve = &cobra.Command{
	Use:   "serve",
	Short: "Run http server",
	Long:  "Run http server",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Starting...\n")

		typ := "mysql"
		dsn := "root:toto42@tcp(192.168.99.100:3306)/starter?charset=utf8&parseTime=True"

		controllers.InitDB(typ, dsn)

		r := gin.Default()

		fmt.Printf("Registering /v1 ...")
		v1.Init(r.Group("/v1"))

		r.Run(":8080")
	},
}
