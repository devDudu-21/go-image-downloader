# go-image-downloader

## Description

This project is an image downloader written in Go. It reads an Excel spreadsheet and downloads the listed images, saving them into an output folder specified by the user.

## Requirements

- Go (version 1.13+ is recommended)
- Library: github.com/xuri/excelize/v2

## Usage

1. Compile the project:

   For Linux and macOS:

   ```bash
   go build -o img-downloader
   ```

   For Windows:

   ```bash
   go build -o img-downloader.exe
   ```

2. Run the binary:

   ```bash
   ./img-downloader
   ```

3. Follow the terminal instructions:

   - Enter the output folder name.
   - Enter the Excel file name (with extension) that contains the URLs.

   The spreadsheet should have one URL per line. For example:
   URL 1  
   URL 2  
   URL 3

## Project Structure

- main.go: Main logic for reading the Excel file and downloading images.
- README.md: Project documentation.

## License

MIT (or any other license of your choice)
