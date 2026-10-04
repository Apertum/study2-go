package main

import (
	"fmt"
	"strings"
)

//Прокси
//Прокси (заместитель) — это структурный паттерн проектирования,
//который позволяет управлять доступом к объекту через замещающий объект.
//Прокси даёт возможность совершать дополнительные действия до или после передачи вызова оригиналу.
//Есть ещё один термин, которым обозначают объект-заместитель, — middleware.
//В случае веб-сервера middleware могут выступать в роли обработчиков запросов,
//обеспечивающих аутентификацию, логирование и т. д.
//Паттерн Прокси используется, когда:
//нужна ленивая инициализация (виртуальный прокси) — если есть объект, создание которого требует больших затрат, прокси может создавать этот объект только тогда, когда он действительно понадобится;
//нужно разграничить доступ к объекту в зависимости от прав вызывающего объекта (защищающий прокси);
//нужно логирование запросов (логирующий прокси) — объект-прокси может хранить всю историю обращений;
//нужно кешировать объекты, результаты запросов и управлять их жизненным циклом («умные» ссылки);
//нужно настроить взаимодействие с удалённым сервером (удалённый прокси) — использование прокси позволяет снизить издержки при передаче данных.

// Server — интерфейс веб-сервера.
type Server interface {
	handleRequest(string, string) (int, string)
}

// WebServer — сервер по умолчанию.
type WebServer struct {
}

// Proxy — прокси-сервер.
type Proxy struct {
	WebServer *WebServer
}

func NewProxyServer(webServer *WebServer) *Proxy {
	return &Proxy{
		WebServer: webServer,
	}
}

func (p *Proxy) handleRequest(url, method string) (int, string) {
	// запрещает доступ к /api/admin
	if strings.HasPrefix(url, "/api/admin") {
		return 403, "Forbidden"
	}
	return p.WebServer.handleRequest(url, method)
}

func (webServer *WebServer) handleRequest(url, method string) (int, string) {
	if !strings.HasPrefix(url, "/api/") {
		return 400, "Bad Request"
	}
	return 200, fmt.Sprintf("%s Request: %s", method, url)
}

func main() {
	proxyServer := NewProxyServer(&WebServer{})
	for _, v := range []string{"/api/info", "/api/admin", "/api/version", "/admin"} {
		httpCode, body := proxyServer.handleRequest(v, "GET")
		fmt.Println(v, httpCode, body)
	}
}
