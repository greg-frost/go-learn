package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(" \n[ БАЗА ]\n ")

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	workingDir = filepath.ToSlash(workingDir)

	targetDir := "base"
	targetDir = filepath.ToSlash(targetDir)

	fmt.Println(workingDir)
	fmt.Println(targetDir)

	workingPath := strings.Split(workingDir, "/")
	targetPath := strings.Split(targetDir, "/")
	fmt.Println(workingPath)
	fmt.Println(targetPath)

	var p1, p2 int
	n1, n2 := len(workingPath), len(targetPath)
	for p1 < n1 && p2 < n2 {
		if workingPath[p1] == targetPath[p2] {
			targetPath = targetPath[1:]
			p2++
		}
		p1++
	}

	fmt.Println(workingPath)
	fmt.Println(targetPath)

	if len(targetPath) > 0 {
		workingDir = filepath.Join(workingDir, filepath.Join(targetPath...))
	}

	fmt.Println(workingDir)

	file, err := os.Open(workingDir + "/main.go")
	fmt.Println(file, err)
}
