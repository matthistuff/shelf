package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/colors"
	"github.com/matthistuff/shelf/data"
	"github.com/matthistuff/shelf/helpers"
	"sort"
	"strconv"
	"strings"
	"time"
)

func CreateObject(c *cli.Context) {
	title := strings.Join(c.Args(), " ")

	object := data.CreateObject(title)

	err := data.Objects().Insert(object)
	helpers.ErrPanic(err)

	fmt.Printf("%s\n", object.Id.Hex())
}

func DeleteObject(c *cli.Context) {
	objectId := helpers.ValidId(c.Args().First())

	object, err := data.GetObject(objectId)
	helpers.ErrExit(err != nil, fmt.Sprintf("Invalid object ID %s!\n", objectId))

	files := data.Files()
	for _, attachment := range object.Attachments {
		files.RemoveId(attachment.Id)
	}

	data.Objects().RemoveId(object.Id)
}

func GetObjects(c *cli.Context) {
	colors.Allow(c)

	query := data.Objects().Find(nil).Sort("-_id")

	page := c.Int("page")
	perPage := 10
	total, _ := query.Count()
	result := []data.Object{}
	query.Skip((page - 1) * perPage).Limit(perPage).All(&result)

	if total > 0 {
		data.ClearCache()
		defer data.FlushCache()

		for index, object := range result {
			fmt.Printf("(%s) %s \"%s\"\n", colors.ShortId(index+1), colors.ObjectId(object.Id.Hex()), object.Title)
			data.SetCache(strconv.Itoa(index+1), object.Id.Hex())
		}
		fmt.Printf("Page %s of %s\n", colors.Bold(strconv.Itoa(page)), colors.Bold(strconv.Itoa(int(total/perPage)+1)))
	}
}

func GetObject(c *cli.Context) {
	colors.Allow(c)

	objectId := helpers.ValidId(c.Args().First())

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

	fmt.Printf("%s\n\n", colors.Bold(object.Title))
	fmt.Printf("%s\n\t%s\n", colors.Header("Created"), object.CreateDate.Format(time.RFC1123))

	if _, exists := attributes["content"]; exists {
		fmt.Printf("\n%s\n\t%s\n", colors.Header("Content"), strings.Join(attributes["content"], ", "))
		delete(attributes, "content")
	}

	if len(attributes) > 0 {
		fmt.Printf("\n%s\n", colors.Header("Attributes"))

		for k := range keys {
			sort.Strings(attributes[keys[k]])

			fmt.Printf("\t%s: %s\n", colors.Bold(keys[k]), strings.Join(attributes[keys[k]], ", "))
		}
	}

	if len(object.Attachments) > 0 {
		data.ClearCache()
		defer data.FlushCache()

		fmt.Printf("\n%s\n", colors.Header("Attachments"))

		for index, attachment := range object.Attachments {
			fmt.Printf("\t(%s) %s: %s (%s)\n", colors.ShortId(index+1), colors.ObjectId(attachment.Id.Hex()), attachment.Filename, attachment.UploadDate.Format(time.RFC1123))
			data.SetCache(strconv.Itoa(index+1), attachment.Id.Hex())
		}
	}
}
