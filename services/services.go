package services

//CheckErro é uma função resposavel por verificar erro retornado do db
func CheckErro(err error) {
	if err != nil {
		panic(err)
	}
}
