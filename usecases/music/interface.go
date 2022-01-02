package music

import "desafio_letras2021/entity"

type Repository interface {
	GetMusics(searchType string, limit int, offset int) ([]entity.Musics, error)
}
