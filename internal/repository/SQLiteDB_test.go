package repository

import (
	"fmt"
	"testing"

	"github.com/ralugr/filter-service/internal/common"
	"github.com/ralugr/filter-service/internal/config"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

func TestCreateTable(t *testing.T) {
	sqldb, err := InitDB()

	if err != nil {
		t.Error("Failed to init database")
	}

	input := []struct {
		name     string
		expected string
	}{
		{"", "QueuedMessages"},
		{"", "RejectedMessages"},
	}

	for tc, tt := range input {
		err = sqldb.db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tt.expected).Scan(&tt.name)

		if err != nil {
			t.Errorf("\nGot unexpected query erro %v", err)
		}

		fmt.Println(" Starting test case ", tc)

		if tt.name != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, tt.name)
		}
	}
}

func TestInsertMessage(t *testing.T) {
	sqldb, err := InitDB()

	if err != nil {
		t.Error("Failed to init database")
	}

	input := []struct {
		m        model.Message
		table    string
		expected interface{}
		error    bool
	}{
		{common.RejectedMockMessage1, "RejectedMessages", nil, false},
		{common.RejectedMockMessage2, "RejectedMessages", nil, false},
		{common.QueuedMockMessag1, "QueuedMessages", nil, false},
		{common.QueuedMockMessag2, "QueuedMessages", nil, false},
		{common.QueuedMockMessag2, "QueuedMessages", "UNIQUE constraint failed: QueuedMessages.uid", true},
	}

	for tc, tt := range input {
		fmt.Println(" Starting test case ", tc)
		result := sqldb.insertMessage(&tt.m, tt.table)

		if !tt.error && result != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, result)
		}

		if tt.error && result.Error() != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, result)
		}
	}
}

func TestStore(t *testing.T) {
	sqldb, err := InitDB()

	if err != nil {
		t.Error("Failed to init database")
	}

	input := []struct {
		m        model.Message
		expected interface{}
		error    bool
	}{
		{common.RejectedMockMessage1, nil, false},
		{common.QueuedMockMessag1, nil, false},
		{common.QueuedMockMessag2, nil, false},
		{common.QueuedMockMessag2, "failed to insert message", true},
		{common.MockMessage1, "message state is unexpected", true},
	}

	for tc, tt := range input {
		fmt.Println(" Starting test case ", tc)
		result := sqldb.Store(&tt.m)

		if !tt.error && result != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, result)
		}

		if tt.error && result.Error() != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, result)
		}
	}
}

func TestGetMessages(t *testing.T) {
	sqldb, err := InitDB()

	if err != nil {
		t.Error("Failed to init database")
	}

	input := []struct {
		m        model.Message
		state    string
		expected interface{}
		error    bool
	}{
		{common.RejectedMockMessage1, model.Rejected, 1, false},
		{common.QueuedMockMessag1, model.Queued, 1, false},
		{common.QueuedMockMessag2, model.Queued, 2, false},
	}

	for tc, tt := range input {
		fmt.Println(" Starting test case ", tc)
		sqldb.Store(&tt.m)

		msg, _ := sqldb.GetMessages(tt.state)
		if len(msg) != tt.expected {
			t.Errorf("\nExpected %v \nActual   %v", tt.expected, len(msg))
		}

	}
}

func InitDB() (*SQLiteDB, error) {
	config := config.Config{
		DBName: "../common/test_storage.db",
		Port:   8080,
		Host:   "",
	}

	sqlitedb, err := NewSQLiteDB(&config)

	if err != nil {
		logger.Warning.Println("Could not initialize db ", err)
		return nil, err
	}

	sqlitedb.db.Exec("DELETE FROM  QueuedMessages")
	sqlitedb.db.Exec("DELETE FROM  RejectedMessages")

	return sqlitedb, nil
}
