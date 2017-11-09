package setting

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {

	res := GetTagVersions()
	for key, value := range res {
		fmt.Println(key, value)
	}
	// t.Fatal()
	fmt.Println("fmt")
}

func TestGetTags(t *testing.T) {
	p := GetTagVersions()
	addrs := make([]string, 0, 0)
	for k, _ := range p {
		addrs = append(addrs, k)
	}
	fmt.Println(addrs)

}
