package parallelWorks

import (
	"sync"
)

const (
	WorkersCount = 28
)

type JobData struct {
	ID   int
	Data interface{}
}

type Worker func(data JobData) JobData

func (worker Worker) goWorker(jobs <-chan JobData, result chan<- JobData, wg *sync.WaitGroup) {
	defer wg.Done()

	for jobData := range jobs {
		result <- worker(jobData)
	}
}

func (worker Worker) Run(dataList []JobData, workersCount int) map[int]interface{} {
	var wg sync.WaitGroup

	jobsChan := make(chan JobData, len(dataList))
	resultsChan := make(chan JobData, len(dataList))

	wg.Add(workersCount)
	for w := 1; w <= workersCount; w++ {
		go worker.goWorker(jobsChan, resultsChan, &wg)
	}

	for _, t := range dataList {
		jobsChan <- t
	}

	close(jobsChan)
	wg.Wait()
	close(resultsChan)

	outData := make(map[int]interface{}, len(dataList))
	for res := range resultsChan {
		outData[res.ID] = res.Data
	}

	return outData
}
