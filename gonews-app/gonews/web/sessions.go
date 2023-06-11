package web

import (
	"database/sql"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

func NewSessionsManager(datasouirceName string) (*scs.SessionManager, error) {

	db, err := sql.Open("postgres", datasouirceName)
	if err != nil {
		return nil, err
	}

	sessions := scs.New()
	sessions.Store = postgresstore.New(db)

	return sessions, nil
}
