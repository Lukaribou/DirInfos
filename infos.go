package main

import (
	"os"
	"path/filepath"
	"strings"
)

// FolderInfos : Infos sur un dossier
type FolderInfos struct {
	TotalSubFolders  uint
	TotalSubFiles    int
	FolderName       string
	DiskName         string
	DirsAccessDenied int
	TotalSize        uint
}

// GetInfos : Retourne les infos sur le dossier
func GetInfos(path string) FolderInfos {
	var temp FolderInfos
	path = strings.ReplaceAll(path, "\\", "/")

	t := strings.Split(path, "/")

	temp.DiskName = strings.Replace(t[0], ":", "", 1)
	temp.FolderName = t[len(t)-1]

	filepath.Walk(path, func(filePath string, infos os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				temp.DirsAccessDenied++
			}
		}

		if infos.IsDir() {
			temp.TotalSubFolders++
		} else {
			temp.TotalSubFiles++
		}
		temp.TotalSize += uint(infos.Size())

		return nil
	})

	return temp
}
