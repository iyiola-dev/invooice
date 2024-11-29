package service

import (
	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/repository"
	"github.com/iyiola-dev/numeris/internal/response"
)

type Service interface {
	  // Auth
	  Register(input inputs.RegisterInput) (*models.User, error)
	  Login(input inputs.LoginInput) (*response.LoginResponse, error)
	  
	
	// Invoice
	CreateInvoice(input inputs.CreateInvoiceInput) (*models.Invoice, error)
	GetInvoiceByID(id uuid.UUID) (*models.Invoice, error)
	GetInvoices(filters map[string]interface{}) ([]models.Invoice, error)
	UpdateInvoice(id uuid.UUID, updates map[string]interface{}) error
	DeleteInvoice(id uuid.UUID) error
	GetInvoiceWithItems(id uuid.UUID) (*models.Invoice, error)
	
	// Payment Details
	CreatePaymentDetails(input inputs.CreatePaymentDetailsInput) (*models.PaymentDetails, error)
	GetPaymentDetailsByInvoiceID(invoiceID uuid.UUID) (*models.PaymentDetails, error)
	UpdatePaymentDetails(id uuid.UUID, updates map[string]interface{}) error
	DeletePaymentDetails(id uuid.UUID) error
	
	// Activity Logs
	GetActivityLogs(filters map[string]interface{}) ([]models.ActivityLog, error)
}


type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
	
