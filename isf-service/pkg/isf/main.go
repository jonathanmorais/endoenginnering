package isf

func InsuFactor() float64 {
	// insuline proportion per carbo is 0.04 or 1:25
	var bgl float64
	var iu float64
	bgl = 6
	iu = 1

	return iu / bgl

}
