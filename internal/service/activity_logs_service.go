package service

import (
	"github.com/iyiola-dev/numeris/internal/models"
	"sort"
)

func (s *service) GetActivityLogs(filters map[string]interface{}) ([]models.ActivityLog, error) {
	// Get logs with preloaded relationships
	logs, err := s.repo.GetActivityLogs(filters)
	if err != nil {
		return nil, err
	}

	// Sort logs by timestamp in descending order (newest first)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Timestamp.After(logs[j].Timestamp)
	})

	return logs, nil
}
