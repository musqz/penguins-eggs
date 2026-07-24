package cmd

import (
	"os"

	"coa/pkg/builder"
	"coa/pkg/distro"
	"coa/pkg/utils"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [distro]",
	Short: "Compile binaries and generate native distribution packages (.deb, PKGBUILD/pkg.tar.zst)",
	Long: `The 'build' command is the integrated packaging tool for the coa/oa ecosystem.
It orchestrates the full compilation of both the C-native engine (oa) and the Go-based orchestrator (coa), triggers the automatic generation of documentation and shell completions, and finally packages everything into native distribution formats like .deb (Debian/Ubuntu) or PKGBUILD / .pkg.tar.zst (Arch Linux).`,
	Example: `  # Compile the ecosystem and generate native package for host distro
  coa tools build
  # Force building Arch Linux package (.pkg.tar.zst) on any distribution
  coa tools build arch`,
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() == 0 {
			utils.Fatal(" Execution aborted. Do NOT run 'coa tools build' with sudo!")
			utils.LogNormal("Compilation must be run as a normal user to avoid " +
				"creating root-owned files and packages in your workspace.")
			os.Exit(1)
		}

		myDistro := distro.NewDistro()
		target := ""
		if len(args) > 0 {
			target = args[0]
		}
		builder.HandleBuild(myDistro, target)
	},
}

func init() {
	toolsCmd.AddCommand(buildCmd)
}
