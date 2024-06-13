package main

import "fmt"

type vertex struct {
	x, y int
}

// third attempt
// optimized solution, broke into two functions to implement backtracking through recursion, and cleaned code structure / removed unnecessary maps
func dfs(p vertex, board [][]byte, word string, index int, visited map[vertex]bool) bool {
	if index == len(word) {
		return true
	}

	if p.x < 0 || p.x >= len(board) || p.y < 0 || p.y >= len(board[0]) || visited[p] || board[p.x][p.y] != word[index] {
		return false
	}

	visited[p] = true
	defer func() { visited[p] = false }() // backtrack

	// search
	directions := []vertex{{p.x + 1, p.y}, {p.x - 1, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}}
	for _, dir := range directions {
		if dfs(dir, board, word, index+1, visited) {
			return true
		}
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	foundWords := make([]string, 0)
	visited := make(map[vertex]bool)

	for _, word := range words {
		wordFound := false
		for i := 0; i < len(board) && !wordFound; i++ {
			for j := 0; j < len(board[i]) && !wordFound; j++ {
				if dfs(vertex{i, j}, board, word, 0, visited) {
					foundWords = append(foundWords, word)
					wordFound = true
				}
			}
		}
	}
	return foundWords
}

func main () {
	// leetcode input
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	// expected output: "oath" and "eat"
	fmt.Println("found words:", findWords(board, words))
}