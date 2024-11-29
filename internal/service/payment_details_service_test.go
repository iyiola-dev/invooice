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

func TestCreatePaymentDetails(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    invoiceID := uuid.New()
    input := inputs.CreatePaymentDetailsInput{
        InvoiceID:      invoiceID,
        AccountName:    "John Doe",
        AccountNumber:  "1234567890",
        BankName:      "Test Bank",
        BankAddress:   "123 Bank St",
        RoutingNumber: "987654321",
        PaymentDueDate: time.Now().AddDate(0, 0, 30),
    }

    mockRepo.On("CreatePaymentDetails", mock.AnythingOfType("*models.PaymentDetails")).Return(nil)

    details, err := svc.CreatePaymentDetails(input)

    assert.NoError(t, err)
    assert.NotNil(t, details)
    assert.Equal(t, input.AccountName, details.AccountName)
    assert.Equal(t, input.AccountNumber, details.AccountNumber)
    mockRepo.AssertExpectations(t)
}

func TestGetPaymentDetailsByInvoiceID(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    invoiceID := uuid.New()
    expected := &models.PaymentDetails{
        ID:           uuid.New(),
        InvoiceID:    invoiceID,
        AccountName:  "John Doe",
        AccountNumber: "1234567890",
    }

    mockRepo.On("GetPaymentDetailsByInvoiceID", invoiceID).Return(expected, nil)

    details, err := svc.GetPaymentDetailsByInvoiceID(invoiceID)

    assert.NoError(t, err)
    assert.Equal(t, expected, details)
    mockRepo.AssertExpectations(t)
}

func TestUpdatePaymentDetails(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    id := uuid.New()
    existingDetails := &models.PaymentDetails{
        ID:           id,
        AccountName:  "John Doe",
        AccountNumber: "1234567890",
    }

    updates := map[string]interface{}{
        "account_name":   "Jane Doe",
        "account_number": "0987654321",
        "bank_name":      "New Bank",
    }

    mockRepo.On("GetPaymentDetailsByInvoiceID", id).Return(existingDetails, nil)
    mockRepo.On("UpdatePaymentDetails", id, mock.AnythingOfType("*models.PaymentDetails")).Return(nil)

    err := svc.UpdatePaymentDetails(id, updates)

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestUpdatePaymentDetails_NotFound(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    id := uuid.New()
    mockRepo.On("GetPaymentDetailsByInvoiceID", id).Return(nil, errors.New("not found"))

    err := svc.UpdatePaymentDetails(id, map[string]interface{}{})

    assert.Error(t, err)
    mockRepo.AssertExpectations(t)
}

func TestDeletePaymentDetails(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    id := uuid.New()
    existingDetails := &models.PaymentDetails{
        ID: id,
    }

    mockRepo.On("GetPaymentDetailsByInvoiceID", id).Return(existingDetails, nil)
    mockRepo.On("DeletePaymentDetails", id).Return(nil)

    err := svc.DeletePaymentDetails(id)

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestDeletePaymentDetails_NotFound(t *testing.T) {
    mockRepo := new(mocks.Repository)
    svc := service.NewService(mockRepo)

    id := uuid.New()
    mockRepo.On("GetPaymentDetailsByInvoiceID", id).Return(nil, errors.New("not found"))

    err := svc.DeletePaymentDetails(id)

    assert.Error(t, err)
    mockRepo.AssertExpectations(t)
}