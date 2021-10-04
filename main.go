package main

import (
	_ "fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wisepythagoras/go-sudoku-gen/sudoku"
)

func getRandomSudoku(c *gin.Context) {
	curr := time.Now().UnixNano() ^ rand.Int63()

	s := &sudoku.Sudoku{Seed: curr}
	s.Init()
	s.Fill()

	puzzle := s.GeneratePuzzle()

	c.JSON(http.StatusOK, []*sudoku.Sudoku{s, puzzle})
}

func main() {
	router := gin.Default()
	rand.Seed(time.Now().UnixNano())

	router.GET("/sudoku/new", getRandomSudoku)

	router.Run("localhost:8080")
}
