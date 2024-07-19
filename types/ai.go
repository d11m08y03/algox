package types

type BloodDemandRequestPayload struct {
	DistrictName               string `json:"District_Name"`
	Date                       string `json:"Date"`
	Population                 int    `json:"Population"`
	NumberOfHospitals          int    `json:"Number_Of_Hospitals"`
	NumberOfAccidents          int    `json:"Number_Of_Accidents"`
	NumberOfSurgeriesPerformed int    `json:"Number_Of_Surgeries_Performed"`
	BloodDriveEvents           int    `json:"Blood_Drive_Events"`
	Seasonality                string `json:"Seasonality"`
}

type BloodDemandResultPayload struct {
	Demand float32 `json:"demand"`
}
