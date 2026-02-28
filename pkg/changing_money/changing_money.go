package changing_money

import "sort"

type Cassette interface {
	GetId() int
	IsIntact() bool       // флаг исправности кассеты
	GetNominal() int64    // номинал купюры (ожидается один из: 5000, 2000, 1000, 500, 200, 100, 50, 10, 5, 2, 1)
	GetNumOfBonds() int64 // количество купюр в кассете
}

type distribution struct {
	d          map[int64]int64
	remains    int64
	numOfBonds int64
}

// state
//
//	хранит состояние поиска
//	поля:
//		cassettes      []Cassette    -- список кассет
//		resultFound    bool	         -- флаг результата. true, если он найден
//		sum            int64         -- текущий остаток размениваемой суммы
//		currIdx        int 	         -- индекс текущей рассматриваемой кассеты
//		prevIdx        int           -- индекс предыдущей рассматриваемой кассеты
//		prevNominalIdx int           -- индекс предыдущей кассеты, из которой были взяты купюры для размена
//										(с номиналом больше, чем в текущей кассете)
//		res            map[int]int64 -- результат поиска
//
// *
type state struct {
	nominals      map[int64]int64
	resultFound   bool
	sum           int64 // текущий остаток размениваемой суммы
	distributions []distribution
	ordered       [][]int64
}

var interestingNominals = map[int64]interface{}{
	5000: nil,
	500:  nil,
}

// ChangeMoney
//
//	cassettes []Cassette -- список касет с купюрами
//	sum int64 -- сумма для размена
//
//	Порядок поиска:
//	1. сортировка по убыванию номинала купюр
//	2. ищем размен
//	3. ищем остаточный размен(см. findFinalChanging())
//
// *
func ChangeMoney(cassettes []Cassette, sum int64) (map[int]int64, bool) {
	st := &state{
		nominals:      getNominals(cassettes),
		resultFound:   false,
		sum:           sum,
		distributions: make([]distribution, 0, 4),
	}

	makeChanging(st)
	if len(st.distributions) <= 0 {
		return map[int]int64{}, false
	}

	d := st.distributions[0]
	res := make(map[int]int64)
	for _, c := range cassettes {
		if !c.IsIntact() || c.GetNumOfBonds() <= 0 {
			continue
		}

		if need, ok := d.d[c.GetNominal()]; ok {
			res[c.GetId()] = min(need, c.GetNumOfBonds())
			d.d[c.GetNominal()] -= res[c.GetId()]
		}
	}

	return res, true
}

func getNominals(cassettes []Cassette) map[int64]int64 {
	nominals := make(map[int64]int64, 6)
	for _, cas := range cassettes {
		if !cas.IsIntact() || cas.GetNumOfBonds() <= 0 {
			continue
		}
		if _, ok := nominals[cas.GetNominal()]; ok {
			nominals[cas.GetNominal()] += cas.GetNumOfBonds()
		} else {
			nominals[cas.GetNominal()] = cas.GetNumOfBonds()
		}
	}

	return nominals
}

func makeChanging(st *state) {
	distrib := findChangingGreedily(st)
	if distrib.remains == 0 {
		st.distributions = append(st.distributions, distrib)
	}

	st.ordered = make([][]int64, 0, len(distrib.d))
	for k, v := range distrib.d {
		st.ordered = append(st.ordered, []int64{k, v})
	}
	sort.Slice(st.ordered, func(i, j int) bool {
		return st.ordered[i][0] > st.ordered[j][0]
	})

	var tempNumOf int64
	for i, v := range st.ordered {
		if _, ok := interestingNominals[v[0]]; ok {
			tempNumOf = v[1]
			makeChangingFrom(st, i)
			st.ordered[i][1] = tempNumOf
		}
	}

	sort.Slice(st.distributions, func(i, j int) bool {
		return st.distributions[i].numOfBonds < st.distributions[j].numOfBonds
	})
}

func makeChangingFrom(st *state, start int) {
	var (
		nominal, numOf int64
		d              distribution
	)

	for i := start; i < len(st.ordered); i++ {
		nominal, numOf = st.ordered[i][0], st.ordered[i][1]
		if _, ok := interestingNominals[nominal]; ok {
			st.nominals[nominal] = numOf - 1
			d = findChangingGreedily(st)
			if d.remains == 0 {
				st.distributions = append(st.distributions, d)
			}
		}
	}
}

// findChangingGreedily
//
// /
func findChangingGreedily(st *state) distribution {
	sum := st.sum
	nominals := getOrdered(st.nominals)
	res := make(map[int64]int64, len(st.nominals))

	var nom, numOf, totalNumOf int64
	for _, p := range nominals {
		nom, numOf = p[0], p[1]
		if sum >= nom {
			var numOfShouldBonds = sum / nom
			if numOfShouldBonds <= numOf {
				res[nom] = numOfShouldBonds
				totalNumOf += numOfShouldBonds
				sum -= numOfShouldBonds * nom
			} else {
				res[nom] = numOf
				totalNumOf += numOf
				sum -= numOf * nom
			}
		}
	}

	return distribution{d: res, remains: sum, numOfBonds: totalNumOf}
}

func getOrdered(noms map[int64]int64) [][]int64 {
	nominals := make([][]int64, 0, 6)
	for k, v := range noms {
		nominals = append(nominals, []int64{k, v})
	}
	sort.Slice(nominals, func(i, j int) bool {
		return nominals[i][0] > nominals[j][0]
	})

	return nominals
}
