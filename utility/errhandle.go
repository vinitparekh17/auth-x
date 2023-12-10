package utility

func ErrorHandler(e error) {
	if e != nil {
		panic(e)
	}
}
