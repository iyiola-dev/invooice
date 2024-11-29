package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/models"
)


func (s *service) CreatePaymentDetails(input inputs.CreatePaymentDetailsInput) (*models.PaymentDetails, error) {
    // Validate invoice exists
    invoice, err := s.repo.GetInvoiceByID(input.InvoiceID)
    if err != nil {
        return nil, errors.New("invalid invoice")
    }

    // Check if payment details already exist
    existing, err := s.repo.GetPaymentDetailsByInvoiceID(input.InvoiceID)
    if err == nil && existing != nil {
        return nil, errors.New("payment details already exist for this invoice")
    }

    details := &models.PaymentDetails{
        ID:            uuid.New(),
        InvoiceID:     invoice.ID,
        AccountName:   input.AccountName,
        AccountNumber: input.AccountNumber,
        BankName:      input.BankName,
        BankAddress:   input.BankAddress,
        RoutingNumber: input.RoutingNumber,
        PaymentDueDate: input.PaymentDueDate,
    }

    err = s.repo.CreatePaymentDetails(details)
    if err != nil {
        return nil, err
    }

    return details, nil
}

func (s *service) GetPaymentDetailsByInvoiceID(invoiceID uuid.UUID) (*models.PaymentDetails, error) {
    return s.repo.GetPaymentDetailsByInvoiceID(invoiceID)
}

func (s *service) UpdatePaymentDetails(id uuid.UUID, updates map[string]interface{}) error {
    details, err := s.repo.GetPaymentDetailsByInvoiceID(id)
    if err != nil {
        return err
    }

    for key, value := range updates {
        switch key {
        case "account_name":
            details.AccountName = value.(string)
        case "account_number":
            details.AccountNumber = value.(string)
        case "bank_name":
            details.BankName = value.(string)
        case "bank_address":
            details.BankAddress = value.(string)
        case "routing_number":
            details.RoutingNumber = value.(string)
        case "payment_due_date":
            details.PaymentDueDate = value.(time.Time)
        }
    }

    return s.repo.UpdatePaymentDetails(id, details)
}

func (s *service) DeletePaymentDetails(id uuid.UUID) error {
    _, err := s.repo.GetPaymentDetailsByInvoiceID(id)
    if err != nil {
        return err
    }
    return s.repo.DeletePaymentDetails(id)
}
