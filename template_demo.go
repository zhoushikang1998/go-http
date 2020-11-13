package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并解析模板内容
		//tmpl, err := template.New("test").Parse("Hello World!")
		//tmpl, err := template.New("test").Parse("The value is: {{.}}")
		tmpl, err := template.New("test").Parse(`
	{{/* 打印参数的值 */}}
	Inventory
			SKU: {{.SKU}}
			Name: {{.Name}}
			UnitPrice: {{.UnitPrice}}
			Quantity: {{.Quantity}}
			Subtotal: {{.Subtotal}}
		`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		/*// 获取 URL 参数的值
		val := r.URL.Query().Get("val")*/
		// 根据 URL 查询参数的值创建 Inventory 实例
		inventory := &Inventory{
			SKU:  r.URL.Query().Get("sku"),
			Name: r.URL.Query().Get("name"),
		}
		inventory.UnitPrice, _ = strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		inventory.Quantity, _ = strconv.ParseInt(r.URL.Query().Get("quantity"), 10, 64)
		// 调用模板对象的渲染方法
		err = tmpl.Execute(w, inventory)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Query().Get("val")))
	})*/

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

// 根据单价和数量计算出总价值
func (i *Inventory) Subtotal() float64 {
	return i.UnitPrice * float64(i.Quantity)
}
