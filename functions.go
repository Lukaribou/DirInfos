package main

import (
	"bytes"
	"strconv"
	"time"
)

// CheckAndPanic : Panic si err != nil
func CheckAndPanic(err error) {
	if err != nil {
		panic("Une erreur est survenue: " + err.Error())
	}
}

// RemoveNonVisibleChars : Renvoie la string sans les caractères non visibles (garde les espaces)
func RemoveNonVisibleChars(str string) string {
	var temp []byte
	for _, l := range []byte(str) {
		if l > 39 {
			temp = append(temp, l)
		}
	}
	return string(temp)
}

// SpaceEvery3 : Renvoie le nombre avec un espace tous les 3 chiffres
func SpaceEvery3(n uint) string {
	return reverse(insertNth(reverse(strconv.FormatUint(uint64(n), 10)), ' ', 3))
}

func reverse(str string) string { // https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func insertNth(str string, toInsert rune, allTheN int) string { // https://stackoverflow.com/questions/33633168/how-to-insert-a-character-every-x-characters-in-a-string-in-golang
	var buffer bytes.Buffer

	for i, rune := range str {
		buffer.WriteRune(rune)
		if i%allTheN == allTheN-1 && i != len(str)-1 {
			buffer.WriteRune(toInsert)
		}
	}

	return buffer.String()
}

// TimestampToDate : Renvoie la date correspondant au timestamp
func TimestampToDate(nano int64) string {
	return time.Unix(0, nano).Format("02/01/2006 15h04m05s")
}
