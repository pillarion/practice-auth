package access

type Matrix struct {
	Access []Access `db:""`
}

type Access struct {
	ID       int64  `db:"id"`
	Endpoint string `db:"endpoint"`
	Role     string `db:"role"`
}
