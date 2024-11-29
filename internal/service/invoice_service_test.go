package service_test

import (
	"testing"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/mocks"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	customerID := uuid.New()
	userID := uuid.New()

	customer := &models.Customer{
		ID: customerID,
	}

	input := inputs.CreateInvoiceInput{
		CustomerID:    customerID,
		UserID:       userID,
		InvoiceNumber: "INV-001",
		IssueDate:    time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 30),
		Currency:     "USD",
		Items: []inputs.CreateInvoiceItemInput{
			{
				Description: "Test Item",
				Quantity:    1,
				UnitPrice:   100,
				Amount:     100,
			},
		},
		SubTotal:    100,
		TotalAmount: 100,
	}

	// Set up expectations
	mockRepo.On("GetCustomerByID", customerID).Return(customer, nil)
	mockRepo.On("CreateInvoice", mock.AnythingOfType("*models.Invoice")).Return(nil)
	mockRepo.On("CreateInvoiceItem", mock.AnythingOfType("*models.InvoiceItem")).Return(nil)
	mockRepo.On("CreateActivityLog", mock.AnythingOfType("*models.ActivityLog")).Return(nil)

	// Execute
	invoice, err := svc.CreateInvoice(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, invoice)
	assert.Equal(t, customerID, invoice.CustomerID)
	assert.Equal(t, userID, invoice.UserID)
	mockRepo.AssertExpectations(t)
}

func TestCreateInvoice_CustomerNotFound(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	input := inputs.CreateInvoiceInput{
		CustomerID: uuid.New(),
	}

	mockRepo.On("GetCustomerByID", input.CustomerID).Return(nil, errors.New("not found"))

	invoice, err := svc.CreateInvoice(input)

	assert.Error(t, err)
	assert.Nil(t, invoice)
	assert.Equal(t, "invalid customer", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetInvoiceByID(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	invoiceID := uuid.New()
	expected := &models.Invoice{
		ID: invoiceID,
	}

	mockRepo.On("GetInvoiceByID", invoiceID).Return(expected, nil)

	invoice, err := svc.GetInvoiceByID(invoiceID)

	assert.NoError(t, err)
	assert.Equal(t, expected, invoice)
	mockRepo.AssertExpectations(t)
}

func TestUpdateInvoice(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	invoiceID := uuid.New()
	userID := uuid.New()
	existingInvoice := &models.Invoice{
		ID:     invoiceID,
		UserID: userID,
		Status: "pending",
	}

	updates := map[string]interface{}{
		"status": "paid",
		"note":   "Payment received",
	}

	mockRepo.On("GetInvoiceByID", invoiceID).Return(existingInvoice, nil)
	mockRepo.On("CreateActivityLog", mock.AnythingOfType("*models.ActivityLog")).Return(nil)
	mockRepo.On("UpdateInvoice", invoiceID, mock.AnythingOfType("*models.Invoice")).Return(nil)

	err := svc.UpdateInvoice(invoiceID, updates)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteInvoice(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	invoiceID := uuid.New()
	userID := uuid.New()
	existingInvoice := &models.Invoice{
		ID:     invoiceID,
		UserID: userID,
	}

	mockRepo.On("GetInvoiceByID", invoiceID).Return(existingInvoice, nil)
	mockRepo.On("CreateActivityLog", mock.AnythingOfType("*models.ActivityLog")).Return(nil)
	mockRepo.On("DeleteInvoice", invoiceID).Return(nil)

	err := svc.DeleteInvoice(invoiceID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
