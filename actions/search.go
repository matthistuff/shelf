package actions
import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"fmt"
)

func Search(c *cli.Context) {
	search := strings.Join(c.Args(), " ")
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
	query.Limit(10).All(&result)

	for _, object := range result {
		fmt.Println(object.Id.Hex())
	}
}