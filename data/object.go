package data

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Object struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title"`
	CreateDate  time.Time     `bson:"createDate" json:"createDate"`
	Attributes  []KeyVal      `bson:"attributes" json:"attributes"`
	Attachments []Attachment  `bson:"attachments" json:"attachments"`
}

type SearchObject struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Score       float64       `bson:"score" json:"score"`
	Object
}

func (o *Object) Update() error {
	return Objects().UpdateId(o.Id, o)
}

func (o *Object) AddAttribute(name string, value string) {
	attribute := KeyVal{
		Name:  name,
		Value: value,
	}

	o.Attributes = append(o.Attributes, attribute)
	o.Update()
}

func (o *Object) AttributeIndex(name string, value string) int {
	for index, attr := range o.Attributes {
		if attr.Name == name && attr.Value == value {
			return index
		}
	}

	return -1
}

func (o Object) HasAttribute(name string, value string) bool {
	return o.AttributeIndex(name, value) > -1
}

func (o *Object) RemoveAttribute(name string, value string) {
	if index := o.AttributeIndex(name, value); index > -1 {
		o.Attributes = append(o.Attributes[:index], o.Attributes[index+1:]...)
		o.Update()
	}
}

func (o *Object) GetAttachment(objectId string) Attachment {
	return o.Attachments[o.AttachmentIndex(objectId)]
}

func (o *Object) AttachmentIndex(objectId string) int {
	for index, attachment := range o.Attachments {
		if attachment.Id.Hex() == objectId {
			return index
		}
	}

	return -1
}

func (o *Object) RemoveAttachment(objectId string) {
	if index := o.AttachmentIndex(objectId); index > -1 {
		o.Attachments = append(o.Attachments[:index], o.Attachments[index+1:]...)
		o.Update()
	}
}

func (o *Object) AddTag(value string) {
	o.AddAttribute("tag", value)
}

func (o Object) HasTag(tag string) bool {
	return o.HasAttribute("tag", tag)
}

func (o Object) RemoveTag(tag string) {
	o.RemoveAttribute("tag", tag)
}

func CreateObject(title string) *Object {
	return &Object{
		Id:         bson.NewObjectId(),
		Title:      title,
		CreateDate: time.Now(),
	}
}

func GetObject(id string) (*Object, error) {
	object := Object{}
	err := Objects().FindId(bson.ObjectIdHex(id)).One(&object)

	return &object, err
}
