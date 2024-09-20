// controllers/generateController.go
package controllers

import (
	"fmt"
	"os"
	"strings"
)

func GenerateControllerFile(projectName, modelName string) {
	controllerFile := fmt.Sprintf("%s/controllers/%s_controller.go", projectName, strings.ToLower(modelName))
	file, _ := os.Create(controllerFile)
	defer file.Close()

	content := fmt.Sprintf(`package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../%s/models"
)

func Get%s(c *gin.Context) {
	var items []models.%s
	c.JSON(http.StatusOK, items)
}

func Create%s(c *gin.Context) {
	var item models.%s
	c.BindJSON(&item)
	c.JSON(http.StatusOK, item)
}

func Update%s(c *gin.Context) {
	var item models.%s
	c.JSON(http.StatusOK, item)
}

func Delete%s(c *gin.Context) {
	var item models.%s
	c.JSON(http.StatusOK, item)
}
`, projectName, modelName, modelName, modelName, modelName, modelName, modelName, modelName, modelName)

	file.WriteString(content)
	fmt.Printf("Controller for %s created successfully!\n", modelName)
}