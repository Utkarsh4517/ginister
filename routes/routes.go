package routes

import (
	"fmt"
	"os"
)

func generateRoutesFile(modelName string) { 
	routesFile := "./routes/routes.go"

	if _, err := os.Stat(routesFile); os.IsNotExist(err) {
        file, _ := os.Create(routesFile)
        defer file.Close()

        file.WriteString("package routes\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"../controllers\"\n)\n\n")
        file.WriteString("func SetupRouter() *gin.Engine {\n\trouter := gin.Default()\n\n")
        file.WriteString("\treturn router\n}")
    }

	file, _ := os.OpenFile(routesFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    defer file.Close()

    // Register routes for the current model
    file.WriteString(fmt.Sprintf("\trouter.GET(\"/%s\", controllers.Get%s)\n", modelName, modelName))
    file.WriteString(fmt.Sprintf("\trouter.POST(\"/%s\", controllers.Create%s)\n", modelName, modelName))
    file.WriteString(fmt.Sprintf("\trouter.PUT(\"/%s/:id\", controllers.Update%s)\n", modelName, modelName))
    file.WriteString(fmt.Sprintf("\trouter.DELETE(\"/%s/:id\", controllers.Delete%s)\n", modelName, modelName))

    fmt.Printf("Routes for %s added successfully!\n", modelName)

}