package controllers

import (
	"fmt"
	"go-gin-contact-api/models"
	"go-gin-contact-api/services"
	"go-gin-contact-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type contactController struct {
	contactService services.ContactService
}

func NewContactController(contactService services.ContactService) *contactController {
	return &contactController{contactService}
}

func (h *contactController) GetAllContacs(ctx *gin.Context) {
	contacts, err := h.contactService.FindAllContacts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var contactsResponse []utils.GetContactResponse

	for _, c := range contacts {
		newContactResponse := convertToContactResponse(c)

		contactsResponse = append(contactsResponse, newContactResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": contactsResponse,
	})
}

func (h *contactController) GetContact(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.contactService.FindContactByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	contactResponse := convertToContactResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}

func (h *contactController) CreateContact(ctx *gin.Context) {
	var contactRequest utils.GetContactRequest

	err := ctx.ShouldBindJSON(&contactRequest)

	if err != nil {
		// log.Fatal(err)

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed &s, condition: %s"+e.Field(), e.ActualTag())

			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	contact, err := h.contactService.CreateContact(contactRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToContactResponse(contact),
	})
}

func (h *contactController) UpdateContact(ctx *gin.Context) {
	var contactRequest utils.GetContactRequest

	err := ctx.ShouldBindJSON(&contactRequest)

	if err != nil {
		// log.Fatal(err)

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed &s, condition: %s"+e.Field(), e.ActualTag())

			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	contact, err := h.contactService.UpdateContact(id, contactRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToContactResponse(contact),
	})
}

func (h *contactController) DeleteContact(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.contactService.DeleteContact(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	contactResponse := convertToContactResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": contactResponse,
	})
}

func convertToContactResponse(c models.Contact) utils.GetContactResponse {
	return utils.GetContactResponse{
		ID:      c.ID,
		Name:    c.Name,
		Phone:   c.Phone,
		Email:   c.Email,
		Address: c.Address,
	}
}
