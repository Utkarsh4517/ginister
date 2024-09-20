/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Utkarsh4517/ginister/config"
	"Utkarsh4517/ginister/controllers"
	"Utkarsh4517/ginister/models"
	"Utkarsh4517/ginister/routes"
	"bufio"
	"fmt"
	"os"
	"strings"
	"Utkarsh4517/ginister/docker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ginister",
	Short: "CLI to generate Go-Gin projects with MySQL and GORM",
	Run:   runGenerator,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runGenerator(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	createProject(projectName, reader)
}

func createProject(projectName string, reader *bufio.Reader) {
    fmt.Println("Creating project:", projectName, "powered by ginister")

    createProjectStructure(projectName)

    config.GenerateConfigFile(projectName)
    err := docker.CreateDockerfile(projectName)
    if err != nil {
        fmt.Printf("Error creating Dockerfile: %v\n", err)
    }

    var modelNames []string

    for {
        fmt.Print("Enter a model name (or press enter to finish): ")
        modelName, _ := reader.ReadString('\n')
        modelName = strings.TrimSpace(modelName)

        if modelName == "" {
            break
        }

        modelNames = append(modelNames, modelName)
        fields := getModelFields(reader)
        models.GenerateModelFile(projectName, modelName, fields)
        controllers.GenerateControllerFile(projectName, modelName)
    }
    routes.GenerateRoutesFile(projectName, modelNames)
    fmt.Println("Project setup complete!")
}


func createProjectStructure(projectName string) {
	directories := []string{"", "/config", "/controllers", "/models", "/routes"}
	for _, dir := range directories {
		err := os.MkdirAll(projectName+dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", projectName+dir, err)
		}
	}
}

func getModelFields(reader *bufio.Reader) []string {
	var fields []string
	for {
		fmt.Print("Enter a field (name type) or press enter to finish: ")
		field, _ := reader.ReadString('\n')
		field = strings.TrimSpace(field)

		if field == "" {
			break
		}

		fields = append(fields, field)
	}
	return fields
}