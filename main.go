package main

import (
	"fmt"
	"os/exec"
)

func main() {
	infos := GetInfos(getCmdUserPosition())

	fmt.Printf(
		`
Infos sur le dossier: "%s"

Disque: %s
Sous dossiers: %d (Accès refusés: %d)
Fichiers: %d
Taille: %do
`,
		infos.FolderName, infos.DiskName, infos.TotalSubFolders, infos.TotalSubFiles, infos.DirsAccessDenied, infos.TotalSize)
}

func getCmdUserPosition() string {
	out, err := exec.Command("cmd", "/C", "echo", "%CD%").CombinedOutput()
	CheckAndPanic(err)
	return RemoveNonVisibleChars(string(out))
}
