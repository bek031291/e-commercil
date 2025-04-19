package product

import (
	"ecommerce/internal/db"
	"ecommerce/models"
	"log"
)

func CreateProduct(product *models.Product) error {
	var id int64
	query := `INSERT INTO products (name, brand, description, price, category_id, created_at, updated_at) 
			VALUES (?, ?, ?, ?, ?, NOW(), NOW()) RETURNING id`
	err := db.DB.Raw(query, product.Name, product.Brand, product.Description, product.Price, product.CategoryID).Scan(id)

	return err.Error
}

func GetProductByID(id uint) (*models.Product, error) {

	var product models.Product
	query := "SELECT * FROM products WHERE id = ?"
	err := db.DB.Raw(query, id).Scan(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, err
}

func UpdateProduct(id uint, updated *models.Product) error {
	query := `update products set name=?,brand=?,description=?,price=?,category_id=?,updated_at=now() where id=?`
	err := db.DB.Exec(query, updated.Name, updated.Brand, updated.Description, updated.Price, updated.CategoryID, id)
	log.Println(err)
	return err.Error
}

func DeleteProduct(id uint) error {
	result := db.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAll() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT * FROM products"
	err := db.DB.Raw(query).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
