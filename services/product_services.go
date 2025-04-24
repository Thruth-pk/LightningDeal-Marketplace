package services

import (
	"LightningDeal_Marketplace/datamodels"
	"LightningDeal_Marketplace/repositories"
)

type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProducts() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductService(productRepository repositories.IProduct) IProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) GetProductByID(id int64) (*datamodels.Product, error) {
	return p.productRepository.SelectByKey(id)
}
func (p *ProductService) GetAllProducts() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}
func (p *ProductService) DeleteProductByID(id int64) bool {
	return p.productRepository.Delete(id)
}
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}
