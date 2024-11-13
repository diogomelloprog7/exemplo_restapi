package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// Comando SQL para criar a tabela de usuários
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`

	// Comando SQL para criar a tabela de eventos
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	// Executa a criação da tabela de usuários
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Não foi possível criar a tabela USERS: " + err.Error())
	}

	// Executa a criação da tabela de eventos
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Não foi possível criar a tabela EVENTS: " + err.Error())
	}
	createRegistrationTable := `

		CREATE TABLE IF NOT EXISTS registros (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id)
	)
	`
	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Nao foi possivel criar a tabela de registros")
	}
}
