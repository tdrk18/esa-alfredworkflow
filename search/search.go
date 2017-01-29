// -*- coding: utf-8 -*-
//---------------------------------#
// File Name     : search/search.go
// Author        : todoroki
// Date Created  : 2017-01-29
//---------------------------------#

package search

import (
	"os"
	"strings"

	"../env"
	"../esa"

	"github.com/codegangsta/cli"
	"github.com/pascalw/go-alfred"
	"github.com/yuichiro-h/go-esa"
)

func SearchCommand(c *cli.Context) {
	var query = strings.Join(c.Args(), " ")

	client := esaClient.NewEsaClient()
	teamName := os.Getenv(env.EsaTeamName)

	response := alfred.NewResponse()
	res, err := client.GetTeamPosts(teamName, &esa.GetTeamPostsRequest{
		Q: esa.String(query),
	})

	if err != nil {
		response.AddItem(&alfred.AlfredResponseItem{
			Valid: true,
			Uid:   "error",
			Title: err.Error(),
			Icon:  "icon/esa-piyo.png",
		})
		response.Print()
		return
	}

	posts := res.Posts

	for _, post := range posts {
		title := post.Name

		var subTitle string
		if post.Category != nil {
			subTitle = *post.Category
		} else {
			subTitle = ""
		}

		response.AddItem(&alfred.AlfredResponseItem{
			Valid:    true,
			Uid:      post.URL,
			Title:    title,
			Subtitle: subTitle + " @" + post.CreatedBy.ScreenName + " " + post.CreatedAt.Format("2016/01/02 15:04:05"),
			Arg:      post.URL,
			Icon:     "icon/esa-tori.png",
		})
	}
	response.Print()
}
