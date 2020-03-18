package models

type Category struct {
	Model

	Displayname string `json:"displayname"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
}
