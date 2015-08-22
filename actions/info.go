package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"github.com/matthistuff/shelf/helper"
)

func GetInfo(c *cli.Context) {
	objectId := helper.ValidId(c.Args().First())

	_, err := data.GetObject(objectId)
	if err == nil {
		GetObject(c)
		return
	}

	_, err = data.GetAttachment(objectId)
	if err == nil {
		GetAttachmentInfo(c)
		return
	}

	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))
}
