package domain

import (
	"errors"
	"github.com/CALLlA-74/cashing/internal/domain/Cassette/dto"
	changingMoney "github.com/CALLlA-74/cashing/pkg/changing_money"
	"time"
)

type ChangingMoneyUC struct{}

func NewUC() *ChangingMoneyUC {
	return &ChangingMoneyUC{}
}

func (uc *ChangingMoneyUC) ChangeMoney(req dto.ChangeMoneyReq) (dto.ChangingResult, error) {
	cass, sum := dto.ToDomain(req)
	if cass == nil || sum == -1 {
		return dto.ChangingResult{}, errors.New("некорректные значения")
	}

	timeStart := time.Now().UnixMilli()
	result := changingMoney.ChangeMoney(cass, sum)
	timeMillis := time.Now().UnixMilli() - timeStart
	return dto.ChangingResult{
		Changing:     uc.mapResult(result),
		TimeChanging: timeMillis,
	}, nil
}

func (uc *ChangingMoneyUC) mapResult(mp map[int]int64) []dto.Pair {
	ans := make([]dto.Pair, 0, len(mp))
	for k, v := range mp {
		ans = append(ans, dto.Pair{k, v})
	}
	return ans
}
