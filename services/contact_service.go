package services

import (
	"go-gin-contact-api/models"
	"go-gin-contact-api/repositories"
	"go-gin-contact-api/utils"
)

type ContactService interface {
	FindAllContacts() ([]models.Contact, error)
	FindContactByID(ID int) (models.Contact, error)
	CreateContact(contactRequest utils.GetContactRequest) (models.Contact, error)
	UpdateContact(ID int, contactRequest utils.GetContactRequest) (models.Contact, error)
	DeleteContact(ID int) (models.Contact, error)
}

type service struct {
	repository repositories.ContactRepository
}

func NewContactService(repository repositories.ContactRepository) *service {
	return &service{repository}
}

func (s *service) FindAllContacts() ([]models.Contact, error) {
	contacts, err := s.repository.FindAllContacts()
	return contacts, err
	// return s.repository.FindAll()
}

func (s *service) FindContactByID(ID int) (models.Contact, error) {
	contact, err := s.repository.FindContactByID(ID)
	return contact, err
}

func (s *service) CreateContact(contactRequest utils.GetContactRequest) (models.Contact, error) {

	contact := models.Contact{
		Name:    contactRequest.Name,
		Phone:   contactRequest.Phone,
		Email:   contactRequest.Email,
		Address: contactRequest.Address,
	}

	newContact, err := s.repository.CreateContact(contact)
	return newContact, err
}

func (s *service) UpdateContact(ID int, contactRequest utils.GetContactRequest) (models.Contact, error) {
	contact, err := s.repository.FindContactByID(ID)

	contact.Name = contactRequest.Name
	contact.Phone = contactRequest.Phone
	contact.Email = contactRequest.Email
	contact.Address = contactRequest.Address

	newContact, err := s.repository.UpdateContact(contact)
	return newContact, err
}

func (s *service) DeleteContact(ID int) (models.Contact, error) {
	contact, err := s.repository.FindContactByID(ID)

	newBook, err := s.repository.DeleteContact(contact)
	return newBook, err
}
