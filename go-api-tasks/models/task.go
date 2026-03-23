package models

// Variaveis de ambiente

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"` // JSON deve ser "done"
	CreatedAt string `json:"created_at"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done INTEGER NOT NULL DEFAULT 0,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`
)
