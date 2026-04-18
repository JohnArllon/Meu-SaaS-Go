package domain

type Appointment struct {
	ID                  string `json:"id"`
	CustomerName        string `json:"customer_name"`
	InstallationAddress string `json:"installation_address"`
	BrandModel          string `json:"brand_model"`
	ServiceType         string `json:"service_type"`
	ProblemDescription  string `json:"problem_description"`
	PartsReplaced       string `json:"parts_replaced"`
	TotalValue          string `json:"total_value"`
	PaymentMethod       string `json:"payment_method"`
}
