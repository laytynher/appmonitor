package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    // Diretório de aplicativos
    appsDir := "/Applications"
    
    files, err := ioutil.ReadDir(appsDir)
    if err != nil {
        log.Fatal(err)
    }
    
    // Lista de aplicativos permitidos
    allowedApps := map[string]bool{
        "Safari.app": true,
        "Pages.app":  true,
    }
    
    // Verifica se os aplicativos instalados são permitidos
    for _, f := range files {
        if f.IsDir() && !allowedApps[f.Name()] {
            fmt.Printf("Aplicativo não permitido detectado: %s\n", f.Name())
        }
    }
}
