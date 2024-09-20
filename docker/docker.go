package docker

import (
	"fmt"
	"os"
)

func CreateDockerfile(projectDir string) error {
	content := `FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
`

	filePath := fmt.Sprintf("%s/Dockerfile", projectDir)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
