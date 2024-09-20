package models
import (
	"fmt"
	"os"
)
func generateModelFile(modelName string, fields []string) {
	modelFile := fmt.Sprintf("./models/%s.go", modelName)
    file, _ := os.Create(modelFile)
    defer file.Close()

    file.WriteString(fmt.Sprintf("package models\n\nimport \"gorm.io/gorm\"\n\ntype %s struct {\n", modelName))
    for _, field := range fields {
        file.WriteString("\t" + field + "\n")
    }
    file.WriteString("\tgorm.Model\n}\n")
    
    fmt.Printf("Model for %s created successfully!\n", modelName)
}