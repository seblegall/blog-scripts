package cmd

import (
	"io"
	"os"

	"github.com/seblegall/blog-scripts/logrus-cobra/pkg"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var v string

var rootCmd = &cobra.Command{
	Use:   "mycmd",
	Short: "mycmd is a test project to integrate logrus",
}

var subCmd = &cobra.Command{
	Use:   "subcmd",
	Short: "subcmd is a subcommand of the main cmd",
	Run: func(cmd *cobra.Command, args []string) {
		runSubCmd()
	},
}

//NewMyCmd return the root cobra command
func NewMyCmd() *cobra.Command {

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, v); err != nil {
			return err
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", logrus.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	rootCmd.AddCommand(NewSubCmd())

	return rootCmd
}

//NewSubCmd return the a sub cobra command
func NewSubCmd() *cobra.Command {
	return subCmd
}

func setUpLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}

func runSubCmd() {
	pkg.MyFunc()
}
