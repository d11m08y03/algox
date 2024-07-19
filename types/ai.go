package types

type BloodDemandRequestPayload struct {
	Temperature float32 `json:"temperature"`
	BloodType   string  `json:"bloodType"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Population  int     `json:"population"`
	Events      int     `json:"events"`
	Date        string  `json:"date"`
}

type BloodDemandResultPayload struct {
	Demand float32 `json:"demand"`
}
