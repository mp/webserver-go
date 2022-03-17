package main

import (
    "log"
    "net/http"
    "os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, Web!\n"))
}

func main() {
    //get the value of the ADDR environment variable
    addr := os.Getenv("ADDR")

    
    if len(addr) == 0 {
        addr = ":80"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/goog", HelloHandler)

    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
