package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/shong91/cryptocurrency/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	templateDir string = "explorer/templates/"
)
var templates *template.Template

func home(rw http.ResponseWriter, r *http.Request) {
	// request 는 pointer[r *http.Request] 를 사용하여 원본을 받는다 - 다양한 형식, 빅데이터 등 대량 데이터도 request 로 들어올 수 있기 때문
	
	// how to try-catch in Go? 
	// Must(): 에러가 있는지 체크해 주는 helper function
	data := homeData{"home", blockchain.GetBlockchain().AllBlocks()}
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
func Start(port int) {
	handler := http.NewServeMux()

	// load all templates
	// golang does NOT support [**/] .. depth 가 있는 경우 이렇게 나눠서 ParseGlob 처리해주어야 함
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml")) // update variable
	
	// handle URL 
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	// Run server; server-side rendering in Go (which is SUPER EASY!) 
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))

}