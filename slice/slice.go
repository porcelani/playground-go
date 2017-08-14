package slice

func Slice0() ([]string, []int) {

	country := []string{}
	var code []int

	return country, code
}

func Slice1() []string {

	country := []string{"USA", "Brazil", "China"}

	return country
}

func SliceMake() []string {

	countries := []string{}
	countries = make([]string, 5)

	return countries
}

func SliceMakeWithAdd() []string {

	countries := []string{}
	countries = make([]string, 5)

	countries[0]="Brazil"
	countries[1]="USA"

	return countries
}

