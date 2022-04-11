package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// RootCmd является командой, вызываемой по-умолчанию
var RootCmd = &cobra.Command{
	Use:   "contractor_panel",
	Short: "`contractor_panel` microservice provides contractor panel functionality",
	Long:  "`contractor_panel` microservice provides contractor panel functionality",
}

// Execute вызывается единожды из main.main().
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		viper.AutomaticEnv()
	})
}
