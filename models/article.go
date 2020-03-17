package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

// gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
// Tag字段，实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能
// time.Now().Unix() 返回当前的时间戳

type Article struct {
	Model

	//索引
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	//简述
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
