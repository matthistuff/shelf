package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/colors"
	"github.com/matthistuff/shelf/data"
	"github.com/matthistuff/shelf/helper"
	"gopkg.in/mgo.v2/bson"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func AddAttachment(c *cli.Context) {
	objectId := helper.ValidId(c.Args().First())
	attachmentPath := c.Args().Get(1)

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	dbFile, err := data.Files().Create("")
	helper.ErrPanic(err)

	file, err := os.Open(attachmentPath)
	helper.ErrPanic(err)
	defer file.Close()

	_, err = io.Copy(dbFile, file)
	helper.ErrPanic(err)

	err = dbFile.Close()
	helper.ErrPanic(err)

	var (
		content  = ""
		metadata = make(map[string]string)
	)

	if strings.ToLower(filepath.Ext(file.Name())) == ".pdf" {
		if c.BoolT("extract-pdf-text") {
			content, metadata, err = helper.ConvertPDF(file)
			helper.ErrPanic(err)
		}
	}

	attachment := data.CreateAttachment(dbFile, path.Base(attachmentPath), content, metadata)
	object.Attachments = append(object.Attachments, *attachment)
	object.Update()

	fmt.Println(attachment.Id.Hex())
}

func GetAttachment(c *cli.Context) {
	objectId := helper.ValidId(c.Args().First())

	file, err := data.Files().OpenId(bson.ObjectIdHex(objectId))
	helper.ErrPanic(err)

	_, err = io.Copy(os.Stdout, file)
	helper.ErrPanic(err)

	err = file.Close()
	helper.ErrPanic(err)
}

func DeleteAttachment(c *cli.Context) {
	objectId := helper.ValidId(c.Args().First())

	err := data.Files().RemoveId(bson.ObjectIdHex(objectId))
	helper.ErrPanic(err)

	query := data.Objects().Find(bson.M{
		"attachments._id": bson.ObjectIdHex(objectId),
	})

	result := []data.Object{}
	query.All(&result)

	for _, object := range result {
		object.RemoveAttachment(objectId)
	}
}

func ListAttachments(c *cli.Context) {
	colors.Allow(c)

	objectId := helper.ValidId(c.Args().First())

	object, err := data.GetObject(objectId)
	helper.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	if len(object.Attachments) > 0 {
		data.ClearCache()
		defer data.FlushCache()

		for index, attachment := range object.Attachments {
			fmt.Printf("(%s) %s \"%s\"\n", colors.ShortId(index+1), colors.ObjectId(attachment.Id.Hex()), attachment.Filename)
			data.SetCache(strconv.Itoa(index+1), attachment.Id.Hex())
		}
	}
}
