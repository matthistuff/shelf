package colors

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var (
	Header   = color.New(color.FgRed, color.Bold).SprintFunc()
	ObjectId = color.New(color.FgGreen).SprintFunc()
	ShortId  = color.New(color.FgMagenta).SprintFunc()
	Bold     = color.New(color.Bold).SprintFunc()
)

func Allow(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}
