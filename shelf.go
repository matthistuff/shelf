package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/actions"
	"github.com/matthistuff/shelf/data"
	"gopkg.in/mgo.v2"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "shelf"
	app.Usage = "A simple document management system"

	app.Commands = []cli.Command{
		{
			Name: "create",
			Usage: "Create an object",
			Action: actions.CreateObject,
		},
		{
			Name: "attach",
			Usage: "Attach a file to an object",
			Action: actions.AddAttachment,
		},
		{
			Name: "attachments",
			Usage: "List all attachments of an object",
			Action: actions.ListAttachments,
		},
		{
			Name: "attribute",
			Usage: "Manage object attributes",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Usage: "Add an attribute to an object",
					Action: actions.AddAttribute,
				},
				{
					Name: "remove",
					Usage: "Remove an attribute from an object",
					Action: actions.RemoveAttribute,
				},
			},
		},
		{
			Name: "tag",
			Usage: "Add a tag to an object",
			Action: actions.AddTag,
		},
		{
			Name: "untag",
			Usage: "Remove a tag from an object",
			Action: actions.RemoveTag,
		},
		{
			Name: "search",
			Usage: "Search objects",
			Action: actions.Search,
		},
	}

	db, session := data.DB()
	defer session.Close()

	// db.objects.find({$text:{$search:"values"}}, {score: { $meta: "textScore" }}).sort({ score: { $meta: "textScore" } }).skip(0).limit(100)
	db.C("objects").EnsureIndex(mgo.Index{
		Key: []string{
			"$text:title",
			"$text:attachments.content",
			"$text:attributes.value",
		},
		Weights: map[string]int{
			"title": 10,
			"attachments.content": 5,
			"attributes.value": 1,
		},
	})

	app.Run(os.Args)
}
