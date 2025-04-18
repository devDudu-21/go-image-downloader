package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/xuri/excelize/v2"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Image Downloader from Spreadsheet")
    fmt.Println("-----------------------------------------")
    
    fmt.Print("Enter the name of the output folder: ")
    output, _ := reader.ReadString('\n')
    output = strings.TrimSpace(output)
    
    err := os.MkdirAll(output, 0777)
    if err != nil {
        fmt.Println("Error creating the folder:", err)
        return
    } else {
        fmt.Println("Folder", output, "created or already exists.")
    }

    fmt.Print("Enter the name of the Excel file with extension: ")
    file, _ := reader.ReadString('\n')
    file = strings.TrimSpace(file)

    f, err := excelize.OpenFile(file)
    if err != nil {
        fmt.Println("Error to open the file:", err)
        return
    }
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println("Failed to close the file:", err)
        }
    }()

    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println("Failed to read the spreadsheet:", err)
        return
    }

    var wg sync.WaitGroup

    semaphore := make(chan struct{}, 20)
    var count int32

    for _, row := range rows {
        for _, cell := range row {
            url := strings.TrimSpace(cell)
            if url == "" {
                continue
            }

            wg.Add(1)
            semaphore <- struct{}{}
            go func(url string) {
                defer wg.Done()
                defer func() { <-semaphore }()

                if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
                    url = "http://" + url
                }

                resp, err := http.Get(url)
                if err != nil {
                    fmt.Println("Failed to send request to", url, ":", err)
                    return
                }
                defer resp.Body.Close()

                if resp.StatusCode != http.StatusOK {
                    fmt.Println("Failed to download the image from", url, ": status", resp.Status)
                    return
                }

                ext := path.Ext(url)
                if ext == "" {
                    ext = ".jpg"
                }

                currentCount := atomic.AddInt32(&count, 1) - 1
                imageName := fmt.Sprintf("image_%d%s", currentCount, ext)
                
                imagePath := filepath.Join(output, imageName)
                out, err := os.Create(imagePath)
                if err != nil {
                    fmt.Println("Failed to create the file", imagePath, ":", err)
                    return
                }
                defer out.Close()

                _, err = io.Copy(out, resp.Body)
                if err != nil {
                    fmt.Println("Failed to save the image from", url, ":", err)
                } else {
                    fmt.Println("[✔️]Image successfully saved as:", imagePath)
                }
            }(url)
        }
    }

    wg.Wait()
}
