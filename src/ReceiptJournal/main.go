package main

import (
	"ReceiptJournal/controller"
	"ReceiptJournal/middleware"
	"ReceiptJournal/model"
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), &middleware.TimeoutMiddleware{new(middleware.GzipMiddleWare)}))

}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"

	fm := template.FuncMap{}
	fm["displayComma"] = func(price float64) string {
		return middleware.RenderFloat("", price)
	}

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Funcs(fm).Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			println(err)
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func connectToDatabase() *sql.DB {
	/*dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		model.Hostadd,
		model.Dbname,
		model.Dbuser,
		model.Dbpass)

	db, err := sql.Open("cloudsqlpostgres", dsn)*/

	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		"ReceiptJournal",
		"kelvin",
		"8456594")

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	var now time.Time
	fmt.Println(db.QueryRow("SELECT NOW()").Scan(&now))
	fmt.Println(now)
	return db
}
