/*
Copyright © 2024 Matheus Lopes <matlopes1999@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "CLI tool to run stress tests",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "", "URL do serviço a ser testado (Required)")
	rootCmd.Flags().Int("requests", 0, "Número total de requests (Required)")
	rootCmd.Flags().Int("concurrency", 0, "Número de chamadas simultâneas (Required)")
	if err := rootCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired("requests"); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired("concurrency"); err != nil {
		panic(err)
	}
}
