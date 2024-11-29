package service_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/mocks"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetActivityLogs(t *testing.T) {
	// Create mock repository
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	// Test data
	userID := uuid.New()
	invoiceID := uuid.New()
	now := time.Now()

	testLogs := []models.ActivityLog{
		{
			ID:        uuid.New(),
			UserID:    userID,
			InvoiceID: &invoiceID,
			Action:    "INVOICE_CREATED",
			Timestamp: now.Add(-time.Hour), // 1 hour ago
		},
		{
			ID:        uuid.New(),
			UserID:    userID,
			InvoiceID: &invoiceID,
			Action:    "INVOICE_UPDATED",
			Timestamp: now, // current time
		},
	}

	filters := map[string]interface{}{
		"user_id": userID,
	}

	// Set up expectations
	mockRepo.On("GetActivityLogs", filters).Return(testLogs, nil)

	// Execute the service method
	logs, err := svc.GetActivityLogs(filters)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, logs)
	assert.Len(t, logs, 2)

	// Verify logs are sorted by timestamp (newest first)
	assert.Equal(t, "INVOICE_UPDATED", logs[0].Action)
	assert.Equal(t, "INVOICE_CREATED", logs[1].Action)

	// Verify mock expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetActivityLogs_WithError(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	filters := map[string]interface{}{
		"user_id": uuid.New(),
	}

	// Set up expectations for error case
	mockRepo.On("GetActivityLogs", filters).Return(nil, assert.AnError)

	// Execute the service method
	logs, err := svc.GetActivityLogs(filters)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, logs)
	mockRepo.AssertExpectations(t)
}

func TestGetActivityLogs_WithInvoiceFilter(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	userID := uuid.New()
	invoiceID := uuid.New()
	now := time.Now()

	testLogs := []models.ActivityLog{
		{
			ID:        uuid.New(),
			UserID:    userID,
			InvoiceID: &invoiceID,
			Action:    "INVOICE_CREATED",
			Timestamp: now,
		},
	}

	filters := map[string]interface{}{
		"user_id":    userID,
		"invoice_id": invoiceID,
	}

	// Set up expectations
	mockRepo.On("GetActivityLogs", filters).Return(testLogs, nil)

	// Execute the service method
	logs, err := svc.GetActivityLogs(filters)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, logs)
	assert.Len(t, logs, 1)
	assert.Equal(t, invoiceID, *logs[0].InvoiceID)
	mockRepo.AssertExpectations(t)
}
