package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//Create
	//db.Create(&Product{
	//	Name:  "Notebook",
	//	Price: 1000.99,
	//})

	//Create batch
	//products := []Product{
	//	{Name: "Notebook", Price: 1000.99},
	//	{Name: "Mouse", Price: 50.00},
	//	{Name: "Keyboard", Price: 100.00},
	//}
	//db.Create(&products)

	//Select one
	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)
	//db.First(&product, "name = ?, Mouse")
	//fmt.Println(product)

	//Selact all
	//var products2 []Product
	//db.Find(&products2)
	//for _, product := range products2 {
	//	fmt.Println(product)
	//}

	//var products []Product
	//db.Limit(20).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//Where
	//var products []Product
	//db.Where("Price = ?", 100).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//Like
	var products []Product
	db.Where("name LIKE ?", "%book").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}
