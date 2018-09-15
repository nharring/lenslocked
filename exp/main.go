package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	foomap := map[string]string{"key1": "value1", "keyfoo": "valuebar"}
	data := struct {
		Name        string
		FavoriteDay string
		UnixTime    int64
		Tags        []string
		KVP         map[string]string
		PrintMap    bool
	}{"John Smith", "Friday", time.Now().Unix(), []string{"foo", "bar"}, foomap, false}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
