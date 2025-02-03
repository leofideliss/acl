package main

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/leofideliss/acl/internal/controller"
)

var (
    porta int
    urlBase string
)

func init(){
    porta = 8888
    urlBase = fmt.Sprintf("http://localhost:%d",porta)
}

func main(){
    // groups
    http.HandleFunc("/api/group/" , AddGroup)
}
