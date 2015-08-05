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
	"strconv"
)

func CreateObject(c *cli.Context) {
	title := strings.Join(c.Args(), " ")

	object := data.CreateObject(title)

	err := data.Objects().Insert(object)
	helpers.ErrPanic(err)

	fmt.Printf("%s\n", object.Id.Hex())
}

func DeleteObject(c *cli.Context) {
	objectId := c.Args().First()

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	files := data.Files()
	for _, attachment := range object.Attachments {
		files.RemoveId(attachment.Id)
	}

	data.Objects().RemoveId(object.Id)
}

func GetObjects(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")

	page := c.Int("page")
	perPage := 10

	query := data.Objects().Find(nil).Sort("-_id")

	total, _ := query.Count()
	result := []data.Object{}
	query.Skip((page-1)*perPage).Limit(perPage).All(&result)

	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	if total > 0 {
		for index, object := range result {
			fmt.Printf("(%s) %s \"%s\"\n", bold(index+1), green(object.Id.Hex()), object.Title)
		}
		fmt.Printf("Page %s of %s\n", bold(strconv.Itoa(page)), bold(strconv.Itoa(int(total/perPage)+1)))
	}
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

	fmt.Printf("%s\n\n", bold(object.Title))
	fmt.Printf("%s\n\t%s\n", red("Created"), object.CreateDate.Format(time.RFC1123))

	if _, exists := attributes["content"]; exists {
		fmt.Printf("\n%s\n\t%s\n", red("Content"), strings.Join(attributes["content"], ", "))
		delete(attributes, "content")
	}

	if len(attributes) > 0 {
		fmt.Printf("\n%s\n", red("Attributes"))

		for k := range keys {
			sort.Strings(attributes[keys[k]])

			fmt.Printf("\t%s: %s\n", bold(keys[k]), strings.Join(attributes[keys[k]], ", "))
		}
	}

	if len(object.Attachments) > 0 {
		fmt.Printf("\n%s\n", red("Attachments"))

		for _, attachment := range object.Attachments {
			fmt.Printf("\t%s: %s (%s)\n", green(attachment.Id.Hex()), attachment.Filename, attachment.UploadDate.Format(time.RFC1123))
		}
	}
}