package main

import (
	"fmt"

	"github.com/vnFuhung2903/vcs-clean-code/abstractFactory"
	"github.com/vnFuhung2903/vcs-clean-code/di"
	"github.com/vnFuhung2903/vcs-clean-code/proxy"
	"github.com/vnFuhung2903/vcs-clean-code/singleton"
)

func main() {
	fmt.Println("Count using singleton pattern:", singleton.Count())
	fmt.Println("Recount using singleton pattern:", singleton.Count())

	proxy := proxy.NewProxy()
	fmt.Println("Get message using proxy pattern:", proxy.SendMessage())

	factory := abstractFactory.NewFactory("test")
	fmt.Println("Get name using abstract factory pattern:", factory.GetName())

	di.RunDI()
}
