package main

import (
	"flag"
	"fmt"
)

var (
	// P is pricipal
	P int
	// nY is years
	nY int
	// R is rate of interest per annum
	R float64
)

func main() {
	flag.IntVar(&P, "p", 50_000_000, "借入金額")
	flag.IntVar(&nY, "ny", 35, "借入年数")
	flag.Float64Var(&R, "r", 0.525, "利息(%)")
	flag.Parse()

	R /= 100
	fmt.Printf("借入金額: %d, 借入年数: %d, 利息: %.3f%%\n", P, nY, 100*R)

	fmt.Println(">>>>>>>>>> 元利均等方式 <<<<<<<<<<")
	matchingTheRepaymentOfPrincipalAndInterest(P, nY, R)
	fmt.Println(">>>>>>>>>> 元金均等方式 <<<<<<<<<<")
	matchingThePrincipalRepayment(P, nY, R)
}

// P: principal
// R: rate of interest per annum
// r: monthly interest rate
// N: times of repayment
// x: replayment monthly of principal
// X: replayment monthly

// 1: P1 = P(1+r)-X
// 2: P2 = P1(1+r)-X = (P(1+r)-X)(1+r)-X = P(1+r)^2-X((1+r)+1)
// 3: P3 = P2(1+r)-X = P(1+r)^3-X((1+r)^2+(1+r)+1)
// ...
// i: Pi = Pi-1(1+r)-X = P(1+r)^i-X((1+r)^(i-1)+(1+r)^(i-2)+...+1) = P(1+r)^i-X((1+r)^i-1)/(1+r-1) = P(1+r)^i-X((1+r)^i-1)/r
// ...
// N: PN = P(1+r)^N-X((1+r)^N-1)/r = 0
//     X = Pr(1+r)^N/((1+r)^N-1)

func matchingTheRepaymentOfPrincipalAndInterest(P int, nY int, R float64) {
	r := R / 12
	N := 12 * nY
	// f: (1+r)^m
	f := 1.0
	for i := 1; i <= N; i++ {
		f *= 1 + r
	}
	X := int(float64(P) * r * f / (f - 1))

	pI := P
	for i := 1; i <= N; i++ {
		rI := int(float64(pI) * r)
		if i%60 == 1 {
			fmt.Printf("%2d年%d月目, 月元金返済額: %6d, 月利息返済額: %5d, 月返済総額: %6d\n", i/12+1, i%12, X-rI, rI, X)
		}
		pI -= X - rI
	}

	fmt.Printf("返済総額:     %d\n", X*N)
	fmt.Printf("返済利息総額: %d\n", X*N-P)
	fmt.Printf("利息割合:     %.2f%%\n", 100*float64(X*N-P)/float64(X*N))
}

func matchingThePrincipalRepayment(B int, nY int, R float64) {
	r := R / 12
	N := 12 * nY
	x := B / N

	pI := B
	var rSum int
	for i := 1; i <= N; i++ {
		rI := int(float64(pI) * r)
		if i%60 == 1 {
			fmt.Printf("%2d年%d月目, 月元金返済額: %6d, 月利息返済額: %5d, 月返済総額: %6d\n", i/12+1, i%12, x, rI, x+rI)
		}
		rSum += rI
		pI -= x
	}

	fmt.Printf("返済総額:     %d\n", B+rSum)
	fmt.Printf("返済利息総額: %d\n", rSum)
	fmt.Printf("利息割合:     %.2f%%\n", 100*float64(rSum)/float64(B+rSum))
}
