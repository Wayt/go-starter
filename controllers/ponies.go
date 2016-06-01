package controllers

import (
	"github.com/wayt/go-starter/models"
)

type ponies struct {
	controller
}

// Ponies controller
var Ponies = ponies{register(models.Pony{})}

func (c ponies) Get(id string) (*models.Pony, error) {

	p := new(models.Pony)
	if err := db.First(p, id).Error; err != nil {
		return nil, wrap(err)
	}
	return p, nil
}

func (c ponies) List() (models.Ponies, error) {

	var ps models.Ponies
	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}

	return ps, nil
}

func (c ponies) Create(name, title string) (*models.Pony, error) {

	p := &models.Pony{
		Name:  name,
		Title: title,
	}

	if err := db.Create(p).Error; err != nil {
		return nil, wrap(err)
	}

	return p, nil
}

func (c ponies) Delete(p *models.Pony) (*models.Pony, error) {

	return p, db.Delete(p).Error
}
