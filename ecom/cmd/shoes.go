/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shoesCmd represents the shoes command
var shoesCmd = &cobra.Command{
	Use:   "shoes",
	Short: "A brief description of your command",
	Long: `
	flag:
		--instock (true/false) global flags from parent (items)

		--number (9,10,11) only for this cmd 
	args:
		args passed are used as search text

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("shoes called")

		instock_from_flag, _ := cmd.Flags().GetBool("instock")
		number_from_flag, _ := cmd.Flags().GetInt16("number")

		if number_from_flag > 10 {
			return fmt.Errorf("imsry your feet is too large: ")
		}

		fmt.Println("instock = ", instock_from_flag)
		fmt.Println("number = ", number_from_flag)

		fmt.Println("args/search-text = ", args)

		return nil
	},
}

func init() {
	itemsCmd.AddCommand(shoesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shoesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	shoesCmd.Flags().Int16P("number", "n", 9, "shoes size = 9,10,11")
}
