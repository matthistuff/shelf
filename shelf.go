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
	app.Usage = "a simple document management system"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "no-color",
			Usage: "disable colored output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "create",
			Usage: "create an object",
			Action: actions.CreateObject,
		},
		{
			Name: "info",
			Usage: "print information about an object",
			Action: actions.GetObject,
		},
		{
			Name: "attach",
			Usage: "attach a file to an object",
			Action: actions.AddAttachment,
		},
		{
			Name: "attachments",
			Usage: "list all attachments of an object",
			Action: actions.ListAttachments,
		},
		{
			Name: "retrieve",
			Usage: "send an attachment to stdout",
			Action: actions.GetAttachment,
		},
		{
			Name: "attribute",
			Usage: "manage object attributes",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Usage: "add an attribute to an object",
					Action: actions.AddAttribute,
				},
				{
					Name: "remove",
					Usage: "remove an attribute from an object",
					Action: actions.RemoveAttribute,
				},
			},
		},
		{
			Name: "tag",
			Usage: "add a tag to an object",
			Action: actions.AddTag,
		},
		{
			Name: "untag",
			Usage: "remove a tag from an object",
			Action: actions.RemoveTag,
		},
		{
			Name: "search",
			Usage: "search objects",
			Action: actions.Search,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "page",
					Value: 1,
					Usage:"search results page",
				},
			},
		},
	}

	_, session := data.DB()
	defer session.Close()

	// db.objects.find({$text:{$search:"values"}}, {score: { $meta: "textScore" }}).sort({ score: { $meta: "textScore" } }).skip(0).limit(100)
	data.Objects().EnsureIndex(mgo.Index{
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
