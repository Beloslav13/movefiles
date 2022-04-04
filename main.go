package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const dirNewImages = "new_images"

var mainFileWin = [...]string{"main", "main.exe", "main.go"}

// Находит файлы скрипта, скомпилированные файлы программы
func findMainFiles(x string) bool {
	i := sort.Search(len(mainFileWin), func(i int) bool { return x <= mainFileWin[i] })
	if i < len(mainFileWin) && mainFileWin[i] == x {
		return true
	}
	return false
}

// Перемещает файл
func moveFile(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Создаёт директорию для перемещения файлов
func createDir(pwd string) string {
	dirSave := fmt.Sprintf("%s/%s", pwd, dirNewImages)
	os.Mkdir(dirSave, 0750)
	return dirSave
}

// Находит файлы которые будут перемещены
func getFiles(path, pwd string) {
	// Создаём папку в которую переместим файлы
	dirSave := createDir(pwd)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Скипаем папку в которую перемещаются файлы
		if path == dirNewImages {
			return filepath.SkipDir
		}

		if !info.IsDir() && !findMainFiles(info.Name()) {
			fmt.Println(info.Name())
			new := fmt.Sprintf("%s/%s", dirSave, info.Name())
			moveFile(path, new)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

func main() {
	pwd, _ := os.Getwd()
	getFiles(".", pwd)
}