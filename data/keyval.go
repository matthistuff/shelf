package data

type KeyVal struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}
