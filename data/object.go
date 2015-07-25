package data

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Object struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string `bson:"title" json:"title"`
	CreateDate  time.Time `bson:"createDate" json:"createDate"`
	Attributes  []Attribute `bson:"attributes" json:"attributes"`
	Attachments []Attachment `bson:"attachments" json:"attachments"`
}

func (o *Object) Update() error {
	db, _ := DB()

	return db.C("objects").UpdateId(o.Id, o)
}

func (o *Object) AddAttribute(name string, value string) {
	attribute := Attribute{
		Name: name,
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
		Id: bson.NewObjectId(),
		Title: title,
		CreateDate: time.Now(),
	}
}

func GetObject(id string) (*Object, error) {
	db, _ := DB()

	object := Object{}
	err := db.C("objects").FindId(bson.ObjectIdHex(id)).One(&object)

	return &object, err
}