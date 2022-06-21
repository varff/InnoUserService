package models

type User struct {
	Username  string
	Password  string
	Phone     int32
	Email     string
	Rate      int32
	IsAnalyst string
}

func UserLogin(phone int32) (string, error) {
	var result User

	return result.Password, nil
}

func UserRegister(Name, Password, Email string, Phone int32) (bool, error) {

	return true, nil
}

func UserCheckRate(phone int32) (int32, error) {

	return 0, nil
}

func UserRateOrder(phone, rate int32) error {
	return nil
}
