package repository

import (
	config "github.com/amirallio/publishen/product/config"
	pb "github.com/amirallio/publishen/product/model/product"
)

//Adapter is product repository interface
type Adapter interface {
	Create(*pb.Product) (string, error)
}

//Repository product
type Repository struct{}

//Create a product into database
func (r *Repository) Create(p *pb.Product) (string, error) {
	db, err := config.DB()
	if err != nil {
		return "", err
	}
	c, err := db.Collection(nil, "products")
	if err != nil {
		return "", err
	}
	doc, err := c.CreateDocument(nil, p)
	if err != nil {
		return "", err
	}
	return doc.ID.String(), nil
}
