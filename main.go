package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shong91/cryptocurrency/blockchain"
)


const (
	port 		string = ":4000"
	templateDir string = "templates/"
)
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// request 는 pointer 를 사용하여 원본을 받는다 - 다양한 형식, 빅데이터 등 대량 데이터도 request 로 들어올 수 있기 때문
	// Fprint: writer 에 출력
	// fmt.Fprint(rw, "Hello from home!")

	// solution 1) there is NO try-catch is Go. We need to do manually
	// tmpl, err := template.ParseFiles("templates/home.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// solution 2) Must() 로 감싸줌 -> 오류가 있다면 err 를 출력함
	// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"home", blockchain.GetBlockchain().AllBlocks()}
	// tmpl.Execute(rw, data)
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		// get data from 'name' prop
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)

	}

}

func main() {
	// chain := blockchain.GetBlockchain()
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Third Block")
	// chain.AddBlock("Fourth Block")
	// for _, block := range chain.AllBlocks() {
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %s\n", block.Hash)
	// 	fmt.Printf("PrevHash: %s\n", block.PrevHash)
	// }

	// golang does NOT support **/
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml")) // update variable
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
