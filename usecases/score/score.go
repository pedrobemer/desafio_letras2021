package score

import (
	"desafio_letras2021/usecases/comparator"
	"errors"
	"fmt"
	"strings"
)

// Measure how much letters from the firstWord variable matches the letter in the
// same position but in the secondWord.
// PS: To work properly the secondWord variable must be greater than or equal in
//     terms of length in comparison with firstWord.
func measureWordScore(firstWord string, secondWord string) (int, error) {
	var scoreWord int

	if len(firstWord) > len(secondWord) {
		return 0, errors.New("firstWord is greater than secondWord")
	}

	fmt.Println(firstWord, secondWord)
	for i, character := range firstWord {
		if character == rune(secondWord[i]) {
			scoreWord += 1
		}
	}
	fmt.Println(scoreWord)

	return scoreWord, nil
}

// Measure the total socre for the musicTitle in comparison with the searchTitle.
func MeasureTotalScore(searchTitle string, musicTitle string) (int, error) {
	var score int
	var featExist bool

	// Separates the words from the strings using the space delimiter.
	sliceMusicTitle := strings.Fields(musicTitle)
	sliceSearchTitle := strings.Fields(searchTitle)

	for _, word := range sliceMusicTitle {
		if word == "feat" {
			featExist = true
			score -= 5
		}
	}

	// Measure the score combining each word from the two strings.
	for _, word := range sliceMusicTitle {
		for _, titleWord := range sliceSearchTitle {

			fmt.Println(strings.Count(titleWord, ""), titleWord, word)
			result := comparator.Insensitive(titleWord, word)
			if result == true {
				// Update score based on the value when the word from the search
				// matches the word from the music. Furthermore, we increment
				// the score based on the number of caracterers in the titleWord.
				score += 10 + strings.Count(titleWord, "") - 1

				if titleWord == "feat" && featExist == true {
					score += 5
				}

			} else {
				if len(titleWord) > len(word) {
					scoreWord, err := measureWordScore(word, titleWord)
					if err != nil {
						return score, err
					}
					score += scoreWord
				} else {
					scoreWord, err := measureWordScore(titleWord, word)
					if err != nil {
						return score, err
					}
					score += scoreWord
				}
			}
		}

	}

	return score, nil
}
