package changing_money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type BaseCassette struct {
	id         int
	isIntact   bool
	nominal    int64
	numOfBonds int64
}

func (bs BaseCassette) GetId() int {
	return bs.id
}

func (bs BaseCassette) IsIntact() bool {
	return bs.isIntact
}

func (bs BaseCassette) GetNominal() int64 {
	return bs.nominal
}

func (bs BaseCassette) GetNumOfBonds() int64 {
	return bs.numOfBonds
}

func MakeCassette(id int, nominal, numOfBonds int64, isIntact bool) Cassette {
	return BaseCassette{
		id:         id,
		isIntact:   isIntact,
		nominal:    nominal,
		numOfBonds: numOfBonds,
	}
}

func checkAns(t *testing.T, inp []Cassette, sum int64, ans map[int]int64) {
	for _, c := range inp {
		if v, ok := ans[c.GetId()]; ok {
			sum -= c.GetNominal() * v
		}
	}
	assert.Equal(t, int64(0), sum)
}

func TestChangeMoney(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 10, true),
		MakeCassette(10, 2000, 10, true),
	}
	sum := int64(26000)
	checkAns(t, inp, sum, ChangeMoney(inp, sum))
}

func TestChangeMoney2(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 10, true),
		MakeCassette(10, 2000, 10, true),
		MakeCassette(1555, 100, 11, true),
	}
	sum := int64(7300)
	checkAns(t, inp, sum, ChangeMoney(inp, sum))
}

func TestChangeMoney3(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 1, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(1555, 100, 1, true),
	}
	sum := int64(5100)
	checkAns(t, inp, sum, ChangeMoney(inp, sum))
}

func TestChangeMoney4(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 1, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(1555, 100, 1, true),
	}
	sum := int64(5126)
	ans := ChangeMoney(inp, sum)
	assert.Equal(t, 0, len(ans))
}

func BenchmarkChangeMoney(b *testing.B) {
	inp := []Cassette{
		MakeCassette(11, 5000, 1, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(1555, 100, 1, true),
	}
	sum := int64(1e18)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChangeMoney(inp, sum)
	}
}
