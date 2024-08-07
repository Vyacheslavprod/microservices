package models

type Note struct {
	Id       string  `json:"id,omitempty" bson:"id,omitempty" `
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`
	Content  *string `json:"content,omitempty" bson:"content,omitempty"`
	AuthorID uint    `json:"author_id,omitempty" bson:"author_id,omitempty"`
}