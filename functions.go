package main

// CheckAndPanic : Panic si err != nil
func CheckAndPanic(err error) {
	if err != nil {
		panic("Une erreur est survenue: " + err.Error())
	}
}
