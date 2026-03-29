package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SearchTerm busca um termo no índice e retorna os arquivos onde foi encontrado
func SearchTerm(term string) []string {
	file, err := os.Open(indexFile)
	if err != nil {
		fmt.Println("Índice não encontrado. Execute a indexação primeiro.")
		return nil
	}
	defer file.Close()

	var results []string
	term = strings.ToLower(term)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 && parts[0] == term {
			// Evita duplicatas
			encontrado := false
			for _, r := range results {
				if r == parts[1] {
					encontrado = true
					break
				}
			}
			if !encontrado {
				results = append(results, parts[1])
			}
		}
	}

	return results
}

// InserirTermo adiciona o termo no arquivo docs/test.txt e no índice
func InserirTermo(term string) {
	term = strings.ToLower(term)

	// Grava no docs/test.txt
	txtFile, err := os.OpenFile("docs/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir docs/test.txt:", err)
		return
	}
	defer txtFile.Close()
	txtFile.WriteString(term + "\n")

	// Grava no índice
	gravarNoIndice(term, "docs/test.txt")

	fmt.Println("Termo adicionado em docs/test.txt e no índice")
}
