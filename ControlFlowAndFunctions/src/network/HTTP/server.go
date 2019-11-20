package HTTP

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func StartCaseConverterServer(port string) {
	fmt.Println("HTTP case converter server started on port:", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi upper case of request is :", strings.ToUpper(r.URL.Path[1:]),"Requseted on port:",os.Getenv("PORT"))
}