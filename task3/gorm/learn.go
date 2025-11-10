package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Postcount int
	Posts     []Post
}
type Post struct {
	gorm.Model
	UserID        uint
	CommentsCount string
	Comments      []Comment
}
type Comment struct {
	gorm.Model
	PostID uint
}

/*
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
func main() {
	dsn := "root:5594210@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	ctx := context.Background()

	//gorm.G[Post](db).Create(ctx, &Post{UserID: 1})
	cm, err := gorm.G[Comment](db).Where("post_id = ?", 1).First(ctx)
	if err != nil {
		fmt.Printf("select err:\n%+v\n", err)
	} else {
		db.Delete(&cm)
	}
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func SelectAllByUser(db *gorm.DB, userid uint, ctx *context.Context) {
	///
	out, err := gorm.G[User](db).Preload("Posts.Comments", nil).Where("id = ?", userid).First(*ctx)
	if err != nil {
		fmt.Printf("select err:\n%+v\n", err)
	}
	fmt.Println(out)
	//Preload("Classes.Students").
	//         Where("name = ?", gradeName).
	//         First(&grade)
}

func SelectMostComment(db *gorm.DB, ctx *context.Context) {

	//var result map[string]interface{}
	var result []map[string]interface{}
	err := gorm.G[Comment](db).Select("post_id, count(post_id) as count").Group("post_id").Order("count desc").Scan(*ctx, &result)
	if err != nil {
		fmt.Printf("select err:\n%+v\n", err)
	}

	fmt.Printf("%+v", result)
	if len(result) > 0 {
		post, err := gorm.G[Post](db).Where("id = ?", result[0]["post_id"]).First(*ctx)
		if err != nil {
			fmt.Printf("select err:\n%+v\n", err)
		}
		fmt.Println(post)
	}
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

func (u *Post) AfterCreate(tx *gorm.DB) (err error) {

	ctx := context.Background()
	rows, err := gorm.G[User](tx).Where("id = ?", u.UserID).Update(ctx, "postcount", gorm.Expr("postcount + 1"))
	if err != nil {
		fmt.Printf("Update err:\n%+v\n", err)
	}
	fmt.Println(rows)
	return
}

type Commentcount struct {
	count uint
}

func (u *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Printf("%+v", u)
	msg := "无评论"
	var result map[string]interface{}
	//var result Commentcount
	ctx := context.Background()
	err = gorm.G[Comment](tx).Where("post_id = ?", u.PostID).Select("post_id, count(post_id) as count").Group("post_id").Scan(ctx, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", result)
	if result["count"] == 0 {
		gorm.G[Post](tx).Where("id = ?", u.PostID).Update(ctx, "Comments_count", msg)
		return
	}
	gorm.G[Post](tx).Where("id = ?", u.PostID).Update(ctx, "Comments_count", result["count"])
	return
}
