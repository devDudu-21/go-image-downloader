# go-image-downloader

## Descrição

Este projeto é um downloader de imagens escrito em Go. Ele lê uma planilha Excel e baixa as imagens listadas, salvando-as em uma pasta de saída definida pelo usuário.

## Requisitos

- Go (versão 1.13+ recomendada)
- Biblioteca: github.com/xuri/excelize/v2

## Uso

1. Compile o projeto:

   ```bash
   go build -o img-downloader
   ```

2. Execute o binário:

   ```bash
   ./img-downloader
   ```

3. Siga as instruções no terminal:

   - Informe o nome da pasta de saída.
   - Informe o nome do arquivo Excel (com extensão) que contém as URLs.

   As planilhas devem conter uma URL por linha. Por exemplo:
   URL 1
   URL 2
   URL 3

## Estrutura do Projeto

- main.go: Lógica principal para leitura do Excel e download das imagens.
- README.md: Documentação do projeto.

## Licença

MIT (ou outra licença de sua preferência)
