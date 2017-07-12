// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Sample": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-code-reading/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-code-reading
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// example (default view)
//
// Identifier: application/vnd.sample+json; view=default
type Sample struct {
	// デフォルト値
	DefaultType string `form:"defaultType" json:"defaultType" xml:"defaultType"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	EnumType string `form:"enumType" json:"enumType" xml:"enumType"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 数字（1〜10）
	IntegerType int `form:"integerType" json:"integerType" xml:"integerType"`
	// デフォルト値
	Reg string `form:"reg" json:"reg" xml:"reg"`
	// 文字（1~10文字）
	StringType string `form:"stringType" json:"stringType" xml:"stringType"`
}

// Validate validates the Sample media type instance.
func (mt *Sample) Validate() (err error) {

	if mt.StringType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stringType"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.EnumType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "enumType"))
	}
	if mt.DefaultType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "defaultType"))
	}
	if mt.Reg == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "reg"))
	}
	return
}

// SampleCollection is the media type for an array of Sample (default view)
//
// Identifier: application/vnd.sample+json; type=collection; view=default
type SampleCollection []*Sample

// Validate validates the SampleCollection media type instance.
func (mt SampleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
