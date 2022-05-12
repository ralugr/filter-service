package repository

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/config"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

const rejMsg string = `
	CREATE TABLE IF NOT EXISTS RejectedMessages (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL UNIQUE,
    	body TEXT NOT NULL,
		state TEXT NOT NULL,
		reason TEXT NOT NULL
  	);`

const queuedMsg string = `
	CREATE TABLE IF NOT EXISTS QueuedMessages (
  		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uid TEXT NOT NULL UNIQUE,
  		body TEXT NOT NULL,
		state TEXT NOT NULL,
		reason TEXT NOT NULL
	);`

const bannedWords string = `
	CREATE TABLE IF NOT EXISTS BannedWords (
  		id INTEGER NOT NULL PRIMARY KEY,
		words TEXT NOT NULL
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

	if err = sqlite.createTable(rejMsg); err != nil {
		return nil, err
	}
	if err = sqlite.createTable(queuedMsg); err != nil {
		return nil, err
	}
	if err = sqlite.createTable(bannedWords); err != nil {
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
		return errors.New("failed to insert message")
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

func (sqlite *SQLiteDB) UpdateBannedWords(bw *model.BannedWords) error {
	_, err := sqlite.db.Exec("REPLACE INTO BannedWords (id, words) VALUES(?,?)", bw.Id, strings.Join(bw.Words, ","))

	if err != nil {
		logger.Warning.Printf("Could not replace words in db %v", err)
		return err
	}
	return nil
}

func (sqlite *SQLiteDB) GetBannedWords() (*model.BannedWords, error) {
	var bw *model.BannedWords
	rows, err := sqlite.db.Query("SELECT * FROM BannedWords LIMIT 1")

	var id int
	var words string

	rows.Next()
	err = rows.Scan(&id, &words)
	rows.Close()

	if words != "" {
		bw = model.NewBannedWords(strings.Split(words, ","))
	}

	if err != nil {
		logger.Warning.Printf("Could not get banned words from db %v", err)
		return bw, err
	}

	logger.Info.Printf("Banned words retrieved from db %v", bw)
	return bw, nil
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
	_, err := sqlite.db.Exec("INSERT INTO "+table+" (uid,body,state,reason) VALUES(?,?,?,?)", message.UID, message.Body, message.State, message.Reason)

	if err != nil {
		logger.Info.Printf("Could not insert message %v into %v table: Reason %v ", message, table, err.Error())
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
