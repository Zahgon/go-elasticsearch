package commands

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	IndexName string

	tWidth int
)

var rootCmd = &cobra.Command{
	Use:   "xkcd",
	Short: "xkcd allows you to index and search xkcd.com",
	Long:  "xkcd indexes comic metadata from xkcd.com into Elasticsearch and searches it from the command line or a local web interface.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&IndexName, "index", "i", "xkcd", "Index name")
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
}

func Execute() { _ = "STUB: not implemented"; return }
