package tax

import "testing"

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
