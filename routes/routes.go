// routes/routes.go
package routes

import (
	"fmt"
	"os"
	"strings"
)

func GenerateRoutesFile(projectName string, modelNames []string) {
	routesFile := fmt.Sprintf("%s/routes/routes.go", projectName)
	file, _ := os.Create(routesFile)
	defer file.Close()

	content := `package routes

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

`

	for _, modelName := range modelNames {
		content += fmt.Sprintf(`	router.GET("/%s", controllers.Get%s)
	router.POST("/%s", controllers.Create%s)
	router.PUT("/%s/:id", controllers.Update%s)
	router.DELETE("/%s/:id", controllers.Delete%s)

`, strings.ToLower(modelName), modelName, strings.ToLower(modelName), modelName,
			strings.ToLower(modelName), modelName, strings.ToLower(modelName), modelName)
	}

	content += `	return router
}
`

	file.WriteString(content)
	fmt.Println("Routes file created successfully!")
}
