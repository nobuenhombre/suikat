package yank

import "time"

type Timing struct {
	PrepareRequest  time.Duration `json:"prepare_request_ns"`  // Подготовка запроса
	Connect         time.Duration `json:"connect_ns"`          // Установка соединения (DNS + TCP + TLS)
	SendRequest     time.Duration `json:"send_request_ns"`     // Отправка заголовков + тела
	TimeToFirstByte time.Duration `json:"ttfb_ns"`             // Ожидание первого байта ответа
	DownloadContent time.Duration `json:"download_content_ns"` // Загрузка тела ответа
	ParseResponse   time.Duration `json:"parse_response_ns"`   // Парсинг ответа
	Total           time.Duration `json:"total_ns"`            // Общее время
}
