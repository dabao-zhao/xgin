package new

import (
	"github.com/spf13/cobra"
)

// 初始化项目

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service project using the template. Example: xgin new helloworld",
	Run:   run,
}

var (
	serviceName string
)

func init() {
	CmdNew.Flags().StringVarP(&serviceName, "name", "n", "", "service name")
}

func run(cmd *cobra.Command, args []string) {

}
