package cntrllrs

type decToBin struct{
	Dec int
  	Bin string
}

func NewDecToBin() *decToBin{
  	obj := new(decToBin)
  	return obj
}
