package database

import (
	"blog/internal/model"
	"context"
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var dbchan chan interface{}

// func init() {
// 	dbchan = make(chan interface{})
// }

type DBHandler struct {
	sync.Mutex
	db *gorm.DB
}

var dbhandler DBHandler

func init() {
	dbhandler.db = SqliteOpen()
	dbhandler.db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}

func SqliteOpen() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AddUser(ctx context.Context, user *model.User) (err error) {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	err = gorm.G[model.User](dbhandler.db).Create(ctx, user)
	return err
}

func GetUser(ctx context.Context, username string) (*model.User, error) {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	finduser, err := gorm.G[model.User](dbhandler.db).Where("username=?", username).First(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return &finduser, err
}

func GetPosts(ctx context.Context, post model.Post) ([]model.Post, error) {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	return gorm.G[model.Post](dbhandler.db).Select("id, title, user_id").Where(post).Find(ctx)
}

func AddPost(ctx context.Context, post *model.Post) error {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	return gorm.G[model.Post](dbhandler.db).Create(ctx, post)
}

func UpdatePost(ctx context.Context, post *model.Post) error {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	_, err := gorm.G[model.Post](dbhandler.db).Updates(ctx, *post)
	return err
}

func DeletePost(ctx context.Context, post *model.Post) error {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	_, err := gorm.G[model.Post](dbhandler.db).Where(post).Delete(ctx)
	return err
}

func AddComment(ctx context.Context, comment *model.Comment) error {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	return gorm.G[model.Comment](dbhandler.db).Create(ctx, comment)
}

func GetComments(ctx context.Context, postid uint) ([]model.Comment, error) {
	dbhandler.Lock()
	defer dbhandler.Unlock()
	comments, err := gorm.G[model.Comment](dbhandler.db).Where("post_id=?", postid).Find(ctx)
	return comments, err
}
