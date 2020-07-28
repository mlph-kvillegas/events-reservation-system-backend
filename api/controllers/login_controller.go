package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mlph-kvillegas/events-reservation-system-backend/api/auth"
	"github.com/mlph-kvillegas/events-reservation-system-backend/api/models"
	"github.com/mlph-kvillegas/events-reservation-system-backend/api/responses"
//	formaterror "github.com/mlph-kvillegas/events-reservation-system-backend/api/utils"
	"golang.org/x/crypto/bcrypt"

	
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	server.CorsHeader(w);
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//return
	}
	//token, err := server.SignIn(user.Email, user.Password)
	token := server.SignIn(user.Email, user.Password)
	//if err != nil {
	//	formattedError := formaterror.FormatError(err.Error())
	//	responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
	//	return
	//}
	
	
	
	urlsJson:=token.ToJson()
	
	responses.JSON(w, http.StatusOK, string(urlsJson))
}

//func (server *Server) SignIn(email, password string) (string, error) {
func (server *Server) SignIn(email string, password string) *ResponseReturnUserAuth {	

	var err error

	user := models.User{}
	req_data_format := &ResponseReturnUserAuth{
		StatusCode:0,
		StatusMessage:"",
	}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		//return "", err
		req_data_format.StatusCode = 2;
		req_data_format.StatusMessage = "Email address does not exist in our records";

	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		//return "", err
		req_data_format.StatusCode = 3;
		req_data_format.StatusMessage = "Password mismatch";
	}
	if req_data_format.StatusCode==0 {
		req_data_format.StatusCode = 1;
		req_data_format.StatusMessage = "Successfully login";
		auth_val,_:=auth.CreateToken(user.ID);
		req_data_format.Auth = auth_val;
	}
	//return auth.CreateToken(user.ID)
	return req_data_format
}
