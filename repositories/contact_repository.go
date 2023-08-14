package repositories

import (
	"go-gin-contact-api/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	FindAllContacts() ([]models.Contact, error)
	FindContactByID(ID int) (models.Contact, error)
	CreateContact(contact models.Contact) (models.Contact, error)
	UpdateContact(contact models.Contact) (models.Contact, error)
	DeleteContact(contact models.Contact) (models.Contact, error)
}

type repository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllContacts() ([]models.Contact, error) {
	var contact []models.Contact
	err := r.db.Find(&contact).Error
	return contact, err
}

func (r *repository) FindContactByID(ID int) (models.Contact, error) {
	var contact models.Contact
	err := r.db.Find(&contact, ID).Error
	return contact, err
}

func (r *repository) CreateContact(contact models.Contact) (models.Contact, error) {
	err := r.db.Create(&contact).Error
	return contact, err
}

func (r *repository) UpdateContact(contact models.Contact) (models.Contact, error) {
	err := r.db.Save(&contact).Error
	return contact, err
}

func (r *repository) DeleteContact(contact models.Contact) (models.Contact, error) {
	err := r.db.Delete(&contact).Error
	return contact, err
}
