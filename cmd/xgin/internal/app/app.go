package app

import "github.com/spf13/cobra"

// 初始化 app
// 区分从 ddl 初始化和只有名称的初始化

// CmdApp represents the new command.
var CmdApp = &cobra.Command{
	Use:   "app",
	Short: "Create a app template",
	Long:  "Create a app project using the template. Example: xgin app helloworld",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

}
