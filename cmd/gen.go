/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/rjprice04/bruman/internal"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates an HTML Page from the input file name",
	Long: `Generates an HTML Page from the input file name

bruman gen -f out.json `,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file");
		if err != nil {
			log.Fatal("Error getting file name", err)
		}
		bruFile := internal.NewBruFile(file)
		internal.CreateHtml(bruFile)


		


		fmt.Println(bruFile.Summary.PassedTests)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")
	genCmd.Flags().String("file", "bru.json", "The path to the json file output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
