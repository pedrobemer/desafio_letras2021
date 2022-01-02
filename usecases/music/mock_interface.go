package music

import (
	"desafio_letras2021/entity"
	"errors"
)

type MockDatabase struct {
}

func NewMusicsPostgres() *MockDatabase {
	return &MockDatabase{}
}

func (m *MockDatabase) GetMusics(searchType string, limit int, offset int) (
	[]entity.Musics, error) {

	if searchType == "DATABASE_ERROR" {
		return nil, errors.New("Some error in the musics table")
	}

	return []entity.Musics{
		{
			Id:   "0",
			Name: "Que Tiro Foi Esse",
		},
		{
			Id:   "1",
			Name: "Havana",
		},
		{
			Id:   "2",
			Name: "Pirata e Tesouro",
		},
		{
			Id:   "3",
			Name: "Aleluia (part. Michely Manuely)",
		},
		{
			Id:   "4",
			Name: "Me Leva Pra Casa",
		},
		{
			Id:   "5",
			Name: "Corpo Sensual (part. Mateus Carrilho)",
		},
		{
			Id:   "6",
			Name: "Safadômetro",
		},
		{
			Id:   "7",
			Name: "Havana feat Young Thug",
		},
		{
			Id:   "8",
			Name: "Trem Bala",
		},
		{
			Id:   "9",
			Name: "Pesadão (part. Marcelo Falcão)",
		},
		{
			Id:   "10",
			Name: "Tem Café (part. Gaab)",
		},
		{
			Id:   "11",
			Name: "Perfect Duet feat Beyoncé",
		},
		{
			Id:   "12",
			Name: "Vidinha de Balada",
		},
		{
			Id:   "13",
			Name: "Vai Malandra (part. MC Zaac, Maejor, Tropkillaz e DJ Yuri Martins)",
		},
		{
			Id:   "14",
			Name: "Amor da Sua Cama",
		},
		{
			Id:   "15",
			Name: "Imaturo",
		},
		{
			Id:   "16",
			Name: "Amar, Amei",
		},
		{
			Id:   "17",
			Name: "Bom Rapaz (part. Jorge e Mateus)",
		},
		{
			Id:   "18",
			Name: "Fica (part. Matheus e Kauan)",
		},
		{
			Id:   "19",
			Name: "Minha Namorada",
		},
	}, nil
}
