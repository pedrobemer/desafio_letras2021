package main

import (
	"bufio"
	"desafio_letras2021/usecases/music"

	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchMusicTitle := scanner.Text()

	// Search the best matches for the music title
	musicScoreInfo := music.SearchMusic(searchMusicTitle, 10)

	fmt.Println("# Resposta:")
	for _, musicInfo := range musicScoreInfo {
		fmt.Println("# ", musicInfo.Score, " pontos, "+musicInfo.Name)
	}
	fmt.Println("#")
	fmt.Println("# --------------------------------------------------")
}
