package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

const indexFile = "index.txt"

var mu sync.Mutex

func processFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			gravarNoIndice(word, filePath)
		}
	}
}

func gravarNoIndice(term string, filePath string) {
	mu.Lock()
	defer mu.Unlock()

	entrada := term + "|" + filePath
	existing, err := os.ReadFile(indexFile)
	if err == nil {
		for _, line := range strings.Split(string(existing), "\n") {
			if strings.TrimSpace(line) == entrada {
				return
			}
		}
	}

	indexF, err := os.OpenFile(indexFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao gravar no índice:", err)
		return
	}
	defer indexF.Close()

	indexF.WriteString(entrada + "\n")
}

func worker(jobs chan string) {
	for filePath := range jobs {
		processFile(filePath)
	}
}
