package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)
	// функция выделения латинских слов/букв из текста
func extractWords(str string) []string {
	regex := regexp.MustCompile(`[a-zA-Z]+`)
	return regex.FindAllString(str, -1)
}

func main() {
	fi, err := os.Open("A.txt")
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		return
	}
	defer fi.Close()

	// получаем размер файла
	stat, err := fi.Stat()
	if err != nil {
		fmt.Println("Unable to get size of file: ", err)
		return
	}

	// читаем файл
	bs := make([]byte, stat.Size())
	_, err = fi.Read(bs)
	if err != nil {
		fmt.Println("Unable to read file: ", err)
		return
	}

	str := string(bs)
	words := extractWords(str)

	// сортируем слова по алфавиту
	sort.Strings(words)

	// склеивание отсортированных слов с добавлением разделителя
	newStr := strings.Join(words, ", ")

	// создаем выходной файл
	fo, err := os.Create("B.txt")
	if err != nil {
		fmt.Println("Unable to create file: ", err)
		return
	}
	defer fo.Close()

	// пишем в файл конечный результат
	fo.WriteString(newStr)
}