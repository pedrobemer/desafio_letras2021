package main

import (
	"bufio"
	"context"
	"desafio_letras2021/database/postgresql"
	"desafio_letras2021/usecases/music"
	"desafio_letras2021/usecases/utils"

	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {

	// Read environment variables for the database connection
	filenamePath := "./"
	filename := "database"
	DB_USER := utils.ViperReadEnvVariable(filenamePath, filename, "DB_USER")
	DB_PASSWORD := utils.ViperReadEnvVariable(filenamePath, filename, "DB_PASSWORD")
	DB_NAME := utils.ViperReadEnvVariable(filenamePath, filename, "DB_NAME")
	DB_PORT := utils.ViperReadEnvVariable(filenamePath, filename, "DB_PORT")
	DB_HOST := utils.ViperReadEnvVariable(filenamePath, filename, "DB_HOST")

	// Create string for the database connection
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	// Connect the PostgreSQL database
	DBpool, err := pgx.Connect(context.Background(), dbinfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer DBpool.Close(context.Background())

	// Create instance to read the musics table from the database
	musicsRepo := postgresql.NewMusicsPostgres(DBpool)

	// Create instance for the usecases
	musicsApp := music.NewApplication(musicsRepo)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchMusicTitle := scanner.Text()

	// Search the best matches for the music title
	musicsScoreInfo := musicsApp.SearchMusic(searchMusicTitle, 10)

	fmt.Println("# Resposta:")
	for _, musicInfo := range musicsScoreInfo {
		fmt.Println("# ", musicInfo.Score, " pontos, "+musicInfo.Name)
	}
	fmt.Println("#")
	fmt.Println("# --------------------------------------------------")
}
