package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "github.com/JuacCrack/GOapi/models"
)

var contacts = []models.Contact{
    {ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
}

// GetContacts godoc
// @Summary Get all contacts
// @Description Get a list of all contacts
// @Tags contacts
// @Produce json
// @Success 200 {array} models.Contact
// @Router /contacts [get]
func GetContacts(c *gin.Context) {
    c.JSON(http.StatusOK, contacts)
}

// GetContact godoc
// @Summary Get a contact by ID
// @Description Get a single contact by its ID
// @Tags contacts
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {object} models.Contact
// @Failure 404 {object} models.ErrorResponse
// @Router /contacts/{id} [get]
func GetContact(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, contact := range contacts {
        if contact.ID == id {
            c.JSON(http.StatusOK, contact)
            return
        }
    }
    c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Contact not found"})
}

// CreateContact godoc
// @Summary Create a new contact
// @Description Create a new contact with the input payload
// @Tags contacts
// @Accept json
// @Produce json
// @Param contact body models.Contact true "Contact"
// @Success 201 {object} models.Contact
// @Failure 400 {object} models.ErrorResponse
// @Router /contacts [post]
func CreateContact(c *gin.Context) {
    var newContact models.Contact
    if err := c.BindJSON(&newContact); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
        return
    }
    newContact.ID = len(contacts) + 1
    contacts = append(contacts, newContact)
    c.JSON(http.StatusCreated, newContact)
}

// UpdateContact godoc
// @Summary Update an existing contact
// @Description Update an existing contact by ID with the input payload
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path int true "Contact ID"
// @Param contact body models.Contact true "Contact"
// @Success 200 {object} models.Contact
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /contacts/{id} [put]
func UpdateContact(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedContact models.Contact
    if err := c.BindJSON(&updatedContact); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
        return
    }
    for i, contact := range contacts {
        if contact.ID == id {
            contacts[i] = updatedContact
            contacts[i].ID = id
            c.JSON(http.StatusOK, contacts[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Contact not found"})
}

// DeleteContact godoc
// @Summary Delete a contact by ID
// @Description Delete a single contact by its ID
// @Tags contacts
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /contacts/{id} [delete]
func DeleteContact(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, contact := range contacts {
        if contact.ID == id {
            contacts = append(contacts[:i], contacts[i+1:]...)
            c.JSON(http.StatusOK, models.ErrorResponse{Message: "Contact deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Contact not found"})
}
