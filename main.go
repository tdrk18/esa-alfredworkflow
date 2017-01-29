// -*- coding: utf-8 -*-
//---------------------------------#
// File Name     : main.go
// Author        : todoroki
// Date Created  : 2017-01-29
//---------------------------------#

package main

import (
	"os"
	"os/exec"

	"./search"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "esa-alfredworkflow"
	version, _ := exec.Command("cat", "VERSION").Output()
	app.Version = string(version)
	app.Commands = []cli.Command{
		{
			Name:   "search",
			Action: search.SearchCommand,
		},
	}
	app.Run(os.Args)
}
