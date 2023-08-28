package version

import (
	"errors"
	"fmt"

	"github.com/App/desktop/scripts/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	VersionCmd.AddCommand(versionUpdateCmd)
}

// Command to check the application version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version App Desktop Version",
	RunE: func(cmd *cobra.Command, args []string) error {

		packageJson, err := config.ReadConfigJson(config.PackageJsonFile)

		if err != nil {
			return err
		}

		packageJsonVersion := packageJson["version"]

		cargoToml, err := config.ReadconfigToml(config.CargoTomlFile)

		if err != nil {
			return err
		}

		cargoTomlVersion := cargoToml["package"].(map[string]any)["version"]

		tauriConfJson, err := config.ReadConfigJson(config.TauriConfFile)

		if err != nil {
			return err
		}

		tauriConfJsonVersion := tauriConfJson["package"].(map[string]any)["version"]

		if !(packageJsonVersion == cargoTomlVersion && cargoTomlVersion == tauriConfJsonVersion) {
			color.HiRed("Error:")
			fmt.Printf("package.json    -> %s\n", packageJsonVersion)
			fmt.Printf("Cargo.Toml      -> %s\n", cargoTomlVersion)
			fmt.Printf("tauri.conf.json -> %s\n", tauriConfJsonVersion)
			return errors.New("Version is inconsistent between package.json, Cargo.tom and tauri.conf.json")
		} else {
			color.HiYellow("App DESKTOP v%s", tauriConfJsonVersion)
		}

		return nil
	},
}
