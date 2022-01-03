package postgresql

import (
	"context"
	"desafio_letras2021/entity"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

type test struct {
	searchType        string
	limit             int
	offset            int
	expectedMusicInfo []entity.Musics
	expectedError     error
}

func TestGetMusicsAll(t *testing.T) {

	created_at := time.Now()
	updated_at := time.Now()

	query := regexp.QuoteMeta(`
	SELECT
		id, created_at, updated_at, title
	FROM musics;
	`)

	columns := []string{"id", "updated_at", "created_at", "title"}

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub entity connection", err)
	}
	defer mock.Close(context.Background())

	rows := mock.NewRows(columns)
	mock.ExpectQuery(query).WithArgs().WillReturnRows(
		rows.AddRow("0a52d206-ed8b-11eb-9a03-0242ac130003", updated_at, created_at,
			"Trem Bala").AddRow("bbbaaa56-ed8b-11eb-9a03-0242ac130003",
			updated_at, created_at, "Havana"))

	Musics := MusicsPostgres{dbpool: mock}

	expectedTest := test{
		searchType: "ALL",
		limit:      0,
		offset:     0,
		expectedMusicInfo: []entity.Musics{
			{
				Id:        "0a52d206-ed8b-11eb-9a03-0242ac130003",
				UpdatedAt: updated_at,
				CreatedAt: created_at,
				Name:      "Trem Bala",
			},
			{
				Id:        "bbbaaa56-ed8b-11eb-9a03-0242ac130003",
				UpdatedAt: updated_at,
				CreatedAt: created_at,
				Name:      "Havana",
			},
		},
		expectedError: nil,
	}

	musicsInfo, err := Musics.GetMusics(expectedTest.searchType,
		expectedTest.limit, expectedTest.offset)
	assert.Equal(t, expectedTest.expectedMusicInfo, musicsInfo)
	assert.Equal(t, expectedTest.expectedError, err)
}

func TestGetMusicsPagination(t *testing.T) {

	created_at := time.Now()
	updated_at := time.Now()

	expectedTest := test{
		searchType: "PAGINATION",
		limit:      1,
		offset:     0,
		expectedMusicInfo: []entity.Musics{
			{
				Id:        "0a52d206-ed8b-11eb-9a03-0242ac130003",
				UpdatedAt: updated_at,
				CreatedAt: created_at,
				Name:      "Trem Bala",
			},
		},
		expectedError: nil,
	}

	query := regexp.QuoteMeta(`
	SELECT
		id, created_at, updated_at, title
	FROM musics
	ORDER_BY updated_at DESC
	LIMIT $1
	OFFSET $2;
	`)

	columns := []string{"id", "updated_at", "created_at", "title"}

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub entity connection", err)
	}
	defer mock.Close(context.Background())

	rows := mock.NewRows(columns)
	mock.ExpectQuery(query).WithArgs(expectedTest.limit, expectedTest.offset).
		WillReturnRows(rows.AddRow("0a52d206-ed8b-11eb-9a03-0242ac130003",
			updated_at, created_at, "Trem Bala"))

	Musics := MusicsPostgres{dbpool: mock}

	musicsInfo, err := Musics.GetMusics(expectedTest.searchType,
		expectedTest.limit, expectedTest.offset)
	assert.Equal(t, expectedTest.expectedMusicInfo, musicsInfo)
	assert.Equal(t, expectedTest.expectedError, err)
}

func TestGetMusicsError(t *testing.T) {

	expectedTest := test{
		searchType:        "PAGINATION",
		limit:             3,
		offset:            2000,
		expectedMusicInfo: nil,
		expectedError:     errors.New("scany: query multiple result rows: Some Error in the database"),
	}

	query := regexp.QuoteMeta(`
	SELECT
		id, created_at, updated_at, title
	FROM musics
	ORDER_BY updated_at DESC
	LIMIT $1
	OFFSET $2;
	`)

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub entity connection", err)
	}
	defer mock.Close(context.Background())

	mock.ExpectQuery(query).WithArgs(expectedTest.limit, expectedTest.offset).
		WillReturnError(errors.New("Some Error in the database"))

	Musics := MusicsPostgres{dbpool: mock}

	musicsInfo, err := Musics.GetMusics(expectedTest.searchType,
		expectedTest.limit, expectedTest.offset)
	assert.Equal(t, expectedTest.expectedMusicInfo, musicsInfo)
	assert.Equal(t, expectedTest.expectedError.Error(), err.Error())
}
