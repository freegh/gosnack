package tdd

import "testing"

func TestFizzBuzz(t *testing.T) {
 testcases := []struct {
  input    int
  expected string
 }{
  {1, "1"},
  {2, "2"},
  {3, "Fizz"},
  {4, "4"},
  {5, "Buzz"},
  {6, "Fizz"},
  {7, "7"},
  {8, "8"},
  {9, "Fizz"},
  {10, "Buzz"},
  {11, "11"},
  {12, "Fizz"},
  {13, "13"},
  {14, "14"},
  {15, "FizzBuzz"},
 }

 for _, v := range testcases {
  actual := FizzBuzz(v.input)
  if actual != v.expected {
   t.Error("expected :", v.expected, " but got ", actual)
  }
 }
}

func TestInputShouldBe1(t *testing.T) {
	//arrage
	actual := FizzBuzz(1)

	//action
	expected := "1"

	//assertion
	if actual != expected {
		t.Error("expected: ", expected, "but got ", actual)
	}
}

func TestInputShouldBe2(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"
	if actual != expected {
		t.Error("expected: ", expected, "but got ", actual)
	}
}

func TestInputShouldBe3(t *testing.T) {
	actual := FizzBuzz(3)
	expected := "Fizz"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe4(t *testing.T) {
	actual := FizzBuzz(4)
	expected := "4"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe5(t *testing.T) {
	actual := FizzBuzz(5)
	expected := "Buzz"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe6(t *testing.T) {
	actual := FizzBuzz(6)
	expected := "Fizz"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe7(t *testing.T) {
	actual := FizzBuzz(7)
	expected := "7"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe8(t *testing.T) {
	actual := FizzBuzz(8)
	expected := "8"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}

func TestInputShouldBe9(t *testing.T) {
	actual := FizzBuzz(9)
	expected := "Fizz"
	if actual != expected {
		t.Errorf("expected %q but got %q\n", expected, actual)
	}
}