package controller

import "fmt"

func AddGroup(w *http.ResponseWriter, r *http.Request){
    body :=make([]byte, r.ContentLength , r.ContentLength)
    r.Body.Read(body)
    fmt.Println(body)
}
