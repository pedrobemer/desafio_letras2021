package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSearchMusic2(b *testing.B) {
	s := []string{
		"Havana feat Young Thug",
	}

	mockedDb := NewMusicsPostgres()
	musicApp := NewApplication(mockedDb)

	for i := 0; i < b.N; i++ {
		for _, str := range s {
			musicApp.SearchMusic2(str, 10)
		}
	}

}

func TestSearchMusic(t *testing.T) {
	type test struct {
		title              string
		bestReturns        int
		expectedMusicSlice sortMusicSlice
	}

	tests := []test{
		{
			title:       "Havana",
			bestReturns: 6,
			expectedMusicSlice: sortMusicSlice{
				{
					Score: 16,
					Name:  "Havana",
				},
				{
					Score: 11,
					Name:  "Havana feat Young Thug",
				},
				{
					Score: 7,
					Name:  "Vai Malandra (part. MC Zaac, Maejor, Tropkillaz e DJ Yuri Martins)",
				},
				{
					Score: 5,
					Name:  "Fica (part. Matheus e Kauan)",
				},
				{
					Score: 4,
					Name:  "Me Leva Pra Casa",
				},
				{
					Score: 4,
					Name:  "Pesadão (part. Marcelo Falcão)",
				},
			},
		},
		{
			title:       "!!!!",
			bestReturns: 5,
			expectedMusicSlice: sortMusicSlice{
				{
					Score: 0,
					Name:  "Aleluia (part. Michely Manuely)",
				},
				{
					Score: 0,
					Name:  "Amar, Amei",
				},
				{
					Score: 0,
					Name:  "Amor da Sua Cama",
				},
				{
					Score: 0,
					Name:  "Bom Rapaz (part. Jorge e Mateus)",
				},
				{
					Score: 0,
					Name:  "Corpo Sensual (part. Mateus Carrilho)",
				},
			},
		},
	}

	mockedDb := NewMusicsPostgres()
	musicApp := NewApplication(mockedDb)

	for _, testCase := range tests {
		musicsInfo := musicApp.SearchMusic(testCase.title, testCase.bestReturns)
		assert.Equal(t, testCase.expectedMusicSlice, musicsInfo)
	}
}
