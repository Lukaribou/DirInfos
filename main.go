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

Sous dossiers: %s (Accès refusés: %s)
Fichiers: %s
Taille: ~%s octets

Date de création: %s
`,
		infos.FolderName,
		infos.AbsolutePath,
		infos.DiskName,
		SpaceEvery3(infos.TotalSubFolders), SpaceEvery3(infos.DirsAccessDenied),
		SpaceEvery3(infos.TotalSubFiles),
		SpaceEvery3(infos.TotalSize),
		infos.CreationTime)
}

func getCmdUserPosition() string {
	out, err := exec.Command("cmd", "/C", "echo", "%CD%").CombinedOutput()
	CheckAndPanic(err)
	return RemoveNonVisibleChars(string(out))
}
