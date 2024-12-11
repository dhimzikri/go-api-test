package models

type GetDataAssignFC struct {
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Namadebitur   string `json:"namadebitur"`
	Mailingadress string `json:"mailingadress"`
	Zipcode       string `json:"zipcode"`
	Daysoverdue   string `json:"daysoverdue"`
	Ptpydate      string `json:"ptpydate"`
	Actionresult  string `json:"actionresult"`
	Saldopokok    string `json:"saldopokok"`
	Angske        string `json:"angske"`
	Parameter     string `json:"parameter"`
	Bucketid      string `json:"bucketid"`
	Jarak         string `json:"jarak"`
	Radius        string `json:"radius"`
	Cgid          string `json:"cgid"`
	Isaqt         int    `json:"isaqt"`
	Ispaid        int    `json:"is_paid"`
	Flagsort      int    `json:"flagsort"`
	Gender        string `json:"gender"`
	Agentname     string `json:"agentname"`
	GPS           string `json:"gps"`
	Colorflagsort string `json:"colorflagsort"`
	Lastactivity  string `json:"lastactivity"`
}

type DataAssignFC struct {
	LeaderID  string `json:"username" example:"1879"`
	Parameter string `json:"parameter" example:"-1"`
	GPS       string `json:"gps" example:"0;0"`
}

type FDataAssignFC struct {
	LeaderID   string `json:"username" example:"1879"`
	Parameter  string `json:"parameter" example:"-1"`
	Searchdata string `json:"searchdata" example:""`
	GPS        string `json:"gps" example:"0;0"`
	Orderby    string `json:"orderby" example:"applicationid"`
}

type GetCollection struct {
	Customerid    string `json:"customerid"`
	Applicationid string `json:"applicationid"`
}

type DataAssignFC2 struct {
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Parameter     string `json:"parameter"`
	Flag          int    `json:"flag"`
	Seqid         string `json:"seqid"`
}

type FindDataAssignFC struct {
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Namadebitur   string `json:"namadebitur"`
	Mailingadress string `json:"mailingadress"`
	Zipcode       string `json:"zipcode"`
	Daysoverdue   string `json:"daysoverdue"`
	Ptpydate      string `json:"ptpydate"`
	Actionresult  string `json:"actionresult"`
	Saldopokok    string `json:"saldopokok"`
	Angske        string `json:"angske"`
	Parameter     string `json:"parameter"`
	Bucketid      string `json:"bucketid"`
	Jarak         string `json:"jarak"`
	Radius        string `json:"radius"`
	Cgid          string `json:"cgid"`
	Isaqt         int    `json:"isaqt"`
	Ispaid        int    `json:"is_paid"`
	Flagsort      int    `json:"flagsort"`
	Gender        string `json:"gender"`
	Agentname     string `json:"agentname"`
	GPS           string `json:"gps"`
	Colorflagsort string `json:"colorflagsort"`
	Lastactivity  string `json:"lastactivity"`
}

type GetTblMappingResult struct {
	Resultid    string `json:"resultid"`
	Refid       string `json:"refid"`
	Ismandatory int    `json:"ismandatory"`
}

type GetTblQuestion struct {
	PK            int    `json:"pk"`
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Customerid    string `json:"customerid"`
	Refid         string `json:"refid"`
	Answer        string `json:"answer"`
	Flag          string `json:"flag"`
	Bussinessdate string `json:"bussinessdate"`
	Lovgroup      string `json:"lovgroup"`
	Questionlabel string `json:"questionlabel"`
	Foto          string `json:"foto"`
	Type          string `json:"type_"`
	Date          string `json:"date"`
}

type LovGroup struct {
	LeaderID string `json:"username" example:"1879"`
}

type GetLovGroup struct {
	Code string `json:"code"`
}

type CustomerCard struct {
	Applicationid string `json:"applicationid" example:"421A202208016092"`
}

type GetCustomerCard struct {
	Applicationid     string `json:"applicationid"`
	Insseqno          string `json:"insseqno"`
	Duedate           string `json:"duedate"`
	Installmentamount string `json:"installmentamount"`
	Paidamount        string `json:"paidamount"`
}

type GetCollectionHistory struct {
	Activitydate string `json:"activitydate"`
	Activity     string `json:"activity"`
	Resultid     string `json:"resultid"`
	Result       string `json:"result"`
}

type GetDataAssignFcHistory struct {
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Namadebitur   string `json:"namadebitur"`
	Mailingadress string `json:"mailingadress"`
	Zipcode       string `json:"zipcode"`
	Daysoverdue   string `json:"daysoverdue"`
	Ptpydate      string `json:"ptpydate"`
	Actionresult  string `json:"actionresult"`
	Saldopokok    string `json:"saldopokok"`
	Angske        string `json:"angske"`
	Parameter     string `json:"parameter"`
	Bucketid      string `json:"bucketid"`
	Jarak         string `json:"jarak"`
	Radius        string `json:"radius"`
	Cgid          string `json:"cgid"`
	Isaqt         int    `json:"isaqt"`
	Ispaid        int    `json:"is_paid"`
	Flagsort      int    `json:"flagsort"`
	Gender        string `json:"gender"`
	Agentname     string `json:"agentname"`
	GPS           string `json:"gps"`
	Colorflagsort string `json:"colorflagsort"`
	Lastactivity  string `json:"lastactivity"`
}

type FieldCollHD struct {
	LeaderID   string `json:"username" example:"1879"`
	Token      string `json:"token" example:"$2y$10$LfW07MZ0QZaXYv093xT0huWA8o1gWwWhJg3UcrXdAhJq47OpHm.5e"`
	Agreementno string `json:"agreementno"`
}

type GetTblFieldCollHD struct {
	PK            int    `json:"pk"`
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Customerid    string `json:"customerid"`
	Refid         string `json:"refid"`
	Questionlabel string `json:"questionlabel"`
	Answer        string `json:"answer"`
	Flag          string `json:"flag"`
	Bussinessdate string `json:"bussinessdate"`
	Type          string `json:"type_"`
	Plandate      string `json:"plandate"`
}

type FindDataAssignFcHistory struct {
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	Namadebitur   string `json:"namadebitur"`
	Mailingadress string `json:"mailingadress"`
	Zipcode       string `json:"zipcode"`
	Daysoverdue   string `json:"daysoverdue"`
	Ptpydate      string `json:"ptpydate"`
	Actionresult  string `json:"actionresult"`
	Saldopokok    string `json:"saldopokok"`
	Angske        string `json:"angske"`
	Parameter     string `json:"parameter"`
	Bucketid      string `json:"bucketid"`
	Cgid          string `json:"cgid"`
	Isaqt         int    `json:"isaqt"`
	Ispaid        int    `json:"is_paid"`
	Flagsort      int    `json:"flagsort"`
	Gender        string `json:"gender"`
	Seqid         string `json:"seqid"`
}

type GetTblResultData struct {
	PK            int    `json:"pk"`
	Branchid      string `json:"branchid"`
	Applicationid string `json:"applicationid"`
	Agreementno   string `json:"agreementno"`
	GPS           string `json:"gps"`
	Plandate      string `json:"plandate"`
	Zipcode       string `json:"zipcode"`
	Customername  string `json:"customername"`
	Alamat        string `json:"alamat"`
	Foto          string `json:"foto"`
	Flag          string `json:"flag"`
}

type DataColl struct {
	LeaderID   string `json:"username" example:"1879"`
	Token      string `json:"token" example:"$2y$10$LfW07MZ0QZaXYv093xT0huWA8o1gWwWhJg3UcrXdAhJq47OpHm.5e"`
	Searchdata string `json:"searchdata"`
}

type GetTblDataColl struct {
	PK          int    `json:"pk"`
	Agentid     string `json:"agentid"`
	Name        string `json:"name"`
	Notactivity string `json:"notactivity"`
	Activity    string `json:"activity"`
}
