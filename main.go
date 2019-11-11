package main

import (
	"fmt"
	"os"
)

const (
	envRepo = "GITHUB_REPOSITORY"
)

func main() {
	repo := os.Getenv(envRepo)
	fmt.Println(repo)
}
