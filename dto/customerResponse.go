package dto

type CustomerResponse struct {
	Id      string `json:"customer_id"`
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	DOB     string `json:"date_of_birth"`
	Status  string `json:"status"`
}
