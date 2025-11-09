package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Age   uint
	Grade string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:5594210@tcp(127.0.0.1:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}

	ctx := context.Background()
	fmt.Printf("%v\n", ctx)
	//migrate the schema
	db.AutoMigrate(&Student{})

	err = gorm.G[Student](db).Create(ctx, &Student{Name: "张三", Age: 20, Grade: "三年级"})
	if err != nil {
		fmt.Printf("Create err:\n%+v\n", err)
	}

	products, err := gorm.G[Student](db).Where("age > ?", 18).Find(ctx)
	if err == nil {
		fmt.Printf("product:\n%+v\n", products)
	} else {
		fmt.Printf("select err:\n%+v\n", err)
	}

	rows, err := gorm.G[Student](db).Where("name = ?", "张三").Update(ctx, "grade", "四年级")
	if err != nil {
		fmt.Printf("Update err:\n%+v\n", err)
	} else {
		fmt.Printf("rows:\n%+v\n", rows)
	}

	_, err = gorm.G[Student](db).Where("age < ?", 15).Delete(ctx)
	if err != nil {
		fmt.Printf("Delete err:\n%+v\n", err)
	}

	err = Trans(db, 1, 2, 100)

	if err != nil {
		fmt.Printf("Trans err:\n%+v\n", err)
	}
}
