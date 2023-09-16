package cmd

import (
	"github.com/spf13/cobra"
)

var (
	vaultendpoint  string
	vaulttoken     string
	vaultcapath    string
	vaultcacert    string
	vaultcert      string
	vaultkey       string
	vaultformat    string
	vaultnamespace string
)

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage Vault context profiles ",
	Long:  `Manage Vault context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
	},
	//Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// fmt.Println("vault called")

	// },
}

func init() {
	vaultCmd.AddCommand(vaultCreateCmd)
	vaultCmd.AddCommand(deleteVaultCmd)
	//vaultCmd.AddCommand(renameCmd)
	vaultCmd.AddCommand(selectVaultCmd)
	vaultCmd.AddCommand(vaultUpdateCmd)
	vaultCmd.AddCommand(listVaultCmd)

}
