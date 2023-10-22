package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"
)

const mdTemplate = `---
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

type TmplData struct {
	Date string
}

func GenerateContent(arg string) error {
	if !isValidArg(arg) {
		return fmt.Errorf("invalid argument, expected 1-3 digit number")
	}

	formattedArg := fmt.Sprintf("%03s", arg)
	now := time.Now().Format("2006-01-02")
	dirName := fmt.Sprintf("%s_%s", formattedArg, now[2:4]+now[5:7]+now[8:10])

	if err := os.Mkdir(dirName, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	mdPath := filepath.Join(dirName, "index.md")
	file, err := os.Create(mdPath)
	if err != nil {
		return fmt.Errorf("failed to create markdown file: %v", err)
	}
	defer file.Close()

	tmpl, err := template.New("markdown").Parse(mdTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	data := TmplData{Date: now}
	if err = tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func isValidArg(arg string) bool {
	matched, _ := regexp.MatchString(`^[0-9]{1,3}$`, arg)
	return matched
}
