package helpers
import (
	"github.com/fatih/color"
	"github.com/codegangsta/cli"
)

var (
	Header = color.New(color.FgRed, color.Bold).SprintFunc()
	ObjectId = color.New(color.FgGreen).SprintFunc()
	ShortId = color.New(color.FgMagenta, color.Bold).SprintFunc()
	Bold = color.New(color.Bold).SprintFunc()
)

func Color(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}