package data
import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type Attachment struct {
	UploadDate time.Time `bson:"uploadDate" json:"uploadDate"`
	Filename   string `bson:"filename" json:"filename"`
	Content    string `bson:"content" json:"content"`
	Id         bson.ObjectId `bson:"_id" json:"id"`
}

func CreateAttachment(file *mgo.GridFile, filename string) *Attachment {
	return &Attachment{
		UploadDate: file.UploadDate(),
		Filename: filename,
		Content: "",
		Id: file.Id().(bson.ObjectId),
	}
}