package repeater

import (
	"time"
)

// При взаимодействии с некоторыми внешними апи - например банк Точка
// возникает потребность повторить запрос, потому что на том конце
// еще не закончилось действие
// репитер повторяет вызов воркера до тех пор пока
// не удовлетворится условие чекера или когда будет достигнут лимит повторов
// повторы происходят с некоторым таймаутом между ними
//
// Example
// =======
//func CreateSomeWorker() repeater.Worker {
//	return func(data interface{}) repeater.WorkerResult {
//		in := data.(YourType)
//
//		// Here Some Action
//
//		result := repeater.WorkerResult{
//			OutData: out,
//			Err: err,
//		}
//
//		return result
//	}
//}
//
//func CreateSomeChecker() repeater.Checker {
//	return func(wr WorkerResult) (bool, error) {
//		// Here Some Action with wr
//
//		return result, err
//	}
//}

type WorkerResult struct {
	OutData interface{}
	Err     error
}

// Выполняемая функция
type Worker func(data interface{}) WorkerResult

// Проверка условия завершения
type Checker func(wr WorkerResult) (bool, error)

type Config struct {
	count      int64
	LimitCount int64
	Timeout    time.Duration
}

type LimitCountExceedError struct {
}

func (e *LimitCountExceedError) Error() string {
	return "Limit Count Exceed"
}

func (worker Worker) Run(inData interface{}, checker Checker, config Config) (interface{}, error) {
	config.count++

	wr := worker(inData)

	canReturn, err := checker(wr)
	if err != nil {
		return nil, err
	}

	if canReturn {
		return wr.OutData, nil
	}

	if config.count >= config.LimitCount {
		return nil, &LimitCountExceedError{}
	}

	time.Sleep(config.Timeout)

	return worker.Run(inData, checker, config)
}
