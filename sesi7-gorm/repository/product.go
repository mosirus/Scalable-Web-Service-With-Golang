package repository

import (
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product *models.Product) error
	GetAllProduct() (*[]models.Product, error)

	//Latihan
	UpdateProduct(*models.Product) error
	DeleteProduct(id uint) (*models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) GetAllProduct() (*[]models.Product, error) {
	var products []models.Product
	err := p.db.Find(&products).Error
	return &products, err
}

//latihan
func (r *productRepo) UpdateProduct(request *models.Product) error {
	err := r.db.Exec("UPDATE products SET name = ?, brand = ?, user_id = ? WHERE id = ?", request.Name, request.Brand, request.UserID, request.ID).Error
	return err
}

func (r *productRepo) DeleteProduct(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id=?", id).Delete(&product).Error
	return &product, err
}
