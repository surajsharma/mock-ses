package stats

import (
	"mock-ses/email"
	"sync"
	"time"
)

type Statistics struct {
	mutex          sync.RWMutex
	EmailsSent     int
	EmailsRejected int
	LastSendTime   time.Time
	DailyCount     map[string]int // key: YYYY-MM-DD
	Recipients     map[string]int // Track recipient frequencies
	Templates      map[string]int // Track template usage
}

func (s *Statistics) RecordSend(input email.SendEmailInput) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	today := time.Now().Format("2006-01-02")
	s.EmailsSent++
	s.DailyCount[today]++
	s.LastSendTime = time.Now()

	// Record recipient statistics
	for _, recipient := range input.Destination.ToAddresses {
		s.Recipients[recipient]++
	}
}

func (s *Statistics) GetStats() map[string]interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return map[string]interface{}{
		"totalEmailsSent":     s.EmailsSent,
		"totalEmailsRejected": s.EmailsRejected,
		"dailyCount":          s.DailyCount,
	}
}
