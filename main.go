package main

import (
	"fmt"

	"github.com/containerd/cgroups"
)

func main() {

	cg, err := cgroups.Load(cgroups.V1, cgroups.StaticPath("system.slice"))
	if err != nil {
		fmt.Println("loading cgroup error:", err)
		return
	}

	metrics, err := cg.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println("err retrieving memory stats:", err)
		return
	}

	rss := metrics.Memory.RSS

	fmt.Printf("memory rss used is:%v", rss)
}
