package main

/*
	Khai báo các Thư Việc cần có cho Service
*/
import (
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	Product "./module/product"
)

/*
	Hàm khởi tạo
	Được chạy khi khởi động service
*/
func init() {
	log.Println("Khởi động Service DEMO Get/Post")
}


func main() {

	router := gin.Default()

	Product.InitUrlProduct(router)

	router.Run()
}
