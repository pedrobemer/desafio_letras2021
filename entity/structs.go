package entity

import "time"

type MusicInfo struct {
	Name  string
	Score int
}

type Musics struct {
	Id        string    `db:"id" json:"id,omitempty"`
	CreatedAt time.Time `db:"created_at" json:",omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:",omitempty"`
	Name      string    `db:"title" json:",omitempty"`
}
