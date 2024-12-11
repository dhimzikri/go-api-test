// This 
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go_api_cimol/config"
	"go_api_cimol/helpers"
	"go_api_cimol/models"

	"net/http"
)

// @Tags AUTH
// @Summary Login into API Cimol
// @Router /auth/login [post]
// @Param request body models.Login true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Success 200 {object} models.UserLogin
// @Failure 400 {object} helpers.ResponseWithoutData
func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	ctx := context.Background()
	// err := helpers.DoRetry(ctx, func(ctx context.Context) error {
	// 	if err := config.DB1.Raw("Exec sp_getUserMobile_Coll ?, ?", login.Username, login.Password).Scan(&user).Error; err != nil {
	// 		helpers.Response(w, 404, "Article not found", nil)
	// 		return helpers.RetryableError(err)
	// 	}
	// 	return nil
	// })
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB1.Raw("Exec sp_getUserMobile_Coll ?, ?", login.Username, login.Password).Scan(&user).Error; err != nil {
			// Log pesan ketika terjadi retry
			fmt.Printf("Retrying request after error: %v\n", err)
			helpers.Response(w, 404, "Article not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// var userlogin models.UserLogin
	// if err := helpers.VerifyPassword(userlogin.UserPassword, login.Password); err != nil {
	// 	helpers.Response(w, 404, "Wrong Email or Password", nil)
	// 	return
	// }

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	responseData := map[string]interface{}{
		"data_user": user,
		"token":     token,
	}

	helpers.Response(w, 200, "Successfully Login", responseData)
}

// @Tags AUTH
// @Summary Logout API Cimol
// @Router /auth/logout [post]
// @Param request body models.Login true "Payload Body [RAW]"
// @Param Authorization header string false  "string example" token(string)
// @Accept json
// @Produces json
// @Success 200 {object} helpers.ResponseWithoutData 
// @Failure 400 {object} helpers.ResponseWithoutData
func Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	// Tambahkan token ke daftar yang sudah logout
	helpers.AddLoggedOutToken(token)

	// Beri respons ke client bahwa logout berhasil
	helpers.Response(w, http.StatusOK, "Logout successful", nil)
}
