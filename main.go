package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	dirFormat  = "060102"
	dateFormat = "2006-01-02"
)

type execDate struct {
	Date string
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error: Invalid number of arguments")
		os.Exit(1)
	}

	num, err := strconv.Atoi(args[1])
	if err != nil || num < 1 || num > 999 {
		fmt.Println("Error: Argument is not a valid number between 1 to 999")
		os.Exit(1)
	}

	now := time.Now()
	dirName := fmt.Sprintf("%03d_%s", num, now.Format(dirFormat))
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error: Failed to create directory")
		os.Exit(1)
	}

	md, err := os.Create(filepath.Join(dirName, "index.md"))
	if err != nil {
		fmt.Println("Error: Failed to create markdown file")
		os.Exit(1)
	}
	defer md.Close()

	tmpl := template.Must(template.New("markdown").Parse(mdTemplate))
	if err != nil {
		fmt.Println("Error: Failed to parse template")
		os.Exit(1)
	}
	ed := execDate{Date: now.Format(dateFormat)}
	err = tmpl.Execute(md, ed)
	if err != nil {
		fmt.Println("Error: Failed to execute template")
		os.Exit(1)
	}
}
