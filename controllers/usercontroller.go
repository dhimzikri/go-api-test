package controllers

import (
	"go_api_cimol/helpers"
	"go_api_cimol/middleware"
	"go_api_cimol/models"
	"net/http"
)

// @Tags MyProfile
// @Summary Get Data Assign FC
// @Router /tools/Me [get]
// @Param Authorization header string false  "string example" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.MyProfile} 
// @Failure 400 {object} helpers.ResponseWithoutData
func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.KeyUserInfo("userinfo")).(*helpers.MyCustomClaims)
	userResponse := &models.MyProfile{
		ID:   user.User_name,
		Name: user.RealName,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
