package dto

import (
	dmCassette "github.com/CALLlA-74/cashing/internal/domain/Cassette"
	chm "github.com/CALLlA-74/cashing/pkg/changing_money"
)

type ChangeMoneyReq struct {
	Cassettes []Cassette `json:"cassettes" validate:"required"`
	Sum       int64      `json:"sum" validate:"required"`
}

type ChangingResult struct {
	Changing     []Pair `json:"changing" validate:"required"`
	TimeChanging int64  `json:"time" validate:"required"`
}

type Pair struct {
	Id    int   `json:"id" validate:"required"`
	Count int64 `json:"count" validate:"required"`
}

type Cassette struct {
	Id         int   `json:"id" validate:"required"`
	IsIntact   bool  `json:"isIntact" validate:"required"`
	Nominal    int64 `json:"nominal" validate:"required"`
	NumOfBonds int64 `json:"numOfBonds" validate:"required"`
}

var availableNominals = map[int64]interface{}{
	5000: nil,
	2000: nil,
	1000: nil,
	500:  nil,
	200:  nil,
	100:  nil,
}

func ToDomain(req ChangeMoneyReq) ([]chm.Cassette, int64) {
	if req.Sum < 0 {
		return nil, -1
	}

	ids := make(map[int]interface{}, len(req.Cassettes))
	validate := func(c Cassette) bool {
		if _, ok := availableNominals[c.Nominal]; !ok || c.NumOfBonds < 0 {
			return false
		}
		if _, ok := ids[c.Id]; ok {
			return false
		}
		ids[c.Id] = nil
		return true
	}

	res := make([]chm.Cassette, 0, len(req.Cassettes))
	for _, c := range req.Cassettes {
		if !validate(c) {
			return nil, -1
		}
		res = append(res, dmCassette.MakeCassette(c.Id, c.IsIntact, c.Nominal, c.NumOfBonds))
	}
	return res, req.Sum
}
