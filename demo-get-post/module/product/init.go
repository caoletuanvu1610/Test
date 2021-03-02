package product

import (
	"log"
	"github.com/gin-gonic/gin"
)

/*
	Khai báo biến Toàn Cục cho Moduel
*/
var ProductList = make(map[string]ProductStruct, 0)

/*
	Hàm khởi tạo
	Được chạy khi khởi động service
*/
func init() {
	log.Println("Khởi động Module Product")
}



func InitUrlProduct(router *gin.Engine) {

	url := router.Group("/product")
	{

		/**
			Lấy Danh sách Sản Phẩm
			....
		*/
		url.GET("/list", func(c *gin.Context) {

			var ProductObj ProductStruct
			c.JSON(200, ProductObj.GetAll())

		})

		/**
			Thêm SP vào danh sách
			....
		*/
		url.POST("/add", func(c *gin.Context) {
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

			var ProductObj ProductStruct
			ProductObj.Id = id
			ProductObj.Name = name
			ProductObj.Image = image
			err := ProductObj.Add()
			if err != nil {
				c.String(400, err.Error())
				return
			}

			c.String(200, "add product success")
		})

		/**
			Bỏ sản phẩm ra khỏi danh sách, theo ID
			....
		*/
		url.DELETE("/delete/:id", func(c *gin.Context) {
			id := c.Param("id")
			var ProductObj ProductStruct
			ProductObj.Id = id

			err := ProductObj.Delete()
			if err != nil {
				c.String(400, err.Error())
				return
			}

			c.String(200, "delete product success")
		})
	}
}
