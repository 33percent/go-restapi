package keys

type mysqlconnection struct {
	DB *keysdefine
}

type keysdefine struct {
	Dialect  string
	Username string
	Password string
	Name     string
}

func makeconnection() *mysqlconnection {
	return &mysqlconnection{
		DB: &keysdefine{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "test",
		},
	}
}
