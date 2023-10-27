package main

import (
   // "embed"
    "fmt"
  //  "io/fs"
    "log"
    "net/http"
    "encoding/json"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
)
/*
go:embed all:views/build
*/
//var svelteStatic embed.FS

type Element struct {
	position int
	name     string
	weight   float64
	symbol   string
}

func main() {

   /* s, err := fs.Sub(svelteStatic, "views/build")
    if err != nil {
        panic(err)
    }
    

    staticServer := http.FileServer(http.FS(s))
*/
    r := chi.NewRouter()
    
    r.Use(cors.Handler(cors.Options{

    AllowedOrigins:   []string{"https://*", "http://*"},

    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))

    
    data := []Element{
		{1, "Hydrogen", 1.0079, "H"},
		{2, "Helium", 4.0026, "He"},
		{3, "Lithium", 6.941, "Li"},
		{4, "Beryllium", 9.0122, "Be"},
		{5, "Boron", 10.811, "B"},
		{1, "Hydrogen", 1.0079, "H"},
		{2, "Helium", 4.0026, "He"},
		{3, "Lithium", 6.941, "Li"},
		{4, "Beryllium", 9.0122, "Be"},
		{5, "Boron", 10.811, "B"},
	}
    
    r.Get("/post/{no}", func(w http.ResponseWriter, r *http.Request){
            
            w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(map[string]interface{}{

            "data": data,
        })
    })

/*
    r.Handle("/", staticServer)
    r.Handle("/_app/*", staticServer)
    r.Handle("/favicon.png", staticServer)
    r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = "/"
        staticServer.ServeHTTP(w, r)
    })
*/
    fmt.Println("Running on port: 8082")
    log.Fatal(http.ListenAndServe(":8082", r))
}