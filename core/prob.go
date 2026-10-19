package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

type FuzzyBetaNumber struct {
	Alpha float64
	Beta  float64
	A     float64
	B     float64
}

func (f *FuzzyBetaNumber) MembershipFunction(x float64) float64 {
	if x < f.A || x > f.B {
		return 0.0
	}

	normalizedX := (x - f.A) / (f.B - f.A)

	betaDist := distuv.Beta{
		Alpha: f.Alpha,
		Beta:  f.Beta,
	}

	pdf := betaDist.Prob(normalizedX)

	var maxPdf float64 = 1.0
	if f.Alpha > 1 && f.Beta > 1 {
		mode := (f.Alpha - 1) / (f.Alpha + f.Beta - 2)
		maxPdf = betaDist.Prob(mode)
	} else {
		for i := 1; i < 100; i++ {
			t := float64(i) / 100.0
			p := betaDist.Prob(t)
			if p > maxPdf {
				maxPdf = p
			}
		}
	}

	mu := pdf / maxPdf
	if mu > 1.0 {
		return 1.0
	}
	return mu
}

func (f *FuzzyBetaNumber) ProbabilityDistributionFunction(x float64) float64 {
	if x < f.A {
		return 0.0
	}
	if x > f.B {
		return 1.0
	}

	normalizedX := (x - f.A) / (f.B - f.A)

	betaDist := distuv.Beta{
		Alpha: f.Alpha,
		Beta:  f.Beta,
	}

	return betaDist.CDF(normalizedX)
}

func main() {
	// Создаем нечеткую величину с параметрами alpha=2.0, beta=5.0 на носителе [2.0, 10.0]
	fuzzyNum := FuzzyBetaNumber{
		Alpha: 2.0,
		Beta:  5.0,
		A:     2.0,
		B:     10.0,
	}

	// Тестируем в нескольких точках
	testPoints := []float64{1.0, 2.0, 4.0, 6.0, 10.0, 12.0}

	fmt.Println("Тестирование нечеткой величины (Бета-распределение):")
	for _, x := range testPoints {
		mu := fuzzyNum.MembershipFunction(x)
		cdf := fuzzyNum.ProbabilityDistributionFunction(x)
		fmt.Printf("x = %5.1f | μ(x) (принадлежность) = %.4f | CDF (вероятность) = %.4f\n", x, mu, cdf)
	}
}
