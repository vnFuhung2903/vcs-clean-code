package proxy

import "net/mail"

type Proxy struct {
	app *App
}

func NewProxy() *Proxy {
	return &Proxy{
		&App{},
	}
}

func (p *Proxy) Handle(url string, email string) uint {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return 403
	}
	return p.app.Handle(url)
}
