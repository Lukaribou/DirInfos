package main

import (
	"os"
	"path/filepath"
)

// FolderInfos : Infos sur un dossier
type FolderInfos struct {
	// Total des dossiers en dessous
	TotalSubFolders int
	// Total des fichiers en dessous
	TotalSubFiles int
}

// GetInfos : Retourne les infos sur le dossier
func GetInfos(path string) (FolderInfos, error) {
	var temp FolderInfos

	filepath.Walk(path, func(filePath string, infos os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if infos.IsDir() {
			temp.TotalSubFolders++
		} else {
			temp.TotalSubFiles++
		}

		return nil
	})

	return temp, nil
}
