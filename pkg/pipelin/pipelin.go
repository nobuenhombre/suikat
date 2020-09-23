package pipelin

import (
	"sync"
)

// func Pattern
type Work func(in, out chan interface{})

type WorkPipeline []Work

func SyncWork(in, out chan interface{}, myWork Work, wg *sync.WaitGroup) {
	myWork(in, out)
	close(out)
	wg.Done()
}

func (wp *WorkPipeline) Run() {
	var wg sync.WaitGroup

	// Выходной поток каждой функции (Work) является входным потоком следующей
	// in -> A ->[out,in]-> B ->[out,in]-> C ->[out,in]-> D -> out
	// для 4-х методов  будет 2 (входящий, исходящий) и 3 промежуточных канала
	// in -> A ->[out,in]-> B ->[out,in]-> C ->[out,in]-> D ->[out,in]-> E -> out
	// для 5-х методов  будет 2 (входящий, исходящий) и 4 промежуточных канала
	// Тут создаем необходимые каналы
	count := len(*wp) + 1
	channels := make([]chan interface{}, 0, count)

	for i := 0; i < count; i++ {
		newChannel := make(chan interface{}, 1)
		channels = append(channels, newChannel)
	}

	for index, work := range *wp {
		wg.Add(1)

		go SyncWork(channels[index], channels[index+1], work, &wg)
	}

	wg.Wait()
}

//func WorkerExample(in, out chan interface{}) {
//	var result resultType
//
//	for dataRaw := range in {
//		dataTyped, ok := dataRaw.(inType)
//		if !ok {
//			// some error processing
//		}
//		// some result processing
//	}
//
//	out <- result
//}

//func FirstWorkerExample(in, out chan interface{}) {
//	inputData := []int{0, 1, 2, 3, 4, 5, 6, 7}
//
//	for _, num := range inputData {
//		out <- num
//	}
//}
