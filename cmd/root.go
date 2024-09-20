/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
    "bufio"
	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "ginister",
	Short: "CLI to generate Go-Gin projects with MySQL and GORM",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the project name: ")
        projectName, _ := reader.ReadString('\n')
        createProject(projectName)

	},
	
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ginister.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createProject(projectName string) {
	fmt.Println("Creating project: ", projectName, " powered by ginister")
	os.Mkdir(projectName, os.ModePerm)
	os.Mkdir(projectName + "/config", os.ModePerm)
    os.Mkdir(projectName + "/controllers", os.ModePerm)
    os.Mkdir(projectName + "/models", os.ModePerm)
    os.Mkdir(projectName + "/routes", os.ModePerm)
}

