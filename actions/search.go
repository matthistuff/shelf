package actions
import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

func Search(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")

	search := strings.Join(c.Args(), " ")
	page := c.Int("page")
	perPage := 10

	db, _ := data.DB()

	result := []data.Object{}
	query := db.C("objects").Find(bson.M{
		"$text":bson.M{
			"$search":search,
		},
	}).Select(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	}).Sort("$textScore:score")
	total, _ := query.Count()
	query.Skip((page-1)*perPage).Limit(perPage).All(&result)

	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	for index, object := range result {
		fmt.Printf("(%s) %s \"%s\"\n", bold(index+1), green(object.Id.Hex()), object.Title)
	}
	fmt.Printf("Page %s of %s\n", bold(strconv.Itoa(page)), bold(strconv.Itoa(int(total/perPage)+1)))
}