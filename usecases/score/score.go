package score

import (
	"desafio_letras2021/usecases/comparator"
	"strings"
)

// Measure how much letters from the firstWord variable matches the letter in the
// same position but in the secondWord.
func measureWordScore(firstWord string, secondWord string) (int, error) {
	var scoreWord int
	var iterationWord, minorWord string

	if len(firstWord) > len(secondWord) {
		iterationWord = secondWord
		minorWord = firstWord
	} else if len(firstWord) < len(secondWord) {
		iterationWord = firstWord
		minorWord = secondWord
	}

	// fmt.Println(firstWord, secondWord)
	for i, character := range iterationWord {
		if character == rune(minorWord[i]) {
			scoreWord += 1
		}
	}
	// fmt.Println(scoreWord)

	return scoreWord, nil
}

// Measure the total socre for the musicTitle in comparison with the searchTitle.
func MeasureTotalScore(searchTitle string, musicTitle string) (int, error) {
	var score int
	var featInMusiTitle, featInSearchTitle bool

	// Separates the words from the strings using the space delimiter.
	sliceMusicTitle := strings.Fields(musicTitle)
	sliceSearchTitle := strings.Fields(searchTitle)

	// Measure the score combining each word from the two strings.
	for _, word := range sliceMusicTitle {

		// Save information when the music title in the database has a "feat" word.
		if word == "feat" {
			featInMusiTitle = true
		}

		for _, titleWord := range sliceSearchTitle {

			// fmt.Println(strings.Count(titleWord, ""), titleWord, word)
			result, _ := comparator.Insensitive(titleWord, word)
			if result == true {
				// Update score based on the value when the word from the search
				// matches the word from the music. Furthermore, we increment
				// the score based on the number of caracterers in the titleWord.
				score += 10 + strings.Count(titleWord, "") - 1

				// Save information when the searched title has a "feat" word.
				if titleWord == "feat" {
					featInSearchTitle = true
				}

			} else {
				// Measure the word score when the titleWord and word are not
				// the same.
				scoreWord, err := measureWordScore(titleWord, word)
				if err != nil {
					return score, err
				}
				score += scoreWord
			}
		}
	}

	// When the searched title does not have "feat" word and the music title in
	// the database has, then, we need to subtract 5 from the search score.
	if featInMusiTitle == true && featInSearchTitle == false {
		score -= 5
	}

	return score, nil
}
