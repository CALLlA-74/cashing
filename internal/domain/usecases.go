package domain

import (
	changingMoney "github.com/CALLlA-74/cashing/pkg/changing_money"
	"time"
)

type ChangingMoneyUC struct{}

func NewUC() *ChangingMoneyUC {
	return &ChangingMoneyUC{}
}

func (uc *ChangingMoneyUC) ChangeMoney(req ChangeMoneyReq) (ChangingResult, error) {
	cass := uc.mapCassettes(req.Cassettes)
	timeStart := time.Time{}.UnixMilli()
	result, isFound := changingMoney.ChangeMoney(cass, req.Sum)
	timeMillis := time.Now().UnixMilli() - timeStart
	return ChangingResult{
		Changing:     uc.mapResult(result),
		TimeChanging: timeMillis,
		IsFound:      isFound,
	}, nil
}

func (uc *ChangingMoneyUC) mapResult(mp map[int]int64) []Pair {
	ans := make([]Pair, 0, len(mp))
	for k, v := range mp {
		ans = append(ans, Pair{k, v})
	}
	return ans
}

func (uc *ChangingMoneyUC) mapCassettes(cass []Cassette) []changingMoney.Cassette {
	res := make([]changingMoney.Cassette, 0, len(cass))
	for _, c := range cass {
		res = append(res, c)
	}
	return res
}
