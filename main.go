package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
type Job struct {
	comicNumber int
}

var client = &http.Client{Timeout: 15 * time.Second}

func fetch(n int) (*Result, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}
	defer resp.Body.Close()
	var data Result
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}
	return &data, nil
}

// This is the channel where we will be sending jobs i.e(Id of the comic to be fetched)
// Workers will take comic id from this job channel and send it's result to the result channel
var jobs = make(chan Job, 100)
var results = make(chan Result, 100)
var resultCollection []Result

// This function will send the comic id's of all the comics to be fetched to the job channel and these jobs will be processed by the worker
func allocateJobs(totalcomic int) {
	for i := 0; i < totalcomic; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)

}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result, err := fetch(job.comicNumber)
		if err != nil {
			log.Printf("Error fetching job %d: %v\n", job.comicNumber, err)
		}
		results <- *result
	}
}

// This function will create workes that will handle the job of sending api requests
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func collectResults(done chan bool) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("I have received the result no:%d", result.Num)
			resultCollection = append(resultCollection, result)
		}
	}
	done <- true
}
func main() {
	done := make(chan bool)
	go allocateJobs(1000)
	go collectResults(done)
	go createWorkerPool(100)
	fmt.Print("Now I am waiting for all worker to finish their task")
	<-done
	fmt.Print("Everything Done from the workers")
}
