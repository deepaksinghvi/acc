/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ledgerCmd represents the ledger command
var ledgerCmd = &cobra.Command{
	Use:   "ledger",
	Short: "Execute \"acc ledger\" to initiate ledger module",
	Long: `This is the ledger module for the accounting project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---Accounting Ledger Module---")
	},
}

func init() {
	rootCmd.AddCommand(ledgerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ledgerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ledgerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
