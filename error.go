package main

func checkError(err error)  {
	if err != nil {
		panic(err.Error())
	}
}