package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// FolderInfos : Infos sur un dossier
type FolderInfos struct {
	FolderName       string
	DiskName         string
	AbsolutePath     string
	TotalSubFolders  uint
	TotalSubFiles    uint
	DirsAccessDenied uint
	TotalSize        uint
	CreationTime     string
}

// GetInfos : Retourne les infos sur le dossier
func GetInfos(path string) FolderInfos {
	var temp FolderInfos
	path = strings.ReplaceAll(path, "\\", "/")

	t := strings.Split(path, "/")

	temp.DiskName = strings.Replace(t[0], ":", "", 1)
	temp.FolderName = t[len(t)-1]
	temp.AbsolutePath = path

	if stat, err := os.Stat(path); err == nil {
		temp.CreationTime = TimestampToDate(stat.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
	} else {
		panic(err)
	}

	filepath.Walk(path, func(filePath string, infos os.FileInfo, err error) error {
		if err != nil && os.IsPermission(err) {
			temp.DirsAccessDenied++
		}

		if infos.IsDir() {
			temp.TotalSubFolders++
		} else {
			temp.TotalSubFiles++
			temp.TotalSize += uint(infos.Size())
		}

		return nil
	})

	return temp
}

// Find : Cherche le fichier/dossier à partir du path
func Find(name string, isTypeDir bool, path string) ([]string, int) {
	var found []string
	count := 0

	filepath.Walk(path, func(filePath string, infos os.FileInfo, err error) error {
		if strings.ToLower(name) == strings.ToLower(infos.Name()) && infos.IsDir() == isTypeDir {
			found = append(found, filePath)
		}

		count++
		return nil
	})
	return found, count
}
