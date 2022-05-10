package repository

//import (
//	"fmt"
//	"github.com/ralugr/filter-service/internal/common"
//	"github.com/ralugr/filter-service/internal/database"
//	"github.com/ralugr/filter-service/internal/logger"
//	"github.com/ralugr/filter-service/internal/model"
//	"testing"
//)
//
//func TestCreateTable(t *testing.T) {
//	sqldb, err := InitDB()
//
//	if err != nil {
//		t.Error("Failed to init database")
//	}
//
//	input := []struct {
//		name     string
//		expected string
//	}{
//		{"", "QueuedMessages"},
//		{"", "RejectedMessages"},
//	}
//
//	for tc, tt := range input {
//		err = sqldb.db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tt.expected).Scan(&tt.name)
//
//		fmt.Println(" Starting test case ", tc)
//
//		if tt.name != tt.expected {
//			t.Errorf("\nExpected %v \nActual   %v", tt.expected, tt.name)
//		}
//	}
//}
//
//func TestInsertRejectedMessage(t *testing.T) {
//	sqldb, err := InitDB()
//
//	if err != nil {
//		t.Error("Failed to init database")
//	}
//
//	input := []struct {
//		m        model.Message
//		reason   string
//		expected bool
//	}{
//		{common.MockMessage1, "Bla", true},
//		{common.MockMessage2, "Lorem", true},
//	}
//
//	for tc, tt := range input {
//		fmt.Println(" Starting test case ", tc)
//		sqldb.InsertRejectedMessage(&tt.m, tt.reason)
//
//		//if tt.name != tt.expected {
//		//	t.Errorf("\nExpected %v \nActual   %v", tt.expected, tt.name)
//		//}
//	}
//}
//
//func TestInsertQueuedMessage(t *testing.T) {
//
//}
//
//func TestGetMessages(t *testing.T) {
//
//}
//
//func InitDB() (*database.SQLiteDB, error) {
//	db, err := database.NewSQLite("test_storage.db")
//
//	if err != nil {
//		logger.Warning.Println("Could not initialize db ", err)
//		return nil, err
//	}
//
//	return db, nil
//}
