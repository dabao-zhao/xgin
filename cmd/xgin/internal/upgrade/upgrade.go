package upgrade

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
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
	err := goInstall(
		"github.com/dabao-zhao/xgin/v1@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}

// goInstall go get path.
func goInstall(path ...string) error {
	for _, p := range path {
		fmt.Printf("go get -u %s\n", p)
		cmd := exec.Command("go", "get", "-u", p)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
