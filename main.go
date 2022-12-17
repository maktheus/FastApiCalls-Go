package main

import (
	"flag"
	"FastApiCaller/events/src"
)

var duration int
var url := []string;


func init() {

	// Definimos a flag "duration" para o tempo de execução do script em segundos.
	flag.IntVar(&duration, "d", 60, "tempo de execução do script em segundos")

// Definimos a flag "urls" para as URLs da API.
flag.StringArrayVar(&urls, "u", []string{"https://api.github.com/users/brunocarvalho"}, "URLs da API")

	// Analisamos as flags.
	flag.Parse()
}

func main() {
	//chamar a funcao que faz a chamada da api
	apiCaller.ApiCallerMainFunction(duration, urls)
}
