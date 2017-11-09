package main

import (
	_ "github.com/wangfmD/rvs/log"
	loger "log"

	"errors"
)

func main() {
	loger.Println("ddd")
}

// package main

// import (
// 	"fmt"
// 	"github.com/satori/go.uuid"
// 	"github.com/wangfmD/rvs/setting"
// )

// func f1() {
// 	uuid1 := uuid.NewV4()
// 	fmt.Printf(",%T", uuid1.String())
// 	fmt.Printf(",%T", uuid1)
// }

// func main() {
// 	testplatform()
// 	// testmedia()
// }

// func testplatform() {
// 	// var p *setting.serveros
// 	res := setting.GetPlatformAddrs()

// 	// fmt.printf("%t", res[0])
// 	// fmt.println(res[0].addr)
// 	for key, value := range res {
// 		fmt.Println(key, value)
// 	}
// }

// // testmedia ...
// func testmedia() {

// 	res := setting.GetMediaAddrs()
// 	for key, value := range res {
// 		fmt.Println(key, value)
// 		fmt.Printf("%T", value.Addr)
// 	}
// }
