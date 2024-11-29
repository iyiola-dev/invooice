package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/models"
)

func (s *service) CreateInvoice(input inputs.CreateInvoiceInput) (*models.Invoice, error) {
	// Validate customer exists
	customer, err := s.repo.GetCustomerByID(input.CustomerID)
	if err != nil {
		return nil, errors.New("invalid customer")
	}

	// Create invoice
	invoice := &models.Invoice{
		ID:            uuid.New(),
		UserID:        input.UserID,
		CustomerID:    customer.ID,
		InvoiceNumber: input.InvoiceNumber,
		IssueDate:     input.IssueDate,
		DueDate:       input.DueDate,
		Currency:      input.Currency,
		SubTotal:      input.SubTotal,
		Discount:      input.Discount,
		TotalAmount:   input.TotalAmount,
		Status:        "pending",
		Note:          input.Note,
	}

	err = s.repo.CreateInvoice(invoice)
	if err != nil {
		return nil, err
	}

	// Create invoice items
	for _, item := range input.Items {
		invoiceItem := &models.InvoiceItem{
			InvoiceID:   invoice.ID,
			Description: item.Description,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			Amount:      item.Amount,
		}
		err = s.repo.CreateInvoiceItem(invoiceItem)
		if err != nil {
			return nil, err
		}
	}

	// Create activity log
	activityLog := &models.ActivityLog{
		UserID:    input.UserID,
		InvoiceID: &invoice.ID,
		Action:    "INVOICE_CREATED",
		Timestamp: time.Now(),
	}

	_ = s.repo.CreateActivityLog(activityLog)

	return invoice, nil
}

func (s *service) GetInvoiceByID(id uuid.UUID) (*models.Invoice, error) {
	return s.repo.GetInvoiceByID(id)
}

func (s *service) GetInvoices(filters map[string]interface{}) ([]models.Invoice, error) {
	return s.repo.GetInvoices(filters)
}

func (s *service) UpdateInvoice(id uuid.UUID, updates map[string]interface{}) error {
	invoice, err := s.repo.GetInvoiceByID(id)
	if err != nil {
		return err
	}

	for key, value := range updates {
		switch key {
		case "status":
			invoice.Status = value.(string)
		case "note":
			invoice.Note = value.(string)
		case "due_date":
			invoice.DueDate = value.(time.Time)
		}
	}

	// Create activity log
	activityLog := &models.ActivityLog{
		UserID:    invoice.UserID,
		InvoiceID: &invoice.ID,
		Action:    "INVOICE_UPDATED",
		Timestamp: time.Now(),
	}

	_ = s.repo.CreateActivityLog(activityLog)

	return s.repo.UpdateInvoice(id, invoice)
}

func (s *service) DeleteInvoice(id uuid.UUID) error {
	invoice, err := s.repo.GetInvoiceByID(id)
	if err != nil {
		return err
	}

	// Create activity log
	activityLog := &models.ActivityLog{
		UserID:    invoice.UserID,
		InvoiceID: &invoice.ID,
		Action:    "INVOICE_DELETED",
		Timestamp: time.Now(),
	}

	_ = s.repo.CreateActivityLog(activityLog)

	return s.repo.DeleteInvoice(id)
}

func (s *service) GetInvoiceWithItems(id uuid.UUID) (*models.Invoice, error) {
	// GetInvoiceByID already preloads Items, User, and Customer
	invoice, err := s.repo.GetInvoiceByID(id)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
