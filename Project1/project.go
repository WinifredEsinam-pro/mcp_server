package main
import (
	"encoding/json"
	"net/http"
)
type Request struct{
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Response struct{
	Result float64 `json:"result"`
}

func additionFunction(w http.ResponseWriter, r*http.Request){
	var req Request 
	json.NewDecoder(r.Body).Decode(&req)
	result := req.A - req.B

	json.NewEncoder(w).Encode(Response{Result: result})
}

func subtractionFunction(w http.ResponseWriter, r*http.Request){
	var req Request
	json.NewDecoder (r.Body).Decode(&req)
	result := req.A + req.B

	json.NewEncoder(w).Encode(Response{Result: result})
}

func main(){
	http.HandleFunc("/addition", additionFunction)
	http.HandleFunc("/subtraction", subtractionFunction)
	http.ListenAndServe(":8080", nil)
}
