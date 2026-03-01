package Cassette

type Cassette struct {
	id         int
	isIntact   bool
	nominal    int64
	numOfBonds int64
}

func MakeCassette(id int, isIntact bool, nominal, numOfBonds int64) Cassette {
	return Cassette{
		id:         id,
		isIntact:   isIntact,
		nominal:    nominal,
		numOfBonds: numOfBonds,
	}
}

func (c Cassette) GetId() int {
	return c.id
}

func (c Cassette) IsIntact() bool {
	return c.isIntact
}

func (c Cassette) GetNominal() int64 {
	return c.nominal
}

func (c Cassette) GetNumOfBonds() int64 {
	return c.numOfBonds
}
