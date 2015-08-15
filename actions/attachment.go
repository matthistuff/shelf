package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"os"
	"io"
	"path"
	"fmt"
	"github.com/matthistuff/shelf/helpers"
	"gopkg.in/mgo.v2/bson"
)


func AddAttachment(c *cli.Context) {
	objectId := c.Args().First()
	filepath := c.Args().Get(1)

	helpers.ErrExit(objectId == "", "No object ID given!")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	dbFile, err := data.Files().Create("")
	helpers.ErrPanic(err)

	file, err := os.Open(filepath)
	helpers.ErrPanic(err)
	defer file.Close()

	_, err = io.Copy(dbFile, file)
	helpers.ErrPanic(err)

	err = dbFile.Close()
	helpers.ErrPanic(err)

	attachment := data.CreateAttachment(dbFile, path.Base(filepath))
	object.Attachments = append(object.Attachments, *attachment)
	object.Update()

	fmt.Println(attachment.Id.Hex())
}

func GetAttachment(c *cli.Context) {
	attachmentId := c.Args().First()

	helpers.ErrExit(attachmentId == "", "No object ID given!")

	file, err := data.Files().OpenId(bson.ObjectIdHex(attachmentId))
	helpers.ErrPanic(err)

	_, err = io.Copy(os.Stdout, file)
	helpers.ErrPanic(err)

	err = file.Close()
	helpers.ErrPanic(err)
}

func ListAttachments(c *cli.Context) {
	helpers.Color(c)

	objectId := c.Args().First()
	helpers.ErrExit(objectId == "", "No object ID given!")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	for index, attachment := range object.Attachments {
		fmt.Printf("(%s) %s \"%s\"\n", helpers.ShortId(index+1), helpers.ObjectId(attachment.Id.Hex()), attachment.Filename)
	}
}