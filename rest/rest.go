package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shong91/cryptocurrency/blockchain"
	"github.com/shong91/cryptocurrency/utils"
)

var port string

type url string

// implement TextMarshaler
// https://pkg.go.dev/encoding#TextMarshaler
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// Add field tag in order to use in JSON
type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
	IgnoreField string `json:"-"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

// example: implement Stringer interface
func (u urlDescription) String() string {
	return "Hello I'm the URL Description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	// sending JSON
	// urlDescription slice
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
			IgnoreField: "ignore",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	// rw.Header().Add("Content-Type", "application/json")

	// Marshalling in order to send by JSON
	// 소프트웨어는 바이트 단위로 데이터를 인식한다.메모리 바이트는 해석하는 틀에 따라 달라지는데, 이러한 변환(논리적 구조를 로우 바이트로 변환)을 '인코딩' 또는 '마샬링Mashaling'이라고 한다.
	// 반대로, byte slice 나 문자열을 논리적 자료 구조로 변환하는 것을 언마샬링Unmashaling 이라 한다.
	// Go에서는 encoding package 에서 마샬링을 담당
	json.NewEncoder(rw).Encode(data)
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s",b)

}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		// decode requestbody
		var addBlockBody addBlockBody // empty struct variable

		// set decoding data in variable - using pointer
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		fmt.Println(addBlockBody)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}

}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // map 형태로 key:value 받아옴
	// Atoi: convert string to int
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		// convert error to string 
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

// create middleware
func jsonContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	// ListenAndServe(port, nil): nil 일 경우 기본 multiplexer 를 사용
	// multiplexer 는 클라이언트가 보낸 요청을 어디로 보낼지 결정하는데, main.go 에서 호출한 두 package의 Start() 가 같은 url 을 호출하고 있음 
	// http.HandleFunc 의 url 이 같기 때문에 multiple registration 오류 발생
	// => use their own multiplexer (http.NewServeMux())

	// use gorillaMux which is more effective than default multiplexer NewServerMux
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleWare)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET") // gorillaMux can use regex
	
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}