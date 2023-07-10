package models

type Blog struct {
	Base
	Title   string `json:"title"`
	Content string `json:"content"`
	Authro  string `json:"author"`
}

func (m Blog) TableName() string {
	return "blog"
}
