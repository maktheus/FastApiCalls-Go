package apiCaller

import (
	"fmt"
	"time"
	"sync"
	"encoding/json"
)


func ApiCallerMainFunction(duration int, urls []string, file string, name string)  {
	
	var Result []ApiResult

	if duration <= 0 {
		fmt.Println("O tempo de execução do script deve ser maior que 0")
		return
	}
	if len(urls) == 0 {
		fmt.Println("É necessário informar pelo menos uma URL")
		return
	}

	//print o tempo de duracao do script em  00:00:00

	var tempoTotal =   time.Duration(duration) * time.Second 

	fmt.Println("Tempo de execução do script: ", tempoTotal)
	
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	// Criamos um grupo de espera para garantir que todas as threads
	// sejam concluídas antes de terminarmos o programa.
	var wg sync.WaitGroup

	// Vamos percorrer cada URL e criar uma thread para fazer a chamada.

	for time.Now().Before(endTime){
		// Vamos percorrer cada URL e criar uma thread para fazer a chamada.
		for _, url := range urls {
			// Adicionamos uma unidade ao contador de threads no grupo de espera.
			wg.Add(1)

			// Iniciamos uma thread e passamos a URL e o grupo de espera como argumentos.
			go func(url string, wg *sync.WaitGroup) {
				// Quando a thread for finalizada, removemos uma unidade do contador no grupo de espera.
				defer wg.Done()
				
				response, err := MakeRequest(url)
				if err != nil {
					fmt.Println(err)
				}
				var apiResult ApiResult
				err = json.Unmarshal(response, &apiResult)
				if err != nil {
					fmt.Println(err)
				}
				Result = append(Result, apiResult)


			}(url, &wg)
		}
	time.Sleep(1 * time.Second)
	}
	
	// Aqui, nós esperamos até que todas as threads tenham sido concluídas.
	wg.Wait()
	
	switch file {
	case "json":
		//chama a funcao que gera o arquivo json
		FileJsonCreator(Result,name)
	case "csv":
		//chama a funcao que gera o arquivo csv
		FileCsvCreator(Result,name)
	case "xlsx":
		//chama a funcao que gera o arquivo xml
		FileXlsxCreator(Result,name)
	case "txt":
		//chama a funcao que gera o arquivo txt
		FileTxtCreator(Result,name)
	default:
		fmt.Println("Tipo de arquivo inválido")
	}
}
