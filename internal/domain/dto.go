package domain

type ChangeMoneyReq struct {
	Cassettes []Cassette `json:"cassettes" validate:"required"`
	Sum       int64      `json:"sum" validate:"required"`
}

type ChangingResult struct {
	Changing     []Pair `json:"changing" validate:"required"`
	TimeChanging int64  `json:"time" validate:"required"`
	IsFound      bool   `json:"isFound" validate:"required"`
}

type Pair struct {
	Id    int   `json:"id" validate:"required"`
	Count int64 `json:"count" validate:"required"`
}
