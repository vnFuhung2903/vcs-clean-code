package proxy

type App struct{}

func (a *App) SendMessage() string {
	return "Message from app"
}
