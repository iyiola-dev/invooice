package repository

import (
	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/db"
	"github.com/iyiola-dev/numeris/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	// User
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUsers(filters map[string]interface{}) ([]models.User, error)
	DeleteUser(id uuid.UUID) error

	// Customer
	CreateCustomer(customer *models.Customer) error
	GetCustomerByID(id uuid.UUID) (*models.Customer, error)
	GetCustomers(filters map[string]interface{}) ([]models.Customer, error)
	DeleteCustomer(id uuid.UUID) error

	// Invoice
	CreateInvoice(invoice *models.Invoice) error
	GetInvoiceByID(id uuid.UUID) (*models.Invoice, error)
	GetInvoices(filters map[string]interface{}) ([]models.Invoice, error)
	DeleteInvoice(id uuid.UUID) error

	// InvoiceItem
	CreateInvoiceItem(item *models.InvoiceItem) error
	GetInvoiceItems(invoiceID uuid.UUID) ([]models.InvoiceItem, error)
	DeleteInvoiceItem(id uuid.UUID) error

	// ActivityLog
	CreateActivityLog(log *models.ActivityLog) error
	GetActivityLogs(filters map[string]interface{}) ([]models.ActivityLog, error)

	// PaymentDetails
	CreatePaymentDetails(details *models.PaymentDetails) error
	GetPaymentDetailsByInvoiceID(invoiceID uuid.UUID) (*models.PaymentDetails, error)
	DeletePaymentDetails(id uuid.UUID) error


    // Update methods
    UpdateUser(id uuid.UUID, user *models.User) error
    UpdateCustomer(id uuid.UUID, customer *models.Customer) error
    UpdateInvoice(id uuid.UUID, invoice *models.Invoice) error
    UpdateInvoiceItem(id uuid.UUID, item *models.InvoiceItem) error
    UpdatePaymentDetails(id uuid.UUID, details *models.PaymentDetails) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &repository{db: db.DB}
}
