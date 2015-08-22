package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/colors"
	"github.com/matthistuff/shelf/data"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func Search(c *cli.Context) {
	colors.Allow(c)

	page := c.Int("page")
	perPage := 10

	searchQuery := data.ParseQuery(c.Args())
	search := []bson.M{}
	if searchQuery.Text != "" {
		search = append(search, bson.M{
			"$text": bson.M{
				"$search": searchQuery.Text,
			},
		})
	}

	for _, attrQuery := range searchQuery.AttributeQuery {
		search = append(search, bson.M{
			"attributes": bson.M{
				"$elemMatch": bson.M{
					"name": attrQuery.Name,
					"value": bson.M{
						"$regex": attrQuery.Value,
					},
				},
			},
		})
	}

	query := data.Objects().Find(bson.M{
		"$and": search,
	}).Select(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	}).Sort("$textScore:score", "-_id")

	total, _ := query.Count()
	result := []data.Object{}
	query.Skip((page - 1) * perPage).Limit(perPage).All(&result)

	if total > 0 {
		for index, object := range result {
			fmt.Printf("(%s) %s \"%s\"\n", colors.ShortId(index+1), colors.ObjectId(object.Id.Hex()), object.Title)
		}
		fmt.Printf("Page %s of %s\n", colors.Bold(strconv.Itoa(page)), colors.Bold(strconv.Itoa(int(total/perPage)+1)))
	}
}
