package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {
	data := []Movie{
		{Title: "钢铁侠3", Year: 2015, Color: false, Actors: []string{"托尼"}},
		{Title: "复仇者联盟4", Year: 2019, Color: true, Actors: []string{"黑寡妇", "美国队长"}},
	}

	// go结构体解析成json
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}

	jsonData2, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData2))
	}

	// json转为go结构体
	var titles []struct {
		Title string
		// 没有定义其他字段,将被舍弃
	}
	if err := json.Unmarshal(jsonData, &titles); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(titles)
	}
}
