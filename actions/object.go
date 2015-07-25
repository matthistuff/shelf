package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/shelf/data"
	"fmt"
	"github.com/matthistuff/shelf/helpers"
	"strings"
)

func CreateObject(c *cli.Context) {
	db, _ := data.DB()
	title := strings.Join(c.Args(), " ")

	object := data.CreateObject(title)

	err := db.C("objects").Insert(object)
	helpers.ErrPanic(err)

	fmt.Printf("%s\n", object.Id.Hex())
}