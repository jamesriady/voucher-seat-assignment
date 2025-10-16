package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "./data/vouchers.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	voucherRepository := repository.NewVoucherRepository(db)
	if err := voucherRepository.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	voucherService := service.NewVoucherService(voucherRepository)

	validate := validator.New(validator.WithRequiredStructEnabled())
	// Register a custom TagNameFunc to use the JSON tag name.
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" { // Ignore fields with "-" json tag
			return ""
		}
		return name
	})

	voucherHandler := handler.NewVoucherHandler(voucherService, validate)

	// Setup router
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/check", voucherHandler.CheckVoucher).Methods("POST")
	api.HandleFunc("/generate", voucherHandler.GenerateVoucher).Methods("POST")

	httpHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Authorization"}),
	)(router)

	srv := &http.Server{
		Handler:      httpHandler,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on port 8000")
	log.Fatal(srv.ListenAndServe())
}
