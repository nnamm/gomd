package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const frontMatter = `---
title: 
description: 
summary: 
categories: [""]
tags: [""]
date: {{.Date}}
lastmod: {{.LastMod}}
slug: 
---


` // Contains 2 blank lines

type Content struct {
	WorkDir  string
	DirNo    string
	FileName string
}

func (ct *Content) makeDir() error {
	var dir = filepath.Join(ct.WorkDir, ct.DirNo)
	err := os.Mkdir(dir, 0755)
	if os.IsExist(err) {
		return errors.New("the directory already exists")
	} else if err != nil {
		return err
	}
	return nil
}

func (ct *Content) createMarkdown() error {
	fp := filepath.Join(ct.WorkDir, ct.DirNo, ct.FileName)
	md, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer func(md *os.File) {
		err := md.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(md)

	tmpl, err := template.New("").Parse(frontMatter)
	if err != nil {
		return err
	}
	n := time.Now().Format("2006-01-02 15:04:05")
	fm := struct {
		Date    string
		LastMod string
	}{
		Date:    n,
		LastMod: n,
	}
	if err := tmpl.Execute(md, fm); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Required args are missing")
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Fail to getwd: %v\n", err)
		return
	}

	ct := &Content{
		WorkDir:  wd,
		DirNo:    os.Args[1],
		FileName: "index.md",
	}

	if err := ct.makeDir(); err != nil {
		fmt.Printf("Fail to make dir: %v\n", err)
		return
	}

	if err = ct.createMarkdown(); err != nil {
		fmt.Printf("Fail to create markdown file: %v\n", err)
		return
	}
}
