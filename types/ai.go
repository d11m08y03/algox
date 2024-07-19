package types

type BloodDemandRequestPayload struct {
	Date      string `json:"Date"`
	BloodType string `json:"bloodType"`
	Hospital  string `json:"hospital"`
}

type BloodDemandResultPayload struct {
	Demand float32 `json:"demand"`
	Stock  float32 `json:"stock"`
}
