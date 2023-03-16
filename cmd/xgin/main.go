package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/app"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd"
	xnew "github.com/dabao-zhao/xgin/cmd/xgin/internal/new"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "xgin",
	Short:   "xgin: A toolkit for Gin.",
	Long:    `xgin: A toolkit for Gin.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(app.CmdApp)
	rootCmd.AddCommand(xnew.CmdNew)
	rootCmd.AddCommand(curd.CmdCurd)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
