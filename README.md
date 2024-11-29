# Numeris - Invoice Management System

A robust invoice management system built with Go, using Gin framework and PostgreSQL.

## Features

- **Authentication**
  - User registration and login
  - JWT-based authentication

- **Invoice Management**
  - Create, read, update, and delete invoices
  - Add invoice items with descriptions, quantities, and prices
  - Calculate subtotals, discounts, and total amounts
  - Track invoice status (pending, paid, etc.)
  - Support for multiple currencies

- **Payment Details**
  - Add bank account details for payments
  - Track payment due dates
  - Update payment information
  - Link payment details to specific invoices

- **Activity Logging**
  - Track all invoice-related activities
  - Record user actions (create, update, delete)
  - Timestamp all activities
  - Filter logs by user and invoice

## Technology Stack

- **Backend**: Go (Gin Framework)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Environment**: godotenv

## Project Structure

.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── db/              # Database connection
│   ├── handlers/        # HTTP request handlers
│   ├── inputs/          # Request input
│   ├── models/          # Database models
│   ├── repository/      # Data access layer
│   ├── response/        # Response structures
│   ├── routes/          # Route definitions
│   ├── service/         # Business logic
│   └── util/            # Utilities and middleware
└── .env                 # Environment variables

## Models

### Invoice
- Primary model for storing invoice information
- Contains relationships with User, Customer, and InvoiceItems
- Tracks financial details like subtotal, discount, and total amount
- Supports multiple currencies and status tracking

### InvoiceItem
- Represents individual line items in an invoice
- Tracks description, quantity, unit price, and total amount
- Links back to parent invoice

### PaymentDetails
- Stores bank account and payment information
- Links to specific invoices
- Includes account name, number, bank details, and routing information
- Tracks payment due dates

### ActivityLog
- Tracks all system activities
- Records user actions on invoices
- Stores timestamps for audit trails
- Links to both users and invoices

