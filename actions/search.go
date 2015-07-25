package actions
import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

func Search(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")

	page := c.Int("page")
	perPage := 10

	searchQuery := data.ParseQuery(c.Args())
	search := []bson.M{}
	if searchQuery.Text != "" {
		search = append(search, bson.M{
			"$text":bson.M{
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

	db, _ := data.DB()
	query := db.C("objects").Find(bson.M{
		"$and": search,
	}).Select(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	}).Sort("$textScore:score", "-_id")

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