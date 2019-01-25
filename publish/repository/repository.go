package repository

import (
	"github.com/amiraliio/publishen/publish/config"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/arangodb/go-driver"
)

//Adapter is publish repository interface
type Adapter interface {
	Create(*pb.Publish) (string, error)
	Get(*pb.Publish) (*pb.Publish, error)
	List(*pb.Request) ([]*pb.Publish, error)
}

//Repository publish
type Repository struct{}

//Create a publish into database
func (r *Repository) Create(p *pb.Publish) (string, error) {
	c, err := config.Collection()
	if err != nil {
		return "", err
	}
	doc, err := c.CreateDocument(nil, p)
	if err != nil {
		return "", err
	}
	return doc.ID.String(), nil
}

//Get a publish
func (r *Repository) Get(p *pb.Publish) (*pb.Publish, error) {
	c, err := config.Collection()
	if err != nil {
		return nil, err
	}
	var pub *pb.Publish
	publish, err := c.ReadDocument(nil, p.Key, &pub)
	if err != nil {
		return nil, err
	}
	pub.Key = publish.Key
	pub.ID = publish.ID.String()
	return pub, nil
}

//List of publishes
func (r *Repository) List(req *pb.Request) ([]*pb.Publish, error) {
	db, err := config.DB()
	if err != nil {
		return nil, err
	}
	query := "FOR p IN publishes RETURN p"
	cursor, err := db.Query(nil, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var publishes []*pb.Publish
	for {
		var pub *pb.Publish
		meta, err := cursor.ReadDocument(nil, &pub)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		pub.ID = meta.ID.String()
		pub.Key = meta.Key
		publishes = append(publishes, pub)
	}
	return publishes, nil
}
