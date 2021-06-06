package balancer

import (
	"chemistry/exception"
	"errors"
	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func isBalanced(formula string) bool {
	balanced := true
	s := make([]rune, 0, len(formula))

	for _, c := range formula {
		if c == '(' {
			s = append(s, c)
		} else if c == ')' {
			if len(s) == 0 {
				balanced = false
				break
			}
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		balanced = false
	}

	return balanced
}

func _GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func _LCM(a, b int64) int64 {
	result := a * b / _GCD(a, b)

	return result
}

func printEquation(equation string) (string, error) {
	var sides []string
	var leftSide []string
	var rightSide []string

	var result string
	var err error = nil

	exception.Block{
		Try: func() {
			equation = strings.ReplaceAll(equation, " ", "")
			sides = strings.Split(equation, "=")
			leftSide = strings.Split(sides[0], "+")
			rightSide = strings.Split(sides[1], "+")
		},
		Catch: func(e exception.Exception) {
			err = errors.New("unreadable equation")
		},
	}.Do()

	if err != nil {
		return result, err
	}

	var leftCompounds []map[string]int
	var rightCompounds []map[string]int
	s1 := make(map[string]bool)
	s2 := make(map[string]bool)
	for i := 0; i < len(leftSide); i++ {
		leftCompounds = append(leftCompounds, parse(leftSide[i]))
		for k := range parse(leftSide[i]) {
			s1[k] = true
		}
	}
	for i := 0; i < len(rightSide); i++ {
		rightCompounds = append(rightCompounds, parse(rightSide[i]))
		for k := range parse(rightSide[i]) {
			s2[k] = true
		}
	}

	result = leftSide[0]
	for idx, v := range leftSide {
		if idx == 0 {
			continue
		}
		result += " + " + v
	}
	result += " = "

	gases := []string{"NH3", "AsH3", "CH4", "C6H6", "N2", "CO", "CCl4", "F2", "ClO2", "C2H4", "Cl2", "H2", "ClO2", "H2S", "NO", "NO2", "O3", "SO2", "O2", "CO2", "C2H2", "N2O5", "SO3", "H2S"}
	sediments := []string{"Mg(OH)2", "Al(OH)3", "Sn(OH)2", "Pb(OH)2", "Cr(OH)3", "Mn(OH)2", "Fe(OH)2", "Fe(OH)3", "Co(OH)2", "Ni(OH)2", "Cu(OH)2", "Zn(OH)2", "Cd(OH)2", "CaF2", "SrF2", "HgF2", "AgCl", "AgBr", "AgI", "HgI2", "SrSO4", "BaSO4", "PbSO4", "HgSO4", "CaSO3", "SrSO3", "SnSO3", "PbSO3", "MnSO3", "FeSO3", "CoSO3", "NiSO3", "Ag2SO3", "SrS", "SnS", "PbS", "MnS", "FeS", "Fe2S3", "CoS", "NiS", "CuS", "Ag2S", "ZnS", "CdS", "HgS", "Sn(NO3)2", "Li3PO4", "Ca3(PO4)2", "Sr3(PO4)2", "Ba3(PO4)2", "AlPO4", "Sn3(PO4)2", "Pb3(PO4)2", "CrPO4", "Mn3(PO4)2", "Fe3(PO4)2", "FePO4", "Co3(PO4)2", "Cu3(PO4)2", "Ni3(PO4)2", "Ag3PO4", "Zn3(PO4)2", "Cd3(PO4)2", "Hg3(PO4)2", "CaCO3", "SrCO3", "BaCO3", "SnCO3", "PbCO3", "MnCO3", "FeCO3", "CoCO3", "NiCO3", "Ag2CO3", "ZnCO3", "CdCO3", "HgCO3", "H2SiO3", "MgSiO3", "CaSiO3", "SrSiO3", "BaSiO3", "Al2(SiO3)3", "SnSiO3", "PbSiO3", "FeSiO3", "CoSiO3", "CdSiO3"}
	if stringInSlice(rightSide[0], gases) {
		rightSide[0] += "[G]"
	}
	if stringInSlice(rightSide[0], sediments) {
		rightSide[0] += "[S]"
	}
	result += rightSide[0]
	for idx, v := range rightSide {
		if idx == 0 {
			continue
		}
		if stringInSlice(v, gases) {
			v += "[G]"
		}
		if stringInSlice(v, sediments) {
			v += "[S]"
		}

		result += " + " + v

	}

	return result, err
}

func balance(equation string) (string, error) {
	var sides []string
	var leftSide []string
	var rightSide []string

	var result string
	var err error = nil

	exception.Block{
		Try: func() {
			equation = strings.ReplaceAll(equation, " ", "")
			sides = strings.Split(equation, "=")
			leftSide = strings.Split(sides[0], "+")
			rightSide = strings.Split(sides[1], "+")
		},
		Catch: func(e exception.Exception) {
			err = errors.New("unreadable equation")
		},
	}.Do()

	if err != nil {
		return result, err
	}

	var leftCompounds []map[string]int
	var rightCompounds []map[string]int
	s1 := make(map[string]bool)
	s2 := make(map[string]bool)
	for i := 0; i < len(leftSide); i++ {
		leftCompounds = append(leftCompounds, parse(leftSide[i]))
		for k := range parse(leftSide[i]) {
			s1[k] = true
		}
	}
	for i := 0; i < len(rightSide); i++ {
		rightCompounds = append(rightCompounds, parse(rightSide[i]))
		for k := range parse(rightSide[i]) {
			s2[k] = true
		}
	}
	var elements []string
	for k := range s1 {
		elements = append(elements, k)
	}
	sort.Strings(elements)
	indexes := make(map[string]int)
	for i := range elements {
		indexes[elements[i]] = i
	}
	numCols := len(leftCompounds) + len(rightCompounds)
	numRows := len(elements)

	arr := make([][]int64, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int64, numCols)
	}

	for col, compound := range leftCompounds {
		for el, num := range compound {
			row := indexes[el]
			arr[row][col] = int64(num)
		}
	}
	col := len(leftCompounds)
	for _, compound := range rightCompounds {
		for el, num := range compound {
			row := indexes[el]
			arr[row][col] = int64(-num)
		}
		col += 1
	}

	nr := func(i int64) rational.Rational {
		return rational.New(i, 1)
	}

	var equations [][]rational.Rational
	for _, v := range arr {
		var eq []rational.Rational
		for _, x := range v {
			eq = append(eq, nr(x))
		}
		eq = append(eq, nr(0))
		equations = append(equations, eq)
	}

	res, gausErr := gaussian.SolveGaussian(equations, false)
	if gausErr != nil {
		log.Fatal(gausErr)
	}

	var lcm int64
	lcm = 1
	for _, v := range res {
		for _, x := range v {
			if x.GetDenominator() != 0 {
				lcm = _LCM(lcm, x.GetDenominator())
			}
		}
	}

	coeffs := make([]int64, 0)
	for _, v := range res {
		for _, x := range v {
			x = x.MultiplyByNum(lcm)
			if x.GetNumerator() != 0 {
				coeffs = append(coeffs, int64(math.Abs(float64(x.GetNumerator()))))
			}
		}
	}

	coeffs = append(coeffs, lcm)
	if (int(coeffs[0])) == 1 {
		result = " " + leftSide[0]
	} else {
		result = strconv.Itoa(int(coeffs[0])) + " " + leftSide[0]
	}
	for idx, v := range leftSide {
		if idx == 0 {
			continue
		}
		if (int(coeffs[idx])) == 1 {
			result += " + " + v
		} else {
			result += " + " + strconv.Itoa(int(coeffs[idx])) + " " + v
		}
	}
	result += " = "

	gases := []string{"NH3", "AsH3", "CH4", "C6H6", "N2", "CO", "CCl4", "F2", "ClO2", "C2H4", "Cl2", "H2", "ClO2", "H2S", "NO", "NO2", "O3", "SO2", "O2", "CO2", "C2H2", "N2O5", "SO3", "H2S"}
	sediments := []string{"Mg(OH)2", "Al(OH)3", "Sn(OH)2", "Pb(OH)2", "Cr(OH)3", "Mn(OH)2", "Fe(OH)2", "Fe(OH)3", "Co(OH)2", "Ni(OH)2", "Cu(OH)2", "Zn(OH)2", "Cd(OH)2", "CaF2", "SrF2", "HgF2", "AgCl", "AgBr", "AgI", "HgI2", "SrSO4", "BaSO4", "PbSO4", "HgSO4", "CaSO3", "SrSO3", "SnSO3", "PbSO3", "MnSO3", "FeSO3", "CoSO3", "NiSO3", "Ag2SO3", "SrS", "SnS", "PbS", "MnS", "FeS", "Fe2S3", "CoS", "NiS", "CuS", "Ag2S", "ZnS", "CdS", "HgS", "Sn(NO3)2", "Li3PO4", "Ca3(PO4)2", "Sr3(PO4)2", "Ba3(PO4)2", "AlPO4", "Sn3(PO4)2", "Pb3(PO4)2", "CrPO4", "Mn3(PO4)2", "Fe3(PO4)2", "FePO4", "Co3(PO4)2", "Cu3(PO4)2", "Ni3(PO4)2", "Ag3PO4", "Zn3(PO4)2", "Cd3(PO4)2", "Hg3(PO4)2", "CaCO3", "SrCO3", "BaCO3", "SnCO3", "PbCO3", "MnCO3", "FeCO3", "CoCO3", "NiCO3", "Ag2CO3", "ZnCO3", "CdCO3", "HgCO3", "H2SiO3", "MgSiO3", "CaSiO3", "SrSiO3", "BaSiO3", "Al2(SiO3)3", "SnSiO3", "PbSiO3", "FeSiO3", "CoSiO3", "CdSiO3"}

	if stringInSlice(rightSide[0], gases) {
		rightSide[0] += "[G]"
	}
	if stringInSlice(rightSide[0], sediments) {
		rightSide[0] += "[S]"
	}
	substanceCoeff := int(coeffs[len(leftSide)])
	if substanceCoeff == 1 {
		result += rightSide[0]
	} else {
		result += strconv.Itoa(int(coeffs[len(leftSide)])) + " " + rightSide[0]
	}
	for idx, v := range rightSide {
		if idx == 0 {
			continue
		}
		if stringInSlice(v, gases) {
			v += "[G]"
		}
		if stringInSlice(v, sediments) {
			v += "[S]"
		}
		substanceCoefficient := int(coeffs[len(leftSide)+idx])
		if substanceCoefficient == 1 {
			result += " + " + v
		} else {
			result += " + " + strconv.Itoa(substanceCoefficient) + " " + v
		}
	}

	return result, err
}

func parse(formula string) map[string]int {
	var maps = make(map[string]int)
	var multiple []int
	str, count := "", ""
	for i := len(formula) - 1; i >= 0; i-- {
		char := formula[i]
		if char >= 48 && char <= 57 {
			count = string(char) + count
		} else {
			if char == 41 {
				atoi, _ := strconv.Atoi(count)
				if len(multiple) > 0 {
					atoi = atoi * multiple[len(multiple)-1]
				}
				multiple = append(multiple, atoi)
				count = ""
			} else if char == 40 {
				multiple = multiple[:len(multiple)-1]
			} else if char >= 97 && char <= 122 {
				str = string(char) + str
			} else if char >= 65 && char <= 90 {
				str = string(char) + str
				nums := 1
				if count == "" {
					count = "1"
				}
				atoi, _ := strconv.Atoi(count)
				if len(multiple) > 0 {
					nums = multiple[len(multiple)-1] * atoi
				} else {
					nums = atoi
				}
				maps[str] += nums
				str = ""
				count = ""
			}
		}
	}

	return maps
}

func Balance(equation string) (string, error) {

	//equation := "NaOH + H3PO4 = Na3PO4 + H2O"
	//equation := "NaOH + H2SO4 = NaSO4 + H2O" nu merge
	//equation := "Fe2O3 + C = Fe+CO2"
	//equation := "PCl5 + H2O = H3PO4 + HCl"
	//equation := "KNO3 + C12H22O11 = N2 + CO2 + H2O + K2CO3"
	//equation := "C2H5OH + O2 = CO + H2O"
	//equation := "Mg(OH)2 + H3PO4 = H2O + Mg3(PO4)2"
	//equation := "Ca3(PO4)2 + SiO2 + C = CaSiO3 + CO + P"
	//equation = "Ni(NO3)2 + NaOH = NaNO3 + Ni(OH)2"

	if isBalanced(equation) {
		return balance(equation)
	}
	return printEquation(equation)
}
