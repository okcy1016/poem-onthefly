package main

import (
	"math/rand"
	"time"
	"strconv"
	"strings"
	"os"
	"net/http"
	"io"
	"fmt"
)


func GenPoeUrl() string {

	poeSongBaseUrl := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry-zhCN/master/poetry/poet.song."
	songCiBaseUrl := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry/master/ci/ci.song."
	poeTangBaseUrl := "https://raw.githubusercontent.com/chinese-poetry/chinese-poetry-zhCN/master/poetry/poet.tang."
	var urlSlice []string
	

	// Generate random number for choosing from poet or songci
	seed := rand.NewSource(time.Now().UnixNano())
	rand01 := rand.New(seed).Intn(2)

	// Choosing from poet and songci
	if rand01 == 0 {

		// Generate random number for choosing from poet song or poet tang
		rand01 = rand.New(seed).Intn(2)

		if rand01 == 0 {
			
			urlSlice = append(urlSlice, poeSongBaseUrl)

			// Get a random number in {0, 1000, 2000, ..., 254000} for choosing a random poet song
			rand0254000 := rand.New(seed).Intn(255) * 1000

			// Choose a random poet song
			urlSlice = append(urlSlice, strconv.Itoa(rand0254000))
			urlSlice = append(urlSlice, ".json")
		} else {
			urlSlice = append(urlSlice, poeTangBaseUrl)

			// Get a random number in {0, 1000, 2000, ..., 57000} for choosing a random poet tang
			rand057000 := rand.New(seed).Intn(58) * 1000

			// Choosing
			urlSlice = append(urlSlice, strconv.Itoa(rand057000))
			urlSlice = append(urlSlice, ".json")
		}
	} else {
		urlSlice = append(urlSlice, songCiBaseUrl)

		// Get a random number in {0, 1000, 2000, ..., 21000} for choosing a random songci
		rand021000 := rand.New(seed).Intn(22) * 1000

		// Choosing a random songci
		urlSlice = append(urlSlice, strconv.Itoa(rand021000))
		urlSlice = append(urlSlice, ".json")
	}

	url := strings.Join(urlSlice, "")

	return url
}

func FetchPoetry(url string) (string, error) {

	urlSlice := strings.Split(url, "/")
	fileName := urlSlice[len(urlSlice) - 1]
	fileLoc := "json/" + fileName

	// Create json dir if it doesn't exist
	_, err := os.Stat("json")
	if os.IsNotExist(err) {
		os.Mkdir("json", 0755)
	}

	// Check file existence first
	_, err = os.Stat(fileLoc)
	if os.IsNotExist(err) {

		fmt.Println("\n下载中", url, "保存至", fileLoc, "\n")

		output, err := os.Create(fileLoc)
		if err != nil {
			
			fmt.Println("创建文件时出现错误", fileLoc, "-", err, "\n")

			return "", err
		}
		defer output.Close()

		resp, err := http.Get(url)
		if err!= nil {

			fmt.Println("下载文件时出现错误", url, "-", err, "\n")

			// Delete downloaded part file when fails
			os.Remove(fileLoc)

			return "", err
		}
		defer resp.Body.Close()

		fileBytes, err := io.Copy(output, resp.Body)
		if err != nil {
			fmt.Println("写入文件时出现错误", fileLoc, "-", err, "\n")

			// Delete file while failing to write
			os.Remove(fileLoc)

			return "", err
		}

		fmt.Println("下载完成: ", (fileBytes / 1000), "KB\n")
	} else {
		fmt.Println("\n已在本地发现诗词json文件，路径：", fileLoc, "\n")
	}

	return fileLoc, nil
}
