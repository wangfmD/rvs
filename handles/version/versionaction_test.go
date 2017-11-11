package version

import (
	"fmt"
	"testing"
)

// func TestSqlQuerybyid(t *testing.T) {
// 	res := sqlQueryVersionById("3d04b9cd-b339-4f9c-8c3b-2dbd4deed48a")
// 	fmt.Println(res)
// }

// TestsqlQueryMediaVersionById ...
func TestSqlQueryMediaVersionById(t *testing.T) {
	res := sqlQueryMediaVersionById("3d04b9cd-b339-4f9c-8c3b-2dbd4deed48a")
	fmt.Println(res)

}
