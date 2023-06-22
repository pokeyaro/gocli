package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pokeyaro/gocli/config"
	"github.com/spf13/cobra"
)

const (
	cmdName = "gocli"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// version subcommand
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(cmdName + " version information: \n")
		fmt.Printf("  Version: %s\n", config.Version)
		fmt.Printf("  Git Commit: %s\n", config.GitCommit)
		fmt.Printf("  Build Time: %s\n", config.BuildTime)
		fmt.Printf("  Go version: %s\n", runtime.Version())
		fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		return
	},
}

// start subcommand
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the project",
	Run: func(cmd *cobra.Command, args []string) {
		// Logical processing of actually starting the project
		fmt.Println("Starting the project...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number")
	rootCmd.Version = config.Version
	rootCmd.SetVersionTemplate(fmt.Sprintf(`{{with .Name}}{{printf "%%s" .}}{{end}} {{printf "%%s\n" .Version}}`))
}
