package repository

import (
	"github.com/amiraliio/publishen/publish/config"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/arangodb/go-driver"
)

//Adapter is publish repository interface
type Adapter interface {
	Create(*pb.Publish) (*pb.Publish, error)
	Get(*pb.Publish) (*pb.Publish, error)
	ListOfUserPublishes(*pb.Request) ([]*pb.Publish, error)
	Update(*pb.Publish) (*pb.Publish, error)
	Delete(*pb.Publish) (*pb.Publish, error)
}

//Repository publish
type Repository struct{}

//Create a publish into database
func (r *Repository) Create(p *pb.Publish) (*pb.Publish, error) {
	c, err := config.Collection()
	if err != nil {
		return nil, err
	}
	meta, err := c.CreateDocument(nil, p)
	if err != nil {
		return nil, err
	}
	p.Id = meta.ID.String()
	p.Key = meta.Key
	return p, nil
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
	pub.Id = publish.ID.String()
	return pub, nil
}

//ListOfUserPublishes of publishes
func (r *Repository) ListOfUserPublishes(req *pb.Request) ([]*pb.Publish, error) {
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
		pub.Id = meta.ID.String()
		pub.Key = meta.Key
		publishes = append(publishes, pub)
	}
	return publishes, nil
}

//Update a Publish
func (r *Repository) Update(p *pb.Publish) (*pb.Publish, error) {
	collection, err := config.Collection()
	if err != nil {
		return nil, err
	}
	meta, err := collection.UpdateDocument(nil, p.Key, p)
	if err != nil {
		return nil, err
	}
	var pub *pb.Publish
	pub.Id = meta.ID.String()
	pub.Key = meta.Key
	return pub, nil
}

//Delete a publish
func (r *Repository) Delete(p *pb.Publish) (*pb.Publish, error) {
	collection, err := config.Collection()
	if err != nil {
		return nil, err
	}
	meta, err := collection.RemoveDocument(nil, p.Key)
	if err != nil {
		return nil, err
	}
	var pub *pb.Publish
	pub.Id = meta.ID.String()
	pub.Key = meta.Key
	return pub, nil
}
