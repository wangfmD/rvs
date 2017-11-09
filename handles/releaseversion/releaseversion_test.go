package releaseversion

import (
	"fmt"
	"testing"
)

// Testre ...
// func TestRe(t *testing.T) {
// 	SqlAddversion()
// }

func TestQuery(t *testing.T) {
	// res := queryTagVers("V4.020")
	res := queryVersByIps("10.1.41.56")
	fmt.Println(res)
}
