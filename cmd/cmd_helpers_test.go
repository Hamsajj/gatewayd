package cmd

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// executeCommandC executes a cobra command and returns the command, output, and error.
// Taken from https://github.com/spf13/cobra/blob/0c72800b8dba637092b57a955ecee75949e79a73/command_test.go#L48.
func executeCommandC(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()

	return buf.String(), err
}

// mustPullPlugin pulls the gatewayd-plugin-cache plugin and returns the path to the archive.
func mustPullPlugin() (string, error) {
	pluginURL := "github.com/gatewayd-io/gatewayd-plugin-cache@v0.2.10"
	fileName := "./gatewayd-plugin-cache-linux-amd64-v0.2.10.tar.gz"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, err := executeCommandC(rootCmd, "plugin", "install", "--pull-only", pluginURL)
		if err != nil {
			return "", err
		}
	}

	return filepath.Abs(fileName) //nolint:wrapcheck
}
