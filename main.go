package main

import (
	"fmt"
	"os/exec"
)

func main() {
	infos := GetInfos(getCmdUserPosition())

	fmt.Printf(
		`
Infos sur le dossier: %s

Chemin absolu: %s
Disque: %s

Sous dossiers: %d (Accès refusés: %d)
Fichiers: %d
Taille: ~%s octets

Date de création: %s
`,
		infos.FolderName,
		infos.AbsolutePath,
		infos.DiskName,
		infos.TotalSubFolders, infos.DirsAccessDenied,
		infos.TotalSubFiles,
		SpaceEvery3(infos.TotalSize),
		infos.CreationTime)
}

func getCmdUserPosition() string {
	out, err := exec.Command("cmd", "/C", "echo", "%CD%").CombinedOutput()
	CheckAndPanic(err)
	return RemoveNonVisibleChars(string(out))
}
