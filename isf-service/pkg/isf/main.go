package isf

func InsuFactor() float32 {
	// insuline proportion per carbo is 0.04 or 1:25
	var bgl float32
	var iu float32
	bgl = 6
	iu = 1

	return iu / bgl

}
