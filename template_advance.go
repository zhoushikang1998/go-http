package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		tmpl, err := template.New("test").Parse(`
{{$name := "Alice"}}
{{$age := 18}}
{{$round := true}}
Name: {{$name}}
Age: {{$age}}
Round: {{$round}}`)
		if err != nil {
			fmt.Fprintf(w, "parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4001", nil))
}
