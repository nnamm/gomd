package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/nnamm/gomd/clock"
)

const frontMatter = `---
title: 
description: 
summary: 
categories: [""]
tags: [""]
clock: {{.clock}}
lastmod: {{.LastMod}}
slug: 
---


` // Contains 2 blank lines

type source struct {
	//WorkDir  string
	DirName     string
	CurrentDate string
	//FileName string
}

func (s *source) setDirName(dirNo string) error {
	l := utf8.RuneCountInString(dirNo)
	if l > 3 {
		return fmt.Errorf("DirNo must be under 3-digits: %d", l)
	}

	getint, err := strconv.Atoi(dirNo)
	if err != nil {
		return fmt.Errorf("DirNo must be all numeric: %v", err)
	}

	s.DirName = fmt.Sprintf("%03d", getint)
	return nil
}

type date struct {
	clocker clock.Clocker
}

func (s *source) getCurrentDate(d date) {
	n := d.clocker.Now()
	s.CurrentDate = n.Format("060102")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Required args are missing")
		return
	}

	s := source{}
	if err := s.setDirName(os.Args[1]); err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Println(s.DirName)

	//wd, err := os.Getwd()
	//if err != nil {
	//  fmt.Printf("Fail to getwd: %v\n", err)
	//  return
	//}
	//
	//ct := &Content{
	//  WorkDir:  wd,
	//  DirNo:    os.Args[1],
	//  FileName: "index.md",
	//}
	//
	//if err := ct.makeDir(); err != nil {
	//  fmt.Printf("Fail to make dir: %v\n", err)
	//  return
	//}
	//
	//if err = ct.createMarkdown(); err != nil {
	//  fmt.Printf("Fail to create markdown file: %v\n", err)
	//  return
	//}
}

//func (ct *Content) makeDir() error {
//  var dir = filepath.Join(ct.WorkDir, ct.DirNo)
//  err := os.Mkdir(dir, 0755)
//  if os.IsExist(err) {
//      return errors.New("the directory already exists")
//  } else if err != nil {
//      return err
//  }
//  return nil
//}
//
//func (ct *Content) createMarkdown() error {
//  fp := filepath.Join(ct.WorkDir, ct.DirNo, ct.FileName)
//  md, err := os.Create(fp)
//  if err != nil {
//      return err
//  }
//  defer func(md *os.File) {
//      err := md.Close()
//      if err != nil {
//          fmt.Println(err)
//      }
//  }(md)
//
//  tmpl, err := template.New("").Parse(frontMatter)
//  if err != nil {
//      return err
//  }
//  n := time.Now().Format("2006-01-02 15:04:05")
//  fm := struct {
//      clock    string
//      LastMod string
//  }{
//      clock:    n,
//      LastMod: n,
//  }
//  if err := tmpl.Execute(md, fm); err != nil {
//      return err
//  }
//  return nil
//}
