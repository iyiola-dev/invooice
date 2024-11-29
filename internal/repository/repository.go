package repository

import (
	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/models"
)


// User implementations
func (r *repository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *repository) GetUsers(filters map[string]interface{}) ([]models.User, error) {
	var users []models.User
	err := r.db.Where(filters).Find(&users).Error
	return users, err
}

func (r *repository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

// Customer implementations
func (r *repository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *repository) GetCustomerByID(id uuid.UUID) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, "id = ?", id).Error
	return &customer, err
}

func (r *repository) GetCustomers(filters map[string]interface{}) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Where(filters).Find(&customers).Error
	return customers, err
}

func (r *repository) DeleteCustomer(id uuid.UUID) error {
	return r.db.Delete(&models.Customer{}, "id = ?", id).Error
}

// Invoice implementations
func (r *repository) CreateInvoice(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *repository) GetInvoiceByID(id uuid.UUID) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.Preload("Customer").
		Preload("User").
		Preload("Items").
		First(&invoice, "id = ?", id).Error
	return &invoice, err
}

func (r *repository) GetInvoices(filters map[string]interface{}) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Preload("Customer").
		Preload("User").
		Preload("Items").
		Where(filters).
		Order("created_at DESC").
		Find(&invoices).Error
	return invoices, err
}

func (r *repository) DeleteInvoice(id uuid.UUID) error {
	return r.db.Delete(&models.Invoice{}, "id = ?", id).Error
}

// InvoiceItem implementations
func (r *repository) CreateInvoiceItem(item *models.InvoiceItem) error {
	return r.db.Create(item).Error
}

func (r *repository) GetInvoiceItems(invoiceID uuid.UUID) ([]models.InvoiceItem, error) {
	var items []models.InvoiceItem
	err := r.db.Where("invoice_id = ?", invoiceID).Find(&items).Error
	return items, err
}

func (r *repository) DeleteInvoiceItem(id uuid.UUID) error {
	return r.db.Delete(&models.InvoiceItem{}, "id = ?", id).Error
}

// ActivityLog implementations
func (r *repository) CreateActivityLog(log *models.ActivityLog) error {
	return r.db.Create(log).Error
}

func (r *repository) GetActivityLogs(filters map[string]interface{}) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	err := r.db.Preload("User").Preload("Invoice").Where(filters).Find(&logs).Error
	return logs, err
}

// PaymentDetails implementations
func (r *repository) CreatePaymentDetails(details *models.PaymentDetails) error {
	return r.db.Create(details).Error
}

func (r *repository) GetPaymentDetailsByInvoiceID(invoiceID uuid.UUID) (*models.PaymentDetails, error) {
	var details models.PaymentDetails
	err := r.db.Where("invoice_id = ?", invoiceID).First(&details).Error
	return &details, err
}

func (r *repository) DeletePaymentDetails(id uuid.UUID) error {
	return r.db.Delete(&models.PaymentDetails{}, "id = ?", id).Error
}

// Update methods
func (r *repository) UpdateUser(id uuid.UUID, user *models.User) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *repository) UpdateCustomer(id uuid.UUID, customer *models.Customer) error {
	return r.db.Model(&models.Customer{}).Where("id = ?", id).Updates(customer).Error
}

func (r *repository) UpdateInvoice(id uuid.UUID, invoice *models.Invoice) error {
	return r.db.Model(&models.Invoice{}).Where("id = ?", id).Updates(invoice).Error
}

func (r *repository) UpdateInvoiceItem(id uuid.UUID, item *models.InvoiceItem) error {
	return r.db.Model(&models.InvoiceItem{}).Where("id = ?", id).Updates(item).Error
}

func (r *repository) UpdatePaymentDetails(id uuid.UUID, details *models.PaymentDetails) error {
	return r.db.Model(&models.PaymentDetails{}).Where("id = ?", id).Updates(details).Error
}