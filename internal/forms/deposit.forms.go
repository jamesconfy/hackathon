package forms

type Deposit struct {
	Status string `json:"status" validate:"required,oneof=Rejected Pending Done"`
}
