package repository

import (
	"github.com/vmware/vic/lib/apiservers/engine/errors"
	"github.com/amiraliio/publishen/publish/config"
	pb "github.com/amiraliio/publishen/publish/model/publish"
)

//Adapter is publish repository interface
type Adapter interface {
	Create(*pb.Publish) (string, error)
	Get(*pb.Publish) (*pb.Publish, error)
	List(*pb.Request) ([]*pb.Publishes, error)
}

//Repository publish
type Repository struct{}

const collection string = "publishes"

//Create a publish into database
func (r *Repository) Create(p *pb.Publish) (string, error) {
	db, err := config.DB()
	if err != nil {
		return "", err
	}
	c, err := db.Collection(nil, collection)
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
	var pub *pb.Publish
	db, err := config.DB()
	if err != nil {
		return nil, err
	}
	collection, err := db.Collection(nil, collection)
	if err != nil {
		return nil, err
	}
	publish, err := collection.ReadDocument(nil, p.Key, &pub)
	if err != nil {
		return nil, err
	}
	pub.Key = publish.Key
	pub.ID = publish.ID.String()
	return pub, nil
}

func (r *Repository) List(req *pb.Request) ([]*pb.Publishes, error){
	//TODO
}