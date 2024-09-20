// models/generateModel.go
package models

import (
	"fmt"
	"os"
	"strings"
)

func GenerateModelFile(projectName, modelName string, fields []string) {
	modelFile := fmt.Sprintf("%s/models/%s.go", projectName, strings.ToLower(modelName))
	file, _ := os.Create(modelFile)
	defer file.Close()

	content := fmt.Sprintf("package models\n\nimport \"gorm.io/gorm\"\n\ntype %s struct {\n", modelName)
	for _, field := range fields {
		content += "\t" + field + "\n"
	}
	content += "\tgorm.Model\n}\n"

	file.WriteString(content)
	fmt.Printf("Model for %s created successfully!\n", modelName)
}
