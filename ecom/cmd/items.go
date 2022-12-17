/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "A brief description of your command",
	Long: `
		flags:
			--instock (for this cmd and its child = shirts and shoes)

		args:
			args passed are used as search text
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("items called")
		instock_from_flag, _ := cmd.Flags().GetBool("instock")
		toggle_from_flag, _ := cmd.Flags().GetBool("toggle")

		fmt.Println("instock = ", instock_from_flag)
		fmt.Println("toggle = ", toggle_from_flag)
		fmt.Println("args/search-text = ", args)

	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	itemsCmd.PersistentFlags().Bool("instock", true, "instock defualt = true")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	itemsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
