package music

import (
	"desafio_letras2021/entity"
	"desafio_letras2021/usecases/score"
	"desafio_letras2021/usecases/treatment"
	"sort"
)

type sortMusicSlice []entity.MusicInfo

func SearchMusic(title string, bestReturns int) sortMusicSlice {

	// Until now, this simulates a search in the database
	// dbMusic := []string{"Havana (feat) Young Thug"}
	dbMusic := []string{"Que Tiro Foi Esse", "Deixe-me Ir",
		"Sobre Nós (Poesia Acústica #2)", "Apelido Carinhoso", "Tô Com Moral No Céu",
		"Lugar Secreto", "Jó", "Perfect", "Fica Tranquilo",
		"Capricorniana (Poesia Acústica #3)", "Amor da Sua Cama", "Nessas Horas",
		"Downtown (part. J Balvin)", "Você Vai Entender", "Aquieta Minh'alma", "Havana",
		"Havana feat Young Thug", "Vai Malandra (part. MC Zaac, Maejor, Tropkillaz e DJ Yuri Martins)",
		"Prioridade", "Trevo (Tu) (part. Tiago Iorc)", "Machika (part. Anitta y Jeon)",
		"Trem Bala", "Moça do Espelho", "Safadômetro", "Eu Cuido de Ti",
		"Too Good At Goodbyes", "Duro Igual Concreto", "Aquela Pessoa",
		"Rap Lord (part. Jonas Bento)", "Contrato", "IDGAF", "De Quem É a Culpa?",
		"Não Troco", "Quase", "Deus É Deus", "Anti-Amor", "Eu Era",
		"Cerveja de Garrafa (Fumaça Que Eu Faço)", "Não Deixo Não", "Rockstar feat 21 Savage",
		"New Rules", "Photograph", "Eu Juro", "Ninguém Explica Deus (part. Gabriela Rocha)",
		"Lindo És", "Bengala e Crochê", "Pirata e Tesouro", "A Libertina",
		"Pesadão (part. Marcelo Falcão)", "Aleluia (part. Michely Manuely)", "Eu Cuido de Ti",
		"Oi", "Céu Azul", "Never Be The Same", "My Life Is Going On", "Imaturo",
		"Gucci Gang", "Cuidado", "K.O.", "Échame La Culpa", "Échame La Culpa feat Luis Fonsi",
		"Tem Café (part. MC Hariel)", "Raridade", "Te Vi Na Rua Ontem",
		"Dona Maria (feat Jorge)", "Fica (part. Matheus e Kauan)", "9 Meses (Oração do Bebê)",
		"Muleque de Vila", "A Vitória Chegou", "Ar Condicionado No 15", "Vida Loka Também Ama",
		"Pegada Que Desgrama", "Transplante (part. Bruno & Marrone)", "Na Conta da Loucura",
		"Tem Café (part. Gaab)", "Apelido Carinhoso", "Perfect Duet", "Perfect Duet feat Beyoncé",
		"Coração de Aço", "Minha Morada", "Amar, Amei", "Regime Fechado", "O Escudo",
		"Minha Namorada", "Quero Conhecer Jesus (O Meu Amado é o Mais Belo)", "Me Leva Pra Casa",
		"Como é Que Faz? (part. Rob Nunes)", "The Scientist", "Bella Ciao",
		"O Que Tiver Que Ser Vai Ser", "Corpo Sensual (part. Mateus Carrilho)", "Cor de Marte",
		"Bom Rapaz (part. Jorge e Mateus)", "Vidinha de Balada", "Não Era Você", "Em Teus Braços", "De Trás Pra Frente",
		"All Of Me", "Believer", "A Música Mais Triste do Ano", "Rabiola",
		"Paraíso (part. Pabllo Vittar)", "Vem Pra Minha Vida",
	}

	musicNameSort := make(sortMusicSlice, bestReturns+1)
	for i := 0; i < len(musicNameSort); i++ {
		musicNameSort[i].Score = -100
	}

	// Treat music title for search
	treatedTitle, err := treatment.TreatString(title)
	if err != nil {
		return sortMusicSlice{}
	}

	for _, musicName := range dbMusic {

		// Treat the music title name from the database to measure the score;
		treatedMusicName, err := treatment.TreatString(musicName)
		if err != nil {
			return []entity.MusicInfo{}
		}

		// Measure the score for the treated title based on the treated music
		// name obtained from the database.
		score, err := score.MeasureTotalScore(treatedTitle, treatedMusicName)
		if err != nil {
			return []entity.MusicInfo{}
		}

		musicNameSort[bestReturns].Name = musicName
		musicNameSort[bestReturns].Score = score

		// fmt.Println("pre-sort:", musicNameSort)

		// Sort the slice based on the descending order of the score and when the
		// score is the same using the alphabetical order.
		sort.Sort(musicNameSort)

		// fmt.Println("pos-sort:", musicNameSort)

		// fmt.Println("Score:", score)

	}

	return musicNameSort[0:10]
}

// Use the optimum sorting algorithm available on Golang. It uses an algorithm that
// has a O(n log(n)) in the worst case scenario. In Golang, this can be done
// using the interface for the sort.Sort() function. This function needs the Len()
// method which basically measures the slice length for the sort algorithm.
func (ms sortMusicSlice) Len() int {
	return len(ms)
}

// This method is part of the ms object. It is necessary for the sort.Interface.
// In our case, as we want a descending sort and when the Score in the "i" and "j"
// has the same valeu we want to sort based on the alphabetical order, then, this
// method will return true only when the score from the "i" is higher than the "j"
// or when the Name from "i" is alphabetically lower than the name from "j" since
// the score are the same for both.
func (ms sortMusicSlice) Less(i, j int) bool {
	return ms[i].Score > ms[j].Score || (ms[i].Score == ms[j].Score &&
		ms[i].Name < ms[j].Name)
}

// This method is part of the ms object. It is necessary for the sort.Interface.
// This Swap function has the purpose to switch the i-th and j-th from the slice
// when the conditions for swap is valid, which happens when the Less method
// returns true.
func (ms sortMusicSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
