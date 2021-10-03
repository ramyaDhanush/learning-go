package main

import (
	"fmt"
	"math"
	"math/big"
)

func bigFloatExp(x *big.Float, power int64, prec uint) *big.Float {
	one:= new(big.Float).SetPrec(prec).SetInt64(1)
	
	if(power == 0) {
		return one
	}
	if (power ==1) {
		return x
	}
	
	if power < 0 {
		// callback into the functions with a positive power
		// return this value on the denominator
		return one.Quo(one, bigFloatExp(x, -1*power, prec))
	}
	
	halfPower := power/2
	remainder := power%2
	
	halfExp := bigFloatExp(x, halfPower, prec)
	remainderExp := bigFloatExp(x, remainder, prec)
	
	result := new(big.Float).SetPrec(prec)
	result.Mul(halfExp, halfExp)
	result.Mul(result, remainderExp)
	
	return result
}

func main() {
	const prec = 6000;
	
	one := new(big.Float).SetPrec(prec).SetInt64(1)
	base := new(big.Float).SetPrec(prec).SetInt64(2)
	base = base.Quo(one, base)
	power:=int64(54)
	
	result:= bigFloatExp(base, power, prec)
	resultString := result.Text('f', 100)
	
	floatResult := math.Pow(0.5, float64(power))
	fmt.Println("Base %v to the %v power is %v\n", base, power, resultString)
	fmt.Println("Doing it with Float64 %v\n\n", floatResult)
		
	result.Add(result, one)
	resultString = result.Text('f', 100)
	floatResult += 1
	
	fmt.Println("Base %v to the %v power is %v\n", base, power, resultString)
	fmt.Printf("Doing it with float64 %v\n", floatResult)
		
	fmt.Printf("%T - %T\n", one, base)
}