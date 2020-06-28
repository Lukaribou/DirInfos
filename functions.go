package main

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
