package product

import (
	"errors"
)

func (this *ProductStruct) GetId() string {
	return this.Id
}

func (this *ProductStruct) GetName() string {
	return this.Name
}

func (this *ProductStruct) GetImage() string {
	return this.Image
}

func (this *ProductStruct) Len() int {
	return len(ProductList)
}

func (this *ProductStruct) GetAll() map[string]ProductStruct {
	return ProductList
}

func (this *ProductStruct) Add() error {
	// Tạo 1 object có Cấu Trức là ProductStruct
	var product ProductStruct
	product.Id = this.Id
	product.Name = this.Name
	product.Image = this.Image

	// Kiểm tra ID đã tồn tại thì báo lỗi 
	// ........


	// Thêm vào danh sách toàn cục `productList`
	ProductList[this.Id] = product
	return nil
}

func (this *ProductStruct) Delete() error {
	if _, ok := ProductList[this.Id] ; !ok {
		return errors.New("Không tồn tại ID: " + this.Id)
	}

	delete(ProductList , this.Id)
	return nil
}