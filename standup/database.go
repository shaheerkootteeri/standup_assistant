package standup

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Connect to PostgreSQL database
var db *sql.DB

func init() {
	// Replace the connection string with your PostgreSQL connection details
	connStr := "postgres://username:password@host:port/database?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

// GetUpdates retrieves updates from the database
func GetUpdates() ([]StandupUpdate, error) {
	rows, err := db.Query("SELECT * FROM standup_updates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []StandupUpdate
	for rows.Next() {
		var update StandupUpdate
		if err := rows.Scan(&update.ID, &update.UserID, &update.Update, &update.Blockers, &update.CreatedAt); err != nil {
			return nil, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}

// InsertUpdate inserts a new update into the database
func InsertUpdate(update StandupUpdate) error {
	_, err := db.Exec("INSERT INTO standup_updates(user_id, update, blockers, created_at) VALUES ($1, $2, $3, $4)",
		update.UserID, update.Update, update.Blockers, update.CreatedAt)
	return err
}
