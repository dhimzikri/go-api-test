package main

import (
	"fmt"
	"go_api_cimol/config"
	"go_api_cimol/routes"
	"net/http"

	_ "go_api_cimol/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title CIMOL API with Golang
// @version 1.0
// @description CIMOL API with Golang.
// @description Golang memiliki keunggulan dalam skalabilitas horizontal dan pemeliharaan kode yang lebih sederhana,
// @description Golang juga memiliki keamanan bawaan dan memudahkan pencegahan beberapa jenis serangan, serta
// @description kinerja yang signifikan dalam kasus pengolahan beban kerja tinggi dan pembuatan API yang membutuhkan respons cepat.
// termsOfService http://swagger.io/terms/

// @contact.name API CNAF MOBILE
// @contact.url https://www.cnaf.co.id/
// @contact.email dimas.zikri@cnaf.co.id

// license.name Apache 2.0
// license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api

// @schemes http https

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("https://localhost:9090/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.ToolsRoutes(router)

	// Path to the SSL certificate and private key files
	certFile := "localhost.crt"
	keyFile := "localhost.key"

	// Start the HTTPS server
	err := http.ListenAndServeTLS(":9090", certFile, keyFile, r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	// log.Println("Server running on port 9090")
	// http.ListenAndServe(":9090", r)
}
