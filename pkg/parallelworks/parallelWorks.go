// Package parallelworks provides start parallel calls some function with collecting results
package parallelworks

import (
	"sync"
)

const (
	WorkersCount                    = 64
	WorkersCountToChanLenMultiplier = 2
)

type JobData struct {
	ID   int
	Data interface{}
}

type Worker func(data JobData) JobData

func (data JobData) WriteToChan(a chan<- JobData) {
	for {
		if len(a) < cap(a) {
			break
		}
	}
	a <- data
}

func (worker Worker) goWorker(jobs <-chan JobData, result chan<- JobData, wg *sync.WaitGroup) {
	defer wg.Done()

	for jobData := range jobs {
		wr := worker(jobData)
		wr.WriteToChan(result)
	}
}

func (worker Worker) sendData(dataList []JobData, jobsChan chan<- JobData) {
	defer close(jobsChan)

	for _, t := range dataList {
		t.WriteToChan(jobsChan)
	}
}

func (worker Worker) readResults(resultsChan <-chan JobData, outData *map[int]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for res := range resultsChan {
		(*outData)[res.ID] = res.Data
	}
}

func (worker Worker) Run(dataList []JobData, workersCount int) map[int]interface{} {
	channelsLength := len(dataList)
	if len(dataList) > WorkersCountToChanLenMultiplier*workersCount {
		channelsLength = WorkersCountToChanLenMultiplier * workersCount
	}

	outData := make(map[int]interface{}, len(dataList))
	jobsChan := make(chan JobData, channelsLength)
	resultsChan := make(chan JobData, channelsLength)

	go worker.sendData(dataList, jobsChan)

	var waitWorkers, waitResults sync.WaitGroup

	waitWorkers.Add(workersCount)

	for w := 1; w <= workersCount; w++ {
		go worker.goWorker(jobsChan, resultsChan, &waitWorkers)
	}

	go func() {
		waitWorkers.Wait()
		close(resultsChan)
	}()

	waitResults.Add(1)

	go worker.readResults(resultsChan, &outData, &waitResults)

	waitResults.Wait()

	return outData
}
