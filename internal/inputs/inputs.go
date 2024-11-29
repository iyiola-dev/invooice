package inputs

import (
	"time"

	"github.com/google/uuid"
)

type RegisterInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Address   string
}

type LoginInput struct {
	Email    string
	Password string
}


type CreateInvoiceInput struct {
    UserID        uuid.UUID
    CustomerID    uuid.UUID
    InvoiceNumber string
    IssueDate     time.Time
    DueDate       time.Time
    Currency      string
    SubTotal      float64
    Discount      float64
    TotalAmount   float64
    Note          string
    Items         []CreateInvoiceItemInput
}

type CreateInvoiceItemInput struct {
    Description string
    Quantity    int
    UnitPrice   float64
    Amount      float64
}


type CreatePaymentDetailsInput struct {
    InvoiceID      uuid.UUID
    AccountName    string
    AccountNumber  string
    BankName       string
    BankAddress    string
    RoutingNumber  string
    PaymentDueDate time.Time
}
