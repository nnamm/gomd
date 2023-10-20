package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const (
	mdTemplate = `---
title: 
description: 
summary: 
categories: [""]
tags: [""]
date: {{.Date}}
lastmod: {{.Date}}
slug: 
---


`
)

var (
	outputDir string
	date      string
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: gomd [1-3 digit number]")
		os.Exit(1)
	}

	input := args[0]
	if !isInvalidInput(input) {
		fmt.Println("Invalid input. Please provide a valid 1-3 digit number.")
		os.Exit(1)
	}

	outputDir = generateOutputDir(input)
	date = getCurrentDate()

	err := createDirectory(outputDir)
	if err != nil {
		fmt.Println("Error: Failed to create directory: ", err)
	}

	err = createMarkdownFile()
	if err != nil {
		fmt.Println("Error: Failed to create markdown file: ", err)
		os.Exit(1)
	}
}

func isInvalidInput(input string) bool {
	if len(input) < 1 || len(input) > 3 {
		return false
	}
	for _, c := range input {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func generateOutputDir(input string) string {
	t := time.Now()
	year, month, day := t.Date()
	yearStr := fmt.Sprintf("%02d", year%100)
	return fmt.Sprintf("%03s_%02s%02d%02d", input, yearStr, month, day)
}

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func createDirectory(dirName string) error {
	return os.MkdirAll(dirName, 0755)
}

func createMarkdownFile() error {
	tmpl, err := template.New("markdown").Parse(mdTemplate)
	if err != nil {
		return err
	}

	fileName := filepath.Join(outputDir, "index.md")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		Date string
	}{Date: date}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
