package controllers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"os"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/mlph-kvillegas/events-reservation-system-backend/api/models"
//	jwt "github.com/dgrijalva/jwt-go"
)

type ResponseReturnUserAuth struct{
	StatusCode int `json:"statuscode"`
	StatusMessage string `json:"statusmessage"`
	Auth string `json:"auth"`
};
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
};

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Venue{}, &models.Reservation{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}
func (server *Server) CorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_URL"));
	w.Header().Set("Access-Control-Allow-Methods", "POST");
	w.Header().Set("Connection", "Keep-Alive");
	w.Header().Set("Content-Type", "application/json");

}
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
func(m *ResponseReturnUserAuth ) ToJson() string{
	urlsJson, jsn_err := json.Marshal( m)

	if jsn_err != nil {
        fmt.Println(jsn_err)
	}

	return string(urlsJson);
}