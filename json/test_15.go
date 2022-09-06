package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {

	movie := Movie{
		Title:  "星际穿越",
		Year:   2014,
		Price:  99,
		Actors: []string{"马修·麦康纳", "安妮·海瑟薇"},
	}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	fmt.Printf("JSON: %s\n", jsonStr)

	var myMovie Movie

	// 反序列化
	json.Unmarshal(jsonStr, &myMovie)

	fmt.Println(myMovie)

}
