package main

import (
	"fmt"

	"github.com/vnFuhung2903/vcs-clean-code/abstractFactory"
	"github.com/vnFuhung2903/vcs-clean-code/proxy"
	"github.com/vnFuhung2903/vcs-clean-code/singleton"
)

func main() {
	fmt.Println("Count using singleton pattern:", singleton.Count())
	fmt.Println("Recount using singleton pattern:", singleton.Count())

	proxy := proxy.NewProxy()
	fmt.Println("Check status using proxy pattern:", proxy.Handle("/create", "user@gmail"))

	mFactory := abstractFactory.NewFactory("mercedes")
	nFactory := abstractFactory.NewFactory("nissan")
	fmt.Println("Check Mercedes logo using abstract factory pattern:", mFactory.Produce().GetLogo())
	fmt.Println("Check Nissan logo using abstract factory pattern:", nFactory.Produce().GetLogo())
}
