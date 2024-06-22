package main

import (
	"encoding/json"
	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser
	jsonData := `{"user":{"name": "John Doe", "age": 23}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	userJSON := value.GetObject("user").String()

	var user User
	if err = json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}
	println(user.Name, user.Age)
	//fmt.Println("User name: %s\n", user.Get("name"))
	//fmt.Println("User age: %s\n", user.Get("age"))
}
