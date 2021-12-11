package main

import (
	"fmt"
	"time"
)

//自定义时间类型
type MyTime time.Time

// 实现Marshaler接口
func (mytime MyTime) MarshalJSON() ([]byte, error) {
	n := time.Time(mytime).Format("2006-01-02 15:04")
	t := "\"" + n + "\""
	return []byte(t), nil
}

type person struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Birthday MyTime `json:"birthday"`
}

func main() {
	timestr := "2021-10-13"
	res, _ := time.Parse("2006-01-02", timestr)
	fmt.Println(res)
	//amy := person{
	//	Name:     "Amy",
	//	Country:  "China",
	//	City:     "Beijing",
	//	Birthday: MyTime("2021-10-13"),
	//}
	//data, _ := json.MarshalIndent(amy, "", "    ")
	//fmt.Println(string(data))
}
