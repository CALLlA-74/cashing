package changing_money

import "sort"

/*
		В данном модуле представлена реализация алгоритма размена для следующих возможных номиналов:
	5000, 2000, 1000, 500, 200, 100, 50, 10, 5, 2, 1. Данная система каноническая, за исключением пар
	некратных номиналов: 5000 и 2000; 500 и 200. Т.к. для 5000 удвоение делает его кратным 2000, то
	при поиске решения жадным алгоритмом мы можем перебрать номинал 5000 максимум на 1. Аналогичное справедливо
	и для пары 500 и 200. В данной реалзации для значений 5000 и 500 введено понятие "интересующие" номиналы.
		Поиск основан на жадном алгоритме. При поиске разменов для каждого номинала пытаемся
	взять максимальное количество купюр (из имеющихся) до превышения остатка суммы размена. А для
	"интересующего" номинала вдобавок рассматриваем вариант взятия купюр на 1 меньше, чем брали изначально.

		Порядок поиска:
		1. агрегируем кассеты по номиналам, отсеивая неисправные и пустые кассеты.
		2. сортируем номиналы в порядке убывания.
		3. рекурсивно перебираем номиналы (по одному за вызов):
		3.1. если встретили "интересующий" номинал, то проверяем варианты размена:
		3.1.1. пробуем взять максимальное количество до превышения остатка
	суммы размена (если купюр этого номинала меньше, то берем все)
		3.1.2. и на 1 меньше, чем в 3.1.1
		3.2. когда проверили все номиналы сохраняем полученный размен в виде {номинал: требуемое количество}
		4. сортируем найденные размены по количеству купюр -- выбираем минимальный
		5. для минимального размена составляем распределение по кассетам вида:
	{id кассеты: количество купюр из нее}
*/

type Cassette interface {
	GetId() int
	IsIntact() bool       // исправность кассеты (true -- исправна)
	GetNominal() int64    // номинал купюры
	GetNumOfBonds() int64 // количество купюр в кассете
}

type distribution struct {
	d          map[int64]int64
	remains    int64
	numOfBonds int64
}

// state
// хранит текущее состояние поиска
//
// *
type state struct {
	nominals      [][]int64       // в каждой ячейке слайса хранится пара (слайс): номинал -- количество купюр
	sum           int64           // текущий остаток размениваемой суммы
	distributions []distribution  // найденные варианты разменов
	currD         map[int64]int64 // текущий вариант размена
	currIdx       int             // индекс текущего рассматриваемого номинала (в nominals)
	numOfBonds    int64           // общее количество купюр в текущем вариант размена
}

var interestingNominals = map[int64]interface{}{
	5000: nil,
	500:  nil,
}

// ChangeMoney
//
// params:
// -- cassettes []Cassette -- список касет с купюрами
// -- sum int64 -- сумма для размена
//
// return: размен суммы в map вида: {id: numOfBonds}, где id -- идентификатор кассеты; numOfBonds --
// количество купюр из этой кассеты. Возвращается размен с наиманьшим количеством купюр из найденных
//
// *
func ChangeMoney(cassettes []Cassette, sum int64) map[int]int64 {
	st := &state{
		nominals:      aggregateNominals(cassettes),
		sum:           sum,
		distributions: make([]distribution, 0, 4),
		currD:         make(map[int64]int64, 6),
		currIdx:       0,
		numOfBonds:    0,
	}

	makeChangingGreedily(st)
	if len(st.distributions) <= 0 {
		return map[int]int64{}
	}

	sort.Slice(st.distributions, func(i, j int) bool {
		return st.distributions[i].numOfBonds < st.distributions[j].numOfBonds
	})
	d := st.distributions[0]
	res := make(map[int]int64)
	for _, c := range cassettes {
		if !c.IsIntact() || c.GetNumOfBonds() <= 0 {
			continue
		}

		if need, ok := d.d[c.GetNominal()]; ok && need > 0 {
			res[c.GetId()] = min(need, c.GetNumOfBonds())
			d.d[c.GetNominal()] -= res[c.GetId()]
		}
	}

	return res
}

func aggregateNominals(cassettes []Cassette) [][]int64 {
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

	return getOrdered(nominals)
}

func makeChangingGreedily(st *state) {
	if st.currIdx >= len(st.nominals) {
		if st.sum == 0 {
			st.distributions = append(st.distributions, distribution{
				d:          deepCopy(st.currD),
				remains:    st.sum,
				numOfBonds: st.numOfBonds,
			})
		}
		return
	}

	var nom = st.nominals[st.currIdx][0]
	var numOf = st.nominals[st.currIdx][1]
	var deltaSum, deltaNumOf int64
	if st.sum >= nom {
		var numOfShouldBonds = st.sum / nom
		if numOfShouldBonds <= numOf {
			deltaNumOf = numOfShouldBonds
			deltaSum = numOfShouldBonds * nom
		} else {
			deltaNumOf = numOf
			deltaSum = numOf * nom
		}
	}
	doDistribUpd(st, nom, deltaSum, deltaNumOf)
	makeChangingGreedily(st)
	rollbackDistribUpd(st, nom, deltaSum, deltaNumOf)

	if _, ok := interestingNominals[nom]; ok {
		deltaNumOf--
		deltaSum -= nom
		if st.sum >= nom {
			doDistribUpd(st, nom, deltaSum, deltaNumOf)
			makeChangingGreedily(st)
			rollbackDistribUpd(st, nom, deltaSum, deltaNumOf)
		}
	}
}

func doDistribUpd(st *state, nom, deltaSum, deltaNumOf int64) {
	st.currIdx++
	if deltaSum > 0 && deltaNumOf > 0 {
		st.currD[nom] = deltaNumOf
		st.numOfBonds += deltaNumOf
		st.sum -= deltaSum
	}
}

func rollbackDistribUpd(st *state, nom, deltaSum, deltaNumOf int64) {
	st.currIdx--
	if deltaSum > 0 && deltaNumOf > 0 {
		st.sum += deltaSum
		st.numOfBonds -= deltaNumOf
		delete(st.currD, nom)
	}
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

func deepCopy(m map[int64]int64) map[int64]int64 {
	cp := make(map[int64]int64, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
