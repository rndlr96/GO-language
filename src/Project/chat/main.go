package main

import (
  "log"
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
  "flag"
  "../trace"
  "os"
)
type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  t.once.Do(func(){
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })

  t.templ.Execute(w, r)
}

func main(){
  var addr = flag.String("addr", ":8080", "The addr of the application.")
  flag.Parse()
  r := newRoom()
  r.tracer = trace.New(os.Stdout)
  http.Handle("/assets/", http.StripPrefix("/asstes", http.FileServer(http.Dir("/path/to/assets/"))))
  http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
  http.Handle("/login", &templateHandler{filename:"login.html"})
  http.HandleFunc("/auth/", loginHandler)
  http.Handle("/room", r)

  go r.run()

  log.Println("Starting web server on", *addr)
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
