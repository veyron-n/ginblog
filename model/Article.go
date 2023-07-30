package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(20); not null" json:"title"`
	Cid     int    `gorm:"type:int; not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200); not null" json:"desc"`
	Context string `gorm:"type:longtext; not null" json:"context"`
	Img     string `gorm:"type:varchar(100);" json:"img"`
}

func (Article) TableName() string {
	return "article" // 返回你想要使用的表名
}

// CheckArticle 查询文章是否存在
func CheckArticle(title string) (code int) {
	var article Article
	db.Select("id").Where("title = ?", title).First(&article)
	fmt.Print("article.ID", article)
	if article.ID > 0 {
		return errmsg.ERROR_ARTICLE_USED
	}
	return errmsg.SUCCESS
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DelArticle 删除文章
func DelArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["context"] = data.Context
	maps["img"] = data.Img
	err := db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArticle 查询分类下的所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIT, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_CATEGORY_NOT_EXIT
	}
	return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.ERROR, total
}
