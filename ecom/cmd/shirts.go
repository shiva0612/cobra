/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shirtsCmd represents the shirts command
var shirtsCmd = &cobra.Command{
	Use:   "shirts",
	Short: "",
	Long: `
		flag:
			--instock (true/false) global flags from parent (items)

			--size (s,m,l) only for this cmd 
		args:
			args passed are used as search text

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shirts called")

		instock_from_flag, _ := cmd.Flags().GetBool("instock")
		size_from_flag, _ := cmd.Flags().GetString("size")

		fmt.Println("instock = ", instock_from_flag)
		fmt.Println("size = ", size_from_flag)

		fmt.Println("args/search-text = ", args)
	},
}

func init() {
	itemsCmd.AddCommand(shirtsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shirtsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	shirtsCmd.Flags().StringP("size", "s", "s", "size = s,m,l")
}
