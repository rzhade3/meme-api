package handler

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "path/filepath"
  "os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  dir, err := os.Getwd()
  if err != nil {
    fmt.Println("Error getting filepath")
    return
  }
  scriptPath := filepath.Join(dir, "db", "script.txt")
  fmt.Println(scriptPath)
  data, err := ioutil.ReadFile(scriptPath)
  if err != nil {
    // fmt.Println("File reading error", err)
    return
  }
  fmt.Fprintf(w, string(data))
}
