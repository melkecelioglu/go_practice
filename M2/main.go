package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// TemplateHandler represents a single template
type TemplateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	addr := flag.String("addr", ":8080", "The address of the application")
	flag.Parse()

	// Create a new ChatRoom
	chatRoom := NewChatRoom()

	http.Handle("/", &TemplateHandler{filename: "chat.html"})
	http.Handle("/room", chatRoom)

	go chatRoom.Run()

	serverAddress := *addr
	log.Printf("Starting web server on %s\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
