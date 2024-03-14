package access

type Matrix struct {
	Access []Access `json:""`
}

type Access struct {
	Endpoint string `json:"endpoint"`
	Role     string `json:"role"`
}
