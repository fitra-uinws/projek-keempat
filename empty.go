package main

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}
}

func main() {
	var v interface{}

	v = 20
	v = v * 9
}

func main() {
	var v interface{}

	v = 20
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

}
