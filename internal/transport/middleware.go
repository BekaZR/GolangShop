package transport

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
	handler http.Handler
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Логирование запроса
	log.Printf("Запрос: %s %s", r.Method, r.URL.Path)

	// Создаем ResponseWriter-обертку для перехвата и логирования ответа
	responseWriter := &ResponseLogger{ResponseWriter: w}

	// Вызываем следующий обработчик в цепочке
	l.handler.ServeHTTP(responseWriter, r)

	// Логирование ответа
	log.Printf("Ответ: %d", responseWriter.statusCode)
}

type ResponseLogger struct {
	http.ResponseWriter
	statusCode int
}

func (r *ResponseLogger) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func Logging(router http.HandlerFunc) http.Handler {
	return &LoggingMiddleware{
		handler: http.HandlerFunc(router),
	}
}
