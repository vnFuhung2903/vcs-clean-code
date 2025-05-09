package proxy

type App struct{}

func (a *App) Handle(url string) uint {
	if url == "/status" {
		return 200
	}

	if url == "/create" {
		return 201
	}
	return 404
}
