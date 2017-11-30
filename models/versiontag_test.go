package models

import (
	// "encoding/json"
	"fmt"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	s := `{"agent":"V1.01.003.R.001","filesrv":"V2.01.005.R.001","ftp":"2.2.2.1","interact":"V2.01.002.R.001","interactbusiness":"V2.01.001.R.001","jycenter":"V2.01.005.R.001","jyresource":"V2.01.007.R.001","mbs":"0.6.13","middlecas":"V2.02.005.R.001","middlecenter":"V2.01.011.R.002","middlecenterfile":"V2.00.012.R.002","middlecenterres":"V2.01.011.R.001","middleclient":"V2.01.012.R.002","middledatabase":"V1.10.80","middledriver":"V2.01.003.R.001","middleresource":"V2.01.008.R.001","middletherepair":"V2.01.002.R.001","middleware-mcu":"V3.02.001.R.001","mysql":"5.6.19.1","nginx":"1.9.13.2","openfire":"4.0.2.1","redis":"3.0.0"}`

// 	// m := make(map[string]string)
// 	v := VersionTag{}
// 	err := json.Unmarshal([]byte(s), &v)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(v)
// 	}
// 	v.Versionid = "3306"
// 	n, e := v.Add()
// 	if e != nil {
// 		fmt.Println(e)
// 	} else {
// 		fmt.Println("insert suc: ", n)
// 	}

// }

//func TestGetAll(t *testing.T)  {
// 	v := VersionTag{Versionid: "22"}
// 	fmt.Println(v.GetAll())
// }

func TestDe(t *testing.T) {
	v := VersionTag{}
	fmt.Println(v.DeleteOne("2"))
}
