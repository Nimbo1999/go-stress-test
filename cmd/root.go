/*
Copyright © 2024 Matheus Lopes <matlopes1999@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"os"
	"time"

	"github.com/nimbo1999/go-stress-test/internal/client"
	"github.com/nimbo1999/go-stress-test/internal/service"
	"github.com/nimbo1999/go-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "CLI tool to run stress tests",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			panic(err)
		}
		requests, err := cmd.Flags().GetInt("requests")
		if err != nil {
			panic(err)
		}
		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			panic(err)
		}

		appHttpClient := client.NewAppHattpClient(time.Second * 10)
		service := service.NewStressTestService(appHttpClient)
		useCase := usecase.NewStressTestUseCase(url, requests, concurrency, service)
		response := useCase.Execute()
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		encoder.Encode(response)
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
