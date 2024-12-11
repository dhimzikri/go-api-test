package controllers

import (
	"context"
	"encoding/json"
	"go_api_cimol/config"
	"go_api_cimol/helpers"
	"go_api_cimol/models"
	"net/http"
	"time"
)

// @Tags DataAssignFC
// @Summary Get Data Assign FC
// @Router /tools/GetDataAssignFC [post]
// @Param request body models.DataAssignFC true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetDataAssignFC} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetDataAssignFC(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.DataAssignFC
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var getdataassignfc []models.GetDataAssignFC
	ctx := context.Background()
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB2.Raw("Exec getDataAssignFC ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Parameter, dataassignfc.GPS).Scan(&getdataassignfc).Error; err != nil {
			helpers.Response(w, 404, "Data not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	totalData := len(getdataassignfc)

	getdata := struct {
		Total int                      `json:"total"`
		Data  []models.GetDataAssignFC `json:"data"`
	}{
		Total: totalData,
		Data:  getdataassignfc,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags DataAssignFC
// @Summary Find Data Assign FC
// @Router /tools/FindDataAssignFC [post]
// @Param request body models.FDataAssignFC true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.FindDataAssignFC} 
// @Failure 400 {object} helpers.ResponseWithoutData
func FindDataAssignFC(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.FDataAssignFC
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var finddataassignfc []models.FindDataAssignFC
	ctx := context.Background()
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB2.Raw("Exec sp_finddataassignfc ?, ?, ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Parameter, dataassignfc.Searchdata, dataassignfc.GPS, dataassignfc.Orderby).Scan(&finddataassignfc).Error; err != nil {
			helpers.Response(w, 404, "Data not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	totalData := len(finddataassignfc)

	getdata := struct {
		Total int                       `json:"total"`
		Data  []models.FindDataAssignFC `json:"data"`
	}{
		Total: totalData,
		Data:  finddataassignfc,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags MappingResultData
// @Summary Get data Mapping Result
// @Router /tools/GetTblMappingResult [get]
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetTblMappingResult} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetTblMappingResult(w http.ResponseWriter, r *http.Request) {
	var gettblmappingresult []models.GetTblMappingResult

	ctx := context.Background()
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB2.Raw("Exec sp_getTblMappingResult").Scan(&gettblmappingresult).Error; err != nil {
			helpers.Response(w, 404, "Data not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	totalData := len(gettblmappingresult)

	getdata := struct {
		Total int                          `json:"total"`
		Data  []models.GetTblMappingResult `json:"data"`
	}{
		Total: totalData,
		Data:  gettblmappingresult,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags Question
// @Summary Get Question
// @Router /tools/GetTblQuestion [post]
// @Param request body models.DataAssignFC2 true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetTblQuestion} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetTblQuestion(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.DataAssignFC2
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var gettblquestion []models.GetTblQuestion
	if err := config.DB2.Raw("Exec sp_setQuestion ?, ?, ?, ?, ?, ?", dataassignfc.Branchid, dataassignfc.Applicationid, dataassignfc.Agreementno, dataassignfc.Parameter, dataassignfc.Flag, dataassignfc.Seqid).Scan(&gettblquestion).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", gettblquestion)
}

// @Tags LovGroup
// @Summary Get LovGroup
// @Router /tools/GetLovGroup [post]
// @Param request body models.DataAssignFC true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetLovGroup} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetLovGroup(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.LovGroup
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var getlovgroup []models.GetLovGroup
	ctx := context.Background()
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB2.Raw("Exec sp_getlovgroup ?", dataassignfc.LeaderID).Scan(&getlovgroup).Error; err != nil {
			helpers.Response(w, 404, "Data not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	totalData := len(getlovgroup)

	getdata := struct {
		Total int                  `json:"total"`
		Data  []models.GetLovGroup `json:"data"`
	}{
		Total: totalData,
		Data:  getlovgroup,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags CustomerCard
// @Summary GetCustomerCard
// @Router /tools/GetCustomerCard [post]
// @Param request body models.CustomerCard true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetCustomerCard} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetCustomerCard(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.CustomerCard
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	currentTime := time.Now()
	getdate := currentTime.Format("2006-01-02")
	var getcustomercard []models.GetCustomerCard
	if err := config.DB2.Raw("Exec spViewCustomerCard ?, ?", dataassignfc.Applicationid, getdate).Scan(&getcustomercard).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", getcustomercard)
}

// @Tags CollectionHistory
// @Summary Get Collection History
// @Router /tools/GetCollectionHistory [post]
// @Param request body models.GetCollection true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetCollectionHistory} 
// @Failure 400 {object} helpers.ResponseWithoutData
func GetCollectionHistory(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.GetCollection
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var getcollectionhistory []models.GetCollectionHistory
	if err := config.DB2.Raw("Exec spCollActivityHistory ?, 1,1000,'all','',1000, ?", dataassignfc.Customerid, dataassignfc.Applicationid).Scan(&getcollectionhistory).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", getcollectionhistory)
}

// @Tags DataAssignFC
// @Summary Get Data Assign FC History
// @Router /tools/GetDataAssignFcHistory [post]
// @Param request body models.DataAssignFC true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetDataAssignFcHistory}
// @Failure 400 {object} helpers.ResponseWithoutData
func GetDataAssignFcHistory(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.DataAssignFC
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var getdataassignfchistory []models.GetDataAssignFcHistory
	if err := config.DB2.Raw("Exec getDataAssignFC_History ?, ?", dataassignfc.LeaderID, dataassignfc.GPS).Scan(&getdataassignfchistory).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	totalData := len(getdataassignfchistory)

	getdata := struct {
		Total int                             `json:"total"`
		Data  []models.GetDataAssignFcHistory `json:"data"`
	}{
		Total: totalData,
		Data:  getdataassignfchistory,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags DataAssignFC
// @Summary Find Data Assign FC History
// @Router /tools/FindDataAssignFcHistory [post]
// @Param request body models.FDataAssignFC true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.FindDataAssignFcHistory}
// @Failure 400 {object} helpers.ResponseWithoutData
func FindDataAssignFcHistory(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.FDataAssignFC
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var finddataassignfchistory []models.FindDataAssignFcHistory
	if err := config.DB2.Raw("Exec sp_finddataassignfc_History ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Searchdata, dataassignfc.GPS).Scan(&finddataassignfchistory).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", finddataassignfchistory)
}

// @Tags DataColl
// @Summary Get Data Coll
// @Router /tools/GetTblDataColl [post]
// @Param request body models.DataColl true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetTblDataColl}
// @Failure 400 {object} helpers.ResponseWithoutData
func GetTblDataColl(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.DataColl
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var gettbldatacoll []models.GetTblDataColl
	ctx := context.Background()
	err := helpers.DoRetry(ctx, func(ctx context.Context) error {
		if err := config.DB2.Raw("Exec sp_getTblDataColl ?, ?, ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Token, dataassignfc.Searchdata, "", "").Scan(&gettbldatacoll).Error; err != nil {
			helpers.Response(w, 404, "Data not found", nil)
			return helpers.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	totalData := len(gettbldatacoll)

	getdata := struct {
		Total int                     `json:"total"`
		Data  []models.GetTblDataColl `json:"data"`
	}{
		Total: totalData,
		Data:  gettbldatacoll,
	}

	helpers.Response(w, 200, "Successfully Get Data", getdata)
}

// @Tags DataColl
// @Summary Get the Result Data
// @Router /tools/GetTblResultData [post]
// @Param request body models.DataAssignFC2 true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetTblResultData}
// @Failure 400 {object} helpers.ResponseWithoutData
func GetTblResultData(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.DataColl
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var gettblresultdata []models.GetTblResultData
	if err := config.DB2.Raw("Exec sp_getResultDataColl ?, ?, ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Token, dataassignfc.Searchdata, "", "").Scan(&gettblresultdata).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", gettblresultdata)
}

// @Tags DataColl
// @Summary Get data Field Coll HD
// @Router /tools/GetTblFieldCollHD [post]
// @Param request body models.FieldCollHD true "Payload Body [RAW]"
// @Accept json
// @Produces json
// @Param Authorization header string false  "token get from login" token(string)
// @tokenUrl string
// @description OAuth protects our entity endpoints
// @Success 200 {object} helpers.ResponseWithData{data=[]models.GetTblFieldCollHD}
// @Failure 400 {object} helpers.ResponseWithoutData
func GetTblFieldCollHD(w http.ResponseWriter, r *http.Request) {
	var dataassignfc models.FieldCollHD
	if err := json.NewDecoder(r.Body).Decode(&dataassignfc); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var getTblFieldCollHD []models.GetTblFieldCollHD
	if err := config.DB2.Raw("Exec sp_getTblFieldCollHD ?, ?, ?", dataassignfc.LeaderID, dataassignfc.Token, dataassignfc.Agreementno).Scan(&getTblFieldCollHD).Error; err != nil {
		helpers.Response(w, 404, "Data not found", nil)
		return
	}

	helpers.Response(w, 200, "Successfully Get Data", getTblFieldCollHD)
}
