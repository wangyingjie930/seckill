package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" imooc:"ID"` //自己设置的标签imooc
	ProductName  string `json:"ProductName" sql:"productName" imooc:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" imooc:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" imooc:"ProductUrl"`
}
