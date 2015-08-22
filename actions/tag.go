package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"github.com/matthistuff/shelf/helper"
	"strings"
)

func AddTag(c *cli.Context) {
	if len(c.Args()) < 2 {
		helper.ErrExit(true, "Not enough arguments provided")
	}

	objectId := helper.ValidId(c.Args().First())
	value := strings.Join(c.Args().Tail(), " ")

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if object.HasTag(value) {
		return
	}

	object.AddTag(value)
}

func RemoveTag(c *cli.Context) {
	if len(c.Args()) < 2 {
		helper.ErrExit(true, "Not enough arguments provided")
	}

	objectId := helper.ValidId(c.Args().First())
	value := strings.Join(c.Args().Tail(), " ")

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	object.RemoveTag(value)
}
