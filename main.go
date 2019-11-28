package main

import(
  "fmt"
  "net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w,"<h1>Whoa, Go is neat and here!</h1>")
}

func about_handler(w, http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"<h1> We r in About page</h1>")
}


func main(){
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/about",about_handler)
  http.ListenAndServe(":8000",nil)
  fmt.Println("Server is running on port 8000")
}


