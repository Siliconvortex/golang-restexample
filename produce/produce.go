package produce 

type Payload struct {
  Stuff Data
}

type Data struct {
  Fruit Fruits
  Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func (fruits *Fruits) Set(f string, i int) {
  (*fruits)[f] = i
}

func (vegetables *Vegetables) Set(v string, i int) {
  (*vegetables)[v] = i
}
