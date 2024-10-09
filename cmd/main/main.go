package main

import (
	//"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/arepala-uml/go-bookstore/pkg/config"
	"github.com/arepala-uml/go-bookstore/pkg/models"
	"github.com/arepala-uml/go-bookstore/pkg/routes"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	// Get the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Build the full path to app.env dynamically
	configPath := filepath.Join(pwd, "app.env")
	log.Info(configPath)

	viper.SetConfigFile(configPath) // Use the dynamic full path to app.env

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})

	viper.Set("PWD", pwd)
}

func init() {
	InitConfig()
	config.Connect()
	models.DB = config.GetDB()
	models.DB.AutoMigrate(&models.Book{})
}

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	hostname := viper.GetString("SERVER_HOST") + ":" + viper.GetString("SERVER_PORT")
	log.Info(hostname)
	log.Fatal(http.ListenAndServe(hostname, r))

}
