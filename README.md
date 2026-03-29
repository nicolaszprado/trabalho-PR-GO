# Índice de Busca

Programa em Go que indexa arquivos `.txt` dentro de um diretório e permite buscar palavras, mostrando em qual diretório e arquivo cada termo foi encontrado. Usa goroutines para processar múltiplos arquivos em paralelo.

## Requisitos

- [Go](https://go.dev/dl/) 1.21 ou superior instalado

## Como rodar

```bash
go run .
```

Ou compilar e executar:

```bash
go build -o search-index.exe .
.\search-index.exe
```

## Uso

O programa tem um menu com 3 opções:

**1 - Indexar diretório**  
Digite o caminho de uma pasta (ex: `docs`). O programa vai percorrer todos os arquivos `.txt` encontrados e indexar cada palavra.

**2 - Buscar termo**  
Digite uma palavra. Se ela estiver no índice, mostra o diretório e o arquivo onde foi encontrada. Se não estiver, o programa adiciona o termo no arquivo `docs/test.txt` e no índice.

**3 - Sair**

## Exemplo

```
1 - Indexar diretório
2 - Buscar termo
3 - Sair
Opção: 1
Diretório: docs
Indexação finalizada

Opção: 2
Digite o termo: teruo
Encontrado em:
  Diretório : docs
  Arquivo   : test.txt
```

## Estrutura dos arquivos

```
.
├── main.go       # menu e lógica de threads
├── indexer.go    # percorre os diretórios (DFS com filepath.Walk)
├── worker.go     # lê os arquivos e grava no índice
├── search.go     # busca e inserção de termos
├── index.txt     # índice gerado (criado automaticamente)
└── docs/         # pasta com arquivos .txt de exemplo
```

O índice fica salvo em `index.txt` no formato `palavra|caminho/do/arquivo.txt`, uma entrada por linha, sem duplicatas.
