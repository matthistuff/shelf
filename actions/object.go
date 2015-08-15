package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"fmt"
	"github.com/matthistuff/shelf/helpers"
	"strings"
	"sort"
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
	helpers.Color(c)

	query := data.Objects().Find(nil).Sort("-_id")

	page := c.Int("page")
	perPage := 10
	total, _ := query.Count()
	result := []data.Object{}
	query.Skip((page-1)*perPage).Limit(perPage).All(&result)

	data.ClearCache()
	defer data.FlushCache()

	if total > 0 {
		for index, object := range result {
			fmt.Printf("(%s) %s \"%s\"\n", helpers.ShortId(index+1), helpers.ObjectId(object.Id.Hex()), object.Title)
			data.SetCache(strconv.Itoa(index+1), object.Id.Hex())
		}
		fmt.Printf("Page %s of %s\n", helpers.Bold(strconv.Itoa(page)), helpers.Bold(strconv.Itoa(int(total/perPage)+1)))
	}
}

func GetObject(c *cli.Context) {
	helpers.Color(c)

	objectId, exists := data.AssertGuid(c.Args().First())
	helpers.ErrExit(objectId == "", "No object ID given!")
	helpers.ErrExit(!exists, fmt.Sprintf("No cached entry %s exists!", c.Args().First()))

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

	fmt.Printf("%s\n\n", helpers.Bold(object.Title))
	fmt.Printf("%s\n\t%s\n", helpers.Header("Created"), object.CreateDate.Format(time.RFC1123))

	if _, exists := attributes["content"]; exists {
		fmt.Printf("\n%s\n\t%s\n", helpers.Header("Content"), strings.Join(attributes["content"], ", "))
		delete(attributes, "content")
	}

	if len(attributes) > 0 {
		fmt.Printf("\n%s\n", helpers.Header("Attributes"))

		for k := range keys {
			sort.Strings(attributes[keys[k]])

			fmt.Printf("\t%s: %s\n", helpers.Bold(keys[k]), strings.Join(attributes[keys[k]], ", "))
		}
	}

	if len(object.Attachments) > 0 {
		data.ClearCache()
		defer data.FlushCache()

		fmt.Printf("\n%s\n", helpers.Header("Attachments"))

		for index, attachment := range object.Attachments {
			fmt.Printf("\t(%s) %s: %s (%s)\n", helpers.ShortId(index+1), helpers.ObjectId(attachment.Id.Hex()), attachment.Filename, attachment.UploadDate.Format(time.RFC1123))
			data.SetCache(strconv.Itoa(index+1), object.Id.Hex())
		}
	}
}