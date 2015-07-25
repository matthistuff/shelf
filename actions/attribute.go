package actions

import (
	"github.com/codegangsta/cli"
	"strings"
	"github.com/matthistuff/shelf/helpers"
	"fmt"
	"github.com/matthistuff/shelf/data"
)

func AddAttribute(c *cli.Context) {
	if len(c.Args()) < 3 {
		helpers.ErrExit(true, "Not enough arguments provided")
	}

	objectId := c.Args().First()
	name := c.Args().Get(1)
	value := strings.Join(c.Args()[2:], " ")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if (object.HasAttribute(name, value)) {
		return
	}

	object.AddAttribute(name, value)
}

func RemoveAttribute(c *cli.Context) {
	if len(c.Args()) < 3 {
		helpers.ErrExit(true, "Not enough arguments provided")
	}

	objectId := c.Args().First()
	name := c.Args().Get(1)
	value := strings.Join(c.Args()[2:], " ")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	object.RemoveAttribute(name, value)
}