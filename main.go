package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


const port string = ":4000"
type URL string 

// implement TextMarshaler
// https://pkg.go.dev/encoding#TextMarshaler
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// Add field tag in order to use in JSON 
type URLDescription struct {
	URL URL `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
	IgnoreField string `json:"-"`

}

// example: implement Stringer interface
func (u URLDescription) String() string {
	return "Hello I'm the URL Description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	// sending JSON
	// URLDescription slice
	data := []URLDescription{
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add a block",
			Payload: "data:string",
			IgnoreField: "ignore",
		},
	}
	rw.Header().Add("Content-Type", "application/json")

	// Marshalling in order to send by JSON 
	// 소프트웨어는 바이트 단위로 데이터를 인식한다.메모리 바이트는 해석하는 틀에 따라 달라지는데, 이러한 변환(논리적 구조를 로우 바이트로 변환)을 '인코딩' 또는 '마샬링Mashaling'이라고 한다.
	// 반대로, byte slice 나 문자열을 논리적 자료 구조로 변환하는 것을 언마샬링Unmashaling 이라 한다. 
  // Go에서는 encoding package 에서 마샬링을 담당
	json.NewEncoder(rw).Encode(data)
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s",b)

}

func main() {
	// lexplorer.Start()
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
