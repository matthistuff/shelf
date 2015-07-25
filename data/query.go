package data
import "regexp"

type Query struct {
	Text           string
	AttributeQuery []AttributeQuery
}

type AttributeQuery struct {
	Name   string
	Value  string
	Exact bool
}

var AttributeQueryReg = regexp.MustCompile("([^:]+):?(.*)")

func ParseQuery(query []string) Query {
	parsed := Query{}

	for _, part := range query {
		matches := AttributeQueryReg.FindStringSubmatch(part)

		if matches[2] == "" {
			parsed.Text += matches[0] + " "
		} else {
			parsed.AttributeQuery = append(parsed.AttributeQuery, AttributeQuery{
				Name: matches[1],
				Value: matches[2],
				Exact: false,
			})
		}
	}

	return parsed
}