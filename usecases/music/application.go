package music

import (
	"desafio_letras2021/entity"
	"desafio_letras2021/usecases/score"
	"desafio_letras2021/usecases/treatment"
	"sort"
)

type Application struct {
	repo Repository
}

//NewApplication create new use case
func NewApplication(r Repository) *Application {
	return &Application{
		repo: r,
	}
}

type sortMusicSlice []entity.MusicInfo

// Application that Search the best musics title in our database based on the
// "title" parameter. In this case the return is a slice with a length equal to
// "bestReturns"""
func (a *Application) SearchMusic(title string, bestReturns int) sortMusicSlice {

	// Get musics from the database
	musicsInfo, err := a.repo.GetMusics("ALL", 0, 0)
	if err != nil {
		return sortMusicSlice{}
	}

	// Pre-allocate memory for the slice responsible to store the best returns
	// based on the measured score.
	musicNameSort := make(sortMusicSlice, bestReturns+1)
	for i := 0; i < len(musicNameSort); i++ {
		musicNameSort[i].Score = -100
	}

	// Treat music title for search
	treatedTitle, err := treatment.TreatString(title)
	if err != nil {
		return sortMusicSlice{}
	}

	for _, musicName := range musicsInfo {

		// Treat the music title name from the database to measure the score.
		treatedMusicName, err := treatment.TreatString(musicName.Name)
		if err != nil {
			return []entity.MusicInfo{}
		}

		// Measure the score for the treated title based on the treated music
		// from the database.
		score, err := score.MeasureTotalScore(treatedTitle, treatedMusicName)
		if err != nil {
			return []entity.MusicInfo{}
		}

		// Save the music in the last position to see if it is in the field of
		// the best returns.
		musicNameSort[bestReturns].Name = musicName.Name
		musicNameSort[bestReturns].Score = score

		// fmt.Println("pre-sort:", musicNameSort)

		// Sort the slice based on the descending order of the score and when the
		// score is the same using the alphabetical order.
		sort.Sort(musicNameSort)

		// fmt.Println("pos-sort:", musicNameSort)

		// fmt.Println("Score:", score)

	}

	// return the best returns based on the bestReturns threshold.
	return musicNameSort[0:bestReturns]
}

func (a *Application) SearchMusic2(title string, bestReturns int) sortMusicSlice {

	musicsInfo := []string{"Que Tiro Foi Esse", "Deixe-me Ir",
		"Sobre N??s (Poesia Ac??stica #2)", "Apelido Carinhoso",
		"T?? Com Moral No C??u", "Lugar Secreto", "J??", "Perfect", "Fica Tranquilo",
		"Capricorniana (Poesia Ac??stica #3)", "Amor da Sua Cama", "Nessas Horas",
		"Downtown (part. J Balvin)", "Voc?? Vai Entender", "Aquieta Minh'alma",
		"Havana", "Havana feat Young Thug",
		"Vai Malandra (part. MC Zaac, Maejor, Tropkillaz e DJ Yuri Martins)",
		"Prioridade", "Trevo (Tu) (part. Tiago Iorc)",
		"Machika (part. Anitta y Jeon)", "Trem Bala", "Mo??a do Espelho",
		"Safad??metro", "Eu Cuido de Ti", "Too Good At Goodbyes",
		"Duro Igual Concreto", "Aquela Pessoa", "Rap Lord (part. Jonas Bento)",
		"Contrato", "IDGAF", "De Quem ?? a Culpa?", "N??o Troco", "Quase",
		"Deus ?? Deus", "Anti-Amor", "Eu Era",
		"Cerveja de Garrafa (Fuma??a Que Eu Fa??o)", "N??o Deixo N??o",
		"Rockstar feat 21 Savage", "New Rules", "Photograph", "Eu Juro",
		"Ningu??m Explica Deus (part. Gabriela Rocha)",
		"Lindo ??s", "Bengala e Croch??", "Pirata e Tesouro", "A Libertina",
		"Pesad??o (part. Marcelo Falc??o)", "Aleluia (part. Michely Manuely)",
		"Eu Cuido de Ti", "Oi", "C??u Azul", "Never Be The Same",
		"My Life Is Going On", "Imaturo", "Gucci Gang", "Cuidado", "K.O.",
		"??chame La Culpa", "??chame La Culpa feat Luis Fonsi",
		"Tem Caf?? (part. MC Hariel)", "Raridade", "Te Vi Na Rua Ontem",
		"Dona Maria (feat Jorge)", "Fica (part. Matheus e Kauan)",
		"9 Meses (Ora????o do Beb??)", "Muleque de Vila", "A Vit??ria Chegou",
		"Ar Condicionado No 15", "Vida Loka Tamb??m Ama", "Pegada Que Desgrama",
		"Transplante (part. Bruno & Marrone)", "Na Conta da Loucura",
		"Tem Caf?? (part. Gaab)", "Apelido Carinhoso", "Perfect Duet",
		"Perfect Duet feat Beyonc??", "Cora????o de A??o", "Minha Morada", "Amar, Amei",
		"Regime Fechado", "O Escudo", "Minha Namorada",
		"Quero Conhecer Jesus (O Meu Amado ?? o Mais Belo)", "Me Leva Pra Casa",
		"Como ?? Que Faz? (part. Rob Nunes)", "The Scientist", "Bella Ciao",
		"O Que Tiver Que Ser Vai Ser", "Corpo Sensual (part. Mateus Carrilho)",
		"Cor de Marte", "Bom Rapaz (part. Jorge e Mateus)", "Vidinha de Balada",
		"N??o Era Voc??", "Em Teus Bra??os", "De Tr??s Pra Frente", "All Of Me",
		"Believer", "A M??sica Mais Triste do Ano", "Rabiola",
		"Para??so (part. Pabllo Vittar)", "Vem Pra Minha Vida",
	}

	// Pre-allocate memory for the slice responsible to store the best returns
	// based on the measured score.
	musicNameSort := make(sortMusicSlice, bestReturns+1)
	for i := 0; i < len(musicNameSort); i++ {
		musicNameSort[i].Score = -100
	}

	// Treat music title for search
	treatedTitle, err := treatment.TreatString(title)
	if err != nil {
		return sortMusicSlice{}
	}

	for _, musicName := range musicsInfo {

		// Treat the music title name from the database to measure the score.
		treatedMusicName, err := treatment.TreatString(musicName)
		if err != nil {
			return []entity.MusicInfo{}
		}

		// Measure the score for the treated title based on the treated music
		// from the database.
		score, err := score.MeasureTotalScore(treatedTitle, treatedMusicName)
		if err != nil {
			return []entity.MusicInfo{}
		}

		// Save the music in the last position to see if it is in the field of
		// the best returns.
		musicNameSort[bestReturns].Name = musicName
		musicNameSort[bestReturns].Score = score

		// fmt.Println("pre-sort:", musicNameSort)

		// Sort the slice based on the descending order of the score and when the
		// score is the same using the alphabetical order.
		sort.Sort(musicNameSort)

		// fmt.Println("pos-sort:", musicNameSort)

		// fmt.Println("Score:", score)

	}

	// return the best returns based on the bestReturns threshold.
	return musicNameSort[0:bestReturns]
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
