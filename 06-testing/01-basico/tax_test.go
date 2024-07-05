package tax

import "testing"

//go test .
//go test -v

//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out

//go test -bench=. -run=^#
//go test -bench=. -run=^# -count=10
//go test -bench=. -run=^# -count=10 -benchtime=1s

//go test -fuzz=. -fuzztime=5s -run=^#

//go help test

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	actual := CalculateTax(amount)

	if actual != expected {
		t.Errorf("Expected %f but got %f", expected, actual)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		//		{1.0, 10.0},
	}

	for _, item := range table {
		actual := CalculateTax(item.amount)
		if actual != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, actual)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTaxWithSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		actual := CalculateTax(amount)
		if amount <= 0 && actual != 0 {
			t.Errorf("Received %f but expected 0", actual)
		}
		if amount > 20000 && actual != 20 {
			t.Errorf("Received %f but expected 20", actual)
		}
	})
}
