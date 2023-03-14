package main

import (
	"log"

	"github.com/spf13/cobra"

	"xgin/internal/app"
	"xgin/internal/model"
	xnew "xgin/internal/new"
	"xgin/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "xgin",
	Short:   "xgin: A toolkit for Gin.",
	Long:    `xgin: A toolkit for Gin.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(model.CmdModel)
	rootCmd.AddCommand(xnew.CmdNew)
	rootCmd.AddCommand(app.CmdApp)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
