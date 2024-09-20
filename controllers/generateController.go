package controllers

import (
	"fmt"
	"os"
)

func generateControllerFile(modelName string) {
	controllerFile := fmt.Sprintf("./controllers/%s_controller.go", modelName)
    file, _ := os.Create(controllerFile)
    defer file.Close()

    file.WriteString(fmt.Sprintf("package controllers\n\nimport (\n\t\"net/http\"\n\t\"github.com/gin-gonic/gin\"\n\t\"../models\"\n)\n\n"))
    
    // Generate CRUD functions
    file.WriteString(fmt.Sprintf("func Get%s(c *gin.Context) {\n\t// Get all %s\n\tvar items []models.%s\n\t// Fetch from DB (error handling to be added)\n\tc.JSON(http.StatusOK, items)\n}\n\n", modelName, modelName, modelName))
    
    file.WriteString(fmt.Sprintf("func Create%s(c *gin.Context) {\n\t// Create %s\n\tvar item models.%s\n\t// Bind request body (error handling to be added)\n\tc.BindJSON(&item)\n\t// Save to DB\n\tc.JSON(http.StatusOK, item)\n}\n\n", modelName, modelName, modelName))
    
    file.WriteString(fmt.Sprintf("func Update%s(c *gin.Context) {\n\t// Update %s\n\tvar item models.%s\n\t// Fetch existing item and update (error handling to be added)\n\tc.JSON(http.StatusOK, item)\n}\n\n", modelName, modelName, modelName))
    
    file.WriteString(fmt.Sprintf("func Delete%s(c *gin.Context) {\n\t// Delete %s\n\tvar item models.%s\n\t// Perform deletion (error handling to be added)\n\tc.JSON(http.StatusOK, item)\n}\n\n", modelName, modelName, modelName))

    fmt.Printf("Controller for %s created successfully!\n", modelName)
}