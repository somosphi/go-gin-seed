package v1

type (
	Todo struct {
		Title     string `json:"title" validate:"required,min=3" bind:"required"`
		Completed bool   `json:"completed"`
	}
)
