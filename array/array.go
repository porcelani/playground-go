package array

func Array0()[3]string {
	country := [3]string{}
	return country
}

func Array1() [3]string {
	var country [3]string
	country[0] = "USA"
	country[1] = "Brazil"
	country[2] = "China"
	return country
}

func Array2()[3]string {
	country := [...]string{"USA", "Brazil", "China"}
	return country
}

func Array3()[3]string {
	country := [3]string{"USA", "Brazil", "China"}
	return country
}
