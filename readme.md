# Ginister

Ginister is a Go CLI tool that helps you quickly scaffold a fully functioning backend application using the Gin framework, GORM, and MySQL. It automates the process of creating a scalable project structure, setting up a database, generating models, and implementing basic CRUD APIs.

## Features

- Automatic project setup with **Gin**, **GORM**, and **MySQL**
- Scalable folder structure generation
- Dynamic model generation with GORM support
- Pre-built CRUD operations for each model

## Prerequisites

- Go 1.18+
- MySQL

## Installation

Clone the repository and build the CLI tool:

```bash
git clone https://github.com/Utkarsh4517/ginister.git
cd ginister
go build -o ginister
```

## Running the CLI

-   ```bash
    ./ginister
    ```

- Follow the prompts!

- `go run main.go`

## Project structure explanation

```
<project_name>/
├── config/          # Database configurations
├── controllers/     # Controller files with CRUD logic
├── models/          # GORM models
├── routes/          # API route management
├── main.go          # Entry point

```