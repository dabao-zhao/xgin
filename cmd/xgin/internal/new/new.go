package new

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/base"
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
	repoURL string
	branch  string
	timeout string
)

func init() {
	repoURL = "https://github.com/dabao-zhao/xgin-layout.git"
	timeout = "60s"

	CmdNew.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	CmdNew.Flags().StringVarP(&branch, "branch", "b", branch, "repo branch")
	CmdNew.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	name := ""
	if len(args) == 0 {
		log.Fatalln("need project name")
	} else {
		name = args[0]
	}

	done := make(chan error, 1)
	go func() {
		done <- create(ctx, wd, name, repoURL, branch)
		return
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			_, _ = fmt.Fprint(os.Stderr, "ERROR: project creation timed out")
			return
		}
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: failed to create project(%s)", ctx.Err().Error())
	case err = <-done:
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: Failed to create project(%s)", err.Error())
		}
	}
}

func create(ctx context.Context, wd, name, layout, branch string) error {
	to := path.Join(wd, name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		return errors.New("folder already exists")
	}

	fmt.Printf("Creating project %s, layout repo is %s, please wait a moment.\n", name, layout)
	repo := base.NewRepo(layout, branch)
	if err := repo.CopyTo(ctx, to, path.Base(name), []string{".git", ".github"}); err != nil {
		return err
	}

	fmt.Printf("Project creation succeeded %s\n", name)

	return nil
}
