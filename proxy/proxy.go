package proxy

import "math/rand"

type Proxy struct {
	app *App
}

func NewProxy() *Proxy {
	return &Proxy{
		&App{},
	}
}

func (p *Proxy) SendMessage() string {
	rand := rand.Int()
	if rand%2 == 0 {
		return "Message from proxy itself"
	}
	return p.app.SendMessage()
}
