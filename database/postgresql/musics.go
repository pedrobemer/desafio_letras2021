package postgresql

import (
	"context"
	"desafio_letras2021/entity"

	"github.com/georgysavva/scany/pgxscan"
	_ "github.com/lib/pq"
)

type MusicsPostgres struct {
	dbpool PgxIface
}

func NewMusicsPostgres(db PgxIface) *MusicsPostgres {
	return &MusicsPostgres{
		dbpool: db,
	}
}

func (m *MusicsPostgres) GetMusics(searchType string, limit int, offset int) (
	[]entity.Musics, error) {
	var musics []entity.Musics

	query := `
	SELECT
		id, created_at, updated_at, title
	FROM musics;
	`

	err := pgxscan.Select(context.Background(), m.dbpool, &musics, query)
	if err != nil {
		return nil, err
	}

	return musics, nil
}
