package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the xgin tools",
	Long:  "Upgrade the xgin tools. Example: xgin upgrade",
	Run:   run,
}

// Run upgrade the kratos tools.
func run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"github.com/dabao-zhao/xgin@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
