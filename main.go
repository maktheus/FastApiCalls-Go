package main

import (
	"flag"
	"FastApiCaller/events/src"
)

var duration int
var urls []string = make([]string, 0)
var file string


func init() {

	// Definimos a flag "duration" para o tempo de execução do script em segundos.
	flag.IntVar(&duration, "d", 60, "tempo de execução do script em segundos")

	// Definimos a flag "urls" para as URLs da API.
	flag.StringArrayVar(&urls, "u", []string{"https://api.github.com/users/brunocarvalho"}, "URLs da API")

	// Definimos a flag "file" para o tipo de arquivo que será gerado.
	flag.StringVar(&file, "f", "json", "tipo de arquivo que será gerado")

	// Analisamos as flags.
	flag.Parse()
}

func main() {
	//chamar a funcao que faz a chamada da api
	apiCaller.ApiCallerMainFunction(duration, urls)
}
