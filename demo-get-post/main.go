package main

/*
	Khai báo các Thư Việc cần có cho Service
*/
import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
	Khai báo Cấu Trúc cần có cho Service
*/
type ProductStruct struct {
	Id         string
	Name       string
	Descrition string
	Image      string
}

/*
	Khai báo biến Toàn Cục cho Service
*/
var productList []ProductStruct

/*
	Hàm khởi tạo
	Được chạy khi khởi động service
*/
func init() {
	log.Println("Khởi động Service DEMO Get/Post")
}

/*
	Hàm chính
	Mỗi service chỉ và phải có 1 hàm main
*/
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}
func main() {

	router := gin.Default()

	router.GET("/product/list", func(c *gin.Context) {
		c.JSON(200, productList)
	})
	router.POST("/product/add", func(c *gin.Context) {

		// Thu thập dữ liệu đầu vào dạng POST
		id := c.PostForm("id")
		name := c.PostForm("name")
		image := c.PostForm("image")

		// Kiểm tra tính hợp lệ của dữ liệu
		if id == "" {
			c.String(400, "`id` not empty")
			return
		}
		if name == "" {
			c.String(400, "`name` not empty")
			return
		}
		if image == "" {
			c.String(400, "`image` not empty")
			return
		}

		// Tạo 1 object có Cấu Trức là ProductStruct
		var product ProductStruct
		product.Id = id
		product.Name = name
		product.Image = image

		// Thêm vào danh sách toàn cục `productList`
		productList = append(productList, product)

		c.String(200, "add product success")
	})

	router.DELETE("/product/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product ProductStruct
		product.Id = id

		fmt.Println(productList)
		productList = remove(productList, id)
		productList = append(productList, product)

	})
	router.Run()
}
