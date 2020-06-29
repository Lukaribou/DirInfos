package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	switch len(os.Args) {
	case 1:
		infosCommand()
	case 4:
		switch os.Args[1] {
		case "--find", "-f":
			findCommand()
		default:
			fmt.Println("")
		}
	default:
		fmt.Println("Nombre d'arguments incorrect !")
	}
}

func infosCommand() {
	startTime := time.Now()
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

%s éléments analysés en %sms
`,
		infos.FolderName,
		infos.AbsolutePath,
		infos.DiskName,
		SpaceEvery3(infos.TotalSubFolders), SpaceEvery3(infos.DirsAccessDenied),
		SpaceEvery3(infos.TotalSubFiles),
		SpaceEvery3(infos.TotalSize),
		infos.CreationTime,
		SpaceEvery3(infos.TotalSubFolders+infos.TotalSubFiles), SpaceEvery3(uint(time.Since(startTime).Milliseconds())))
}

func findCommand() {
	startTime := time.Now()
	found, count := Find(
		os.Args[3],
		func() bool {
			if os.Args[2] == "file" {
				return false
			} else if os.Args[2] == "folder" || os.Args[2] == "dir" {
				return true
			} else {
				panic(fmt.Sprintf("L'argument \"%s\" n'est pas valide. Il doit être soit: file, folder (ou dir)", os.Args[2]))
			}
		}(),
		getCmdUserPosition())
	if len(found) == 0 {
		fmt.Printf("Aucune correspondance trouvée pour \"%s\". %d éléments analysés.",
			strings.ToLower(os.Args[3]),
			count)
	} else {
		fmt.Printf("%d résultat(s) pour la recherche de \"%s\". %d éléments analysés en %dms.\n\n",
			len(found),
			strings.ToLower(os.Args[3]),
			count,
			time.Since(startTime).Milliseconds())
		for _, f := range found {
			fmt.Println("\t", f)
		}
	}
}

func getCmdUserPosition() string {
	out, err := exec.Command("cmd", "/C", "echo", "%CD%").CombinedOutput()
	CheckAndPanic(err)
	return RemoveNonVisibleChars(string(out))
}
