package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"fmt"
	"github.com/matthistuff/shelf/helpers"
	"strings"
	"sort"
	"github.com/fatih/color"
	"time"
)

func CreateObject(c *cli.Context) {
	db, _ := data.DB()
	title := strings.Join(c.Args(), " ")

	object := data.CreateObject(title)

	err := db.C("objects").Insert(object)
	helpers.ErrPanic(err)

	fmt.Printf("%s\n", object.Id.Hex())
}

func GetObject(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")

	objectId := c.Args().First()
	helpers.ErrExit(objectId == "", "No object ID given!")

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	attributes := make(map[string][]string)
	for _, attribute := range object.Attributes {
		attributes[attribute.Name] = append(attributes[attribute.Name], attribute.Value)
	}

	keys := make([]string, 0, len(attributes))
	for k := range attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	red := color.New(color.FgRed, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Printf("%s\n", red(object.Title))
	fmt.Printf("Created at %s\n", object.CreateDate.Format(time.RFC1123))
	fmt.Printf("\n%s\n", red("Attributes"))

	for k := range keys {
		sort.Strings(attributes[keys[k]])

		fmt.Printf("\t%s: %s\n", bold(keys[k]), strings.Join(attributes[keys[k]], ", "))
	}

	fmt.Printf("\n%s\n", red("Attachments"))
	for _, attachment := range object.Attachments {
		fmt.Printf("\t%s: %s (%s)\n", green(attachment.Id.Hex()), attachment.Filename, attachment.UploadDate.Format(time.RFC1123))
	}
}