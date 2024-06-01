package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 100, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")
}

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func worker(id int, wg *sync.WaitGroup, ch chan<- Result, client *http.Client, url string, requests int) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			ch <- Result{StatusCode: 0, Duration: duration}
			continue
		}
		ch <- Result{StatusCode: resp.StatusCode, Duration: duration}
		resp.Body.Close()
	}
}

func main() {
	flag.Parse()
	if url == "" {
		fmt.Println("A URL do serviço deve ser fornecida.")
		flag.Usage()
		return
	}

	client := &http.Client{}
	results := make(chan Result, requests)
	var wg sync.WaitGroup

	start := time.Now()

	// Dividir as requisições entre os workers
	requestsPerWorker := requests / concurrency
	extraRequests := requests % concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(i, &wg, results, client, url, requestsPerWorker)
	}

	// Distribuir requisições extras, se houver
	if extraRequests > 0 {
		wg.Add(1)
		go worker(concurrency, &wg, results, client, url, extraRequests)
	}

	// Esperar que todas as goroutines terminem
	wg.Wait()
	close(results)

	totalDuration := time.Since(start)
	statusCount := make(map[int]int)
	var totalRequests int

	for result := range results {
		totalRequests++
		statusCount[result.StatusCode]++
	}

	fmt.Printf("Tempo total gasto: %s\n", totalDuration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", statusCount[200])

	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCount {
		if code != 200 {
			fmt.Printf("HTTP %d: %d\n", code, count)
		}
	}
}
