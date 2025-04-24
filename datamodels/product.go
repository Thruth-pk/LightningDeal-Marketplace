package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" view:"id"`
	ProductName  string `json:"ProductName" sql:"product_name" view:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"product_num" view:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"product_img" view:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"product_url" view:"ProductUrl"`
}
