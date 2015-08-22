package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"github.com/matthistuff/shelf/helper"
	"strings"
)

func AddAttribute(c *cli.Context) {
	if len(c.Args()) < 3 {
		helper.ErrExit(true, "Not enough arguments provided")
	}

	objectId := helper.ValidId(c.Args().First())

	name := c.Args().Get(1)
	value := strings.Join(c.Args()[2:], " ")

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if object.HasAttribute(name, value) {
		return
	}

	object.AddAttribute(name, value)
}

func RemoveAttribute(c *cli.Context) {
	if len(c.Args()) < 3 {
		helper.ErrExit(true, "Not enough arguments provided")
	}

	objectId := helper.ValidId(c.Args().First())

	name := c.Args().Get(1)
	value := strings.Join(c.Args()[2:], " ")

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	object.RemoveAttribute(name, value)
}
