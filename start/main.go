package start

import (
    "fmt"
    "os"
)

func CreateMainFile(projectDir string) error {
    content := `package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Run(":8080")
}
`

    filePath := fmt.Sprintf("%s/main.go", projectDir)
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    return err
}