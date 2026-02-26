package changing_money

import "sort"

type Cassette interface {
	GetId() int
	IsIntact() bool       // флаг исправности кассеты
	GetNominal() int64    // номинал купюры
	GetNumOfBonds() int64 // количество купюр в кассете
}

// state
//
//	хранит состояние поиска
//
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
	cassettes      []Cassette
	resultFound    bool
	sum            int64         // текущий остаток размениваемой суммы
	currIdx        int           // индекс текущей рассматриваемой кассеты
	prevIdx        int           // индекс предыдущей рассматриваемой кассеты
	prevNominalIdx int           // индекс предыдущей кассеты, из которой были взяты купюры для размена (с номиналом больше, чем в текущей кассете)
	res            map[int]int64 // результат поиска
}

// ChangeMoney
//
//	cassettes []Cassette -- список касет с купюрами
//	sum int64 -- сумма для размена
//
//	Порядок поиска:
//	1. сортировка по убыванию номинала купюр
//	2. (см. findChanging())
//
// *
func ChangeMoney(cassettes []Cassette, sum int64) map[int]int64 {
	sort.Slice(cassettes, func(i, j int) bool {
		return cassettes[i].GetNominal() > cassettes[j].GetNominal()
	})

	st := &state{
		cassettes:      cassettes,
		resultFound:    false,
		sum:            sum,
		currIdx:        0,
		prevIdx:        -1,
		prevNominalIdx: -1,
		res:            make(map[int]int64),
	}
	findChanging(st)

	if st.resultFound {
		return st.res
	}
	return map[int]int64{}
}

// findChanging
//
//	Поиск размена в невозрастающей последовательности кассет.
//
//	порядок поиска:
//	1. для каждой кассеты:
//	1.1. если сумма st.sum больше номинала текущей купюры,
//		то пытаемся взять макс. количество купюр до превышения суммы sum
//	1.2. иначе если выполняется условие возврата (см. isNeedBackPrevBond()),
//		то запускаем альтернативный путь поиска размена:
//	1.2.1	прибавляем к сумме st.sum одну купюру предыдущего номинала,
//		вычтенного у нее (находим ее по st.prevNominalIdx),
//	1.2.2	и делаем новый вызов findChanging, выполняя поиск с текущей кассеты (с позиции st.currIdx)
//	1.2.3 если поиск завершился неудачно, то снова берем купюру из предыдущей касеты (по st.prevNominalIdx)
//		и продолжаем поиск далее
//
// *
func findChanging(st *state) {
	var (
		c, prevC         Cassette
		numOfShouldBonds int64
		tempSum          int64
		tempPrevIdx      int
	)
	for i := st.currIdx; i < len(st.cassettes); i++ {
		c = st.cassettes[i]
		st.currIdx = i
		if st.prevIdx != -1 && st.cassettes[st.prevIdx].GetNominal() > c.GetNominal() {
			st.prevNominalIdx = st.prevIdx
		}

		if !c.IsIntact() || c.GetNumOfBonds() <= 0 {
			continue
		}

		if st.sum >= c.GetNominal() {
			numOfShouldBonds = st.sum / c.GetNominal()
			if numOfShouldBonds <= c.GetNumOfBonds() {
				st.res[c.GetId()] = numOfShouldBonds
				st.sum -= numOfShouldBonds * c.GetNominal()
				st.prevIdx = i
			}
		} else if isNeedBackPrevBond(st) {
			prevC = st.cassettes[st.prevNominalIdx]

			st.res[prevC.GetId()]--
			tempSum = st.sum
			st.sum += prevC.GetNominal()
			tempPrevIdx = st.prevNominalIdx
			st.prevNominalIdx = -1
			findChanging(st)

			if st.resultFound {
				return
			}
			st.prevNominalIdx = tempPrevIdx
			st.res[prevC.GetId()]++
			st.sum = tempSum
		}
	}

	if st.sum == 0 {
		st.resultFound = true
	}
}

func isNeedBackPrevBond(st *state) bool {
	if st.prevNominalIdx == -1 {
		return false
	}

	c := st.cassettes[st.currIdx]
	prevC := st.cassettes[st.prevNominalIdx]
	return prevC.GetNominal()%c.GetNominal() != 0
}
