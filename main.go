package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    caminho := "/Applications"
    
    arquivos, err := ioutil.ReadDir(caminho)
    if err != nil {
        log.Fatal(err)
    }
    
    permitidos := map[string]bool{
        "Safari.app": true,
        "Pages.app":  true,
    }
    
    for _, arquivo := range arquivos {
        if arquivo.IsDir() && !permitidos[arquivo.Name()] {
            fmt.Printf("Encontrado: %s\n", arquivo.Name())
        }
    }
}
