package main

import (
	"flag"
	"FastApiCaller/events/src"
)

var duration int


func init() {

	// Definimos a flag "duration" para o tempo de execução do script em segundos.
	flag.IntVar(&duration, "d", 60, "tempo de execução do script em segundos")

	// Analisamos as flags.
	flag.Parse()
}

func main() {
	//chamar a funcao que faz a chamada da api
	apiCaller.ApiCallerMainFunction(duration)
}