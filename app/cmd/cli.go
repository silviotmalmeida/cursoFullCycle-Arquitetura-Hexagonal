/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

// instanciando as variáveis a serem recebidas como parâmetros
var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("cli called")

		// iniciando o adapter de CLI para manipulação de product
		res, err := cli.Run(&productService, action, productId, productName, productPrice)
		// em caso de erro, retorna-o
		if err != nil {
			fmt.Println(err.Error())
		}
		// exibe a resposta
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// cadastrando os parâmetros a serem recebidos
	cliCmd.Flags().StringVarP(&action, "action", "a", "get", "Possible values: create, enable or disable.")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product price")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
