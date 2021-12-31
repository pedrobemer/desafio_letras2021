package main

import (
	"bufio"
	"desafio_letras2021/usecases/music"

	"fmt"
	"os"

	"golang.org/x/text/unicode/norm"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchMusicTitle := scanner.Text()

	fmt.Println(norm.NFD)

	// Escreva a solução aqui
	musicScoreInfo := music.SearchMusic(searchMusicTitle)
	// response := searchMusicTitle

	fmt.Println("# Resposta:")
	for _, musicInfo := range musicScoreInfo {
		fmt.Println("# ", musicInfo.Score, " pontos, "+musicInfo.Name)
	}
	fmt.Println("#")
	fmt.Println("# --------------------------------------------------")
}
