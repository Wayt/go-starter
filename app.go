package main

import (
	"github.com/spf13/cobra"
	"github.com/wayt/go-starter/cmd"
)

func main() {

	rootCmd := &cobra.Command{Use: "starter"}

	rootCmd.AddCommand(cmd.Serve, cmd.Migrate)
	rootCmd.Execute()
}
