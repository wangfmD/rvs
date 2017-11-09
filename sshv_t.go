package main

import (
	"fmt"
	"github.com/wangfmD/rvs/sshv"
)

func main() {
	fmt.Println("++++++")
	v, err := sshv.GetVersionMaps("root", "xungejiaoyu2011", "10.1.41.61")
	fmt.Println(v)
	fmt.Println(err)
	fmt.Println("--end--")

}
