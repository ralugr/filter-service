package repository

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/config"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

const rejMsg string = `
	CREATE TABLE IF NOT EXISTS RejectedMessages (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uid INTEGER NOT NULL UNIQUE,
    	body TEXT NOT NULL,
		state TEXT NOT NULL,
		reason TEXT NOT NULL
  	);`

const queuedMsg string = `
	CREATE TABLE IF NOT EXISTS QueuedMessages (
  		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uid INTEGER NOT NULL UNIQUE,
  		body TEXT NOT NULL,
		state TEXT NOT NULL,
		reason TEXT NOT NULL
	);`

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(cfg *config.Config) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", cfg.DBName)
	if err != nil {
		logger.Warning.Printf("Could not open %v %v", cfg.DBName, err)
		return nil, err
	}
	sqlite := &SQLiteDB{db}

	err = sqlite.createTable(rejMsg)
	if err != nil {
		logger.Warning.Println("Could not create table %v", err)
		return nil, err
	}

	err = sqlite.createTable(queuedMsg)
	if err != nil {
		logger.Warning.Println("Could not create table %v", err)
		return nil, err
	}

	return sqlite, nil
}

func (sqlite *SQLiteDB) Store(message *model.Message) error {
	var err error
	if message.State == model.Rejected {
		err = sqlite.insertMessage(message, "RejectedMessages")
	} else if message.State == model.Queued {
		err = sqlite.insertMessage(message, "QueuedMessages")
	} else {
		logger.Info.Printf("Message %v does not belong in Rejected or Queued table. Aborting insert.", message)
		return errors.New("message state is unexpected")
	}

	if err != nil {
		logger.Warning.Printf("Storing message %v failed ", err)
	}

	return nil
}
func (sqlite *SQLiteDB) GetMessages(state string) ([]model.Message, error) {
	var msg *sql.Rows
	var err error

	if state == model.Rejected {
		msg, err = sqlite.selectAll("RejectedMessages")
	} else if state == model.Queued {
		msg, err = sqlite.selectAll("QueuedMessages")
	}
	defer msg.Close()

	if err != nil {
		logger.Warning.Println("SELECT failed ", err)
		return nil, err
	}

	return adapter.ConvertRowsToMessages(msg)
}

func (sqlite *SQLiteDB) createTable(query string) error {
	if _, err := sqlite.db.Exec(query); err != nil {
		logger.Warning.Println("Could not create table", err)
		return err
	}
	return nil
}

func (sqlite *SQLiteDB) insertMessage(message *model.Message, table string) error {
	logger.Info.Printf("Starting INSERT query on %v table, message %v", table, message)
	_, err := sqlite.db.Exec("INSERT INTO "+table+" (uid,body,state,reason) VALUES(?,?,?,?);", message.UID, message.Body, message.State, message.Reason)

	if err != nil {
		logger.Info.Printf("Could not insert message %v into %v table ", message, table)
		return err
	} else {
		logger.Info.Printf("Insert successful for message %v into %v table ", message, table)
	}

	return nil
}

func (sqlite *SQLiteDB) selectAll(table string) (*sql.Rows, error) {
	logger.Info.Printf("Starting SELECT query on %v table", table)
	msg, err := sqlite.db.Query("SELECT * FROM " + table)

	if err != nil {
		logger.Warning.Println("SELECT failed ", err)
		return nil, err
	} else {
		logger.Info.Printf("SELECT successful on %v table", table)
	}

	return msg, err
}
