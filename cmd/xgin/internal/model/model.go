package model

import "github.com/spf13/cobra"

// 根据 ddl 初始化 model

// CmdModel represents the new command.
var CmdModel = &cobra.Command{
	Use:   "model",
	Short: "Create a model template",
	Long:  "Create a model project using the template. Example: xgin model helloworld",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

}
