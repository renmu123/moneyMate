package models

type Category struct {
	ID   int
	Name string `validate:"required" json:"name"`
	Icon string
}
