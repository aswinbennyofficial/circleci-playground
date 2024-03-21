package main

import(
	"net/http"
	"fmt"
	"os"
	"log"
)

func main(){
	http.HandleFunc("/", handler)

	log.Println("Server starting at 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	// Get the secret key from the environment
	secret:=os.Getenv("SECRET_KEY")
	fmt.Fprintf(w, "Hello World, The secret env is %s",secret)

}