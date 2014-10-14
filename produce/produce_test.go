package produce

import "testing"

func TestFruitsSet(t *testing.T) {
  fruits := make(Fruits)

  if fruits.Set("Apples", 5); fruits["Apples"] != 5 {
    t.Errorf("Apples should equal 5")
  }
}

func TestVegetablesSet(t *testing.T) {
  vegetables := make(Vegetables)

  if vegetables.Set("Tomatos", 9); vegetables["Tomatos"] != 9 {
    t.Errorf("Tomatos should equal 9")
  }
}
