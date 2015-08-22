package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Attachment struct {
	Id         bson.ObjectId     `bson:"_id" json:"id"`
	UploadDate time.Time         `bson:"uploadDate" json:"uploadDate"`
	Filename   string            `bson:"filename" json:"filename"`
	MetaData   map[string]string `bson:"metadata" json:"metadata"`
	Content    string            `bson:"content" json:"content"`
}

func CreateAttachment(file *mgo.GridFile, filename string, content string, metadata map[string]string) *Attachment {
	return &Attachment{
		Id:         file.Id().(bson.ObjectId),
		UploadDate: file.UploadDate(),
		Filename:   filename,
		MetaData:   metadata,
		Content:    content,
	}
}
