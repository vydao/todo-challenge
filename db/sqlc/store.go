package db

import (
	"database/sql"
)

type Store interface {
	Querier
	// MakeCompetition(ctx context.Context, challengeID int64, challengerID int64, rivalID int64) error
}

// SQLStore provides all functions to execute SQL queries
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// func (store *SQLStore) MakeCompetition(ctx context.Context, challengeID int64, challengerID int64, rivalID int64) error {
// 	tx, err := store.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()
// 	qtx := store.Queries.WithTx(tx)
// 	comp, err := qtx.CreateCompetition(ctx, CreateCompetitionParams{
// 		ChallengeID:  challengeID,
// 		ChallengerID: challengerID,
// 		RivalID:      rivalID,
// 		Status:       string(api.Upcoming),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	todos, err := qtx.GetTodosByChallenge(ctx, challengeID)
// 	if err != nil || err == sql.ErrNoRows {
// 		return err
// 	}
// 	// Bulk import competition todos
// 	stmt, err := tx.Prepare(pq.CopyIn("competition_todos", "competition_id", "todo_id"))
// 	if err != nil {
// 		return err
// 	}
// 	for _, todo := range todos {
// 		_, err = stmt.Exec(comp.ID, todo.ID)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	_, err = stmt.Exec()
// 	if err != nil {
// 		return err
// 	}
// 	err = stmt.Close()
// 	if err != nil {
// 		return err
// 	}
// 	return tx.Commit()
// }
