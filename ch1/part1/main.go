package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	// flag.stringvar 绑定的顺序与值无关，name的值 与终端输入的顺序有关（后替换前）
	flag.StringVar(&name, "name", "GO语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "GO语言编程之旅", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}
