package main

import (
	"fmt"
	"path/filepath"
	"sync"
)

func main() {

	var option int

	for option != 3 {

		fmt.Println("1 - Indexar diretório")
		fmt.Println("2 - Buscar termo")
		fmt.Println("3 - Sair")

		fmt.Scanln(&option)

		if option == 1 {

			var dir string

			fmt.Print("Diretório: ")
			fmt.Scanln(&dir)

			jobs := make(chan string, 100)

			var wg sync.WaitGroup

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					worker(jobs)
				}()
			}

			indexDirectory(dir, jobs)

			close(jobs)

			wg.Wait()

			fmt.Println("Indexação finalizada")

		} else if option == 2 {

			var term string

			fmt.Print("Digite o termo: ")
			fmt.Scanln(&term)

			results := SearchTerm(term)

			if len(results) == 0 {

				fmt.Println("Termo não encontrado. Adicionando ao arquivo...")
				InserirTermo(term)

			} else {

				fmt.Println("Encontrado em:")

				for _, r := range results {
					fmt.Printf("  Diretório : %s\n", filepath.Dir(r))
					fmt.Printf("  Arquivo   : %s\n\n", filepath.Base(r))
				}
			}
		}
	}
}
