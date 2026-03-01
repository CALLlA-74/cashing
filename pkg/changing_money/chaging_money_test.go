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
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
}

func TestChangeMoney2(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 10, true),
		MakeCassette(10, 2000, 10, true),
		MakeCassette(1555, 100, 11, true),
	}
	sum := int64(7300)
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
}

func TestChangeMoney3(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 1, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(1555, 100, 1, true),
	}
	sum := int64(5100)
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
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

func TestChangeMoney5(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 20, true),
		MakeCassette(10, 2000, 5, true),
	}
	sum := int64(101000)
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
}

func TestChangeMoney6(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 18, true),
		MakeCassette(10, 2000, 10, true),
	}
	sum := int64(101000)
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
}

func TestChangeMoney7(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 2, true),
		MakeCassette(10, 2000, 10, true),
		MakeCassette(90, 500, 10, true),
		MakeCassette(80, 200, 10, true),
	}
	sum := int64(7100)
	res := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, res)
}

func TestChangeMoney8(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 2, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(90, 1000, 1, true),
	}
	sum := int64(9000)
	res := ChangeMoney(inp, sum)
	assert.Equal(t, 0, len(res))
}

func TestChangeMoney9(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 5000, 1, true),
		MakeCassette(10, 2000, 1, true),
		MakeCassette(1555, 100, 1, true),
		MakeCassette(1556, 10, 2, true),
		MakeCassette(1557, 5, 2, true),
		MakeCassette(1558, 1, 3, true),
	}
	sum := int64(5126)
	ans := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, ans)
}

func TestChangeMoney10(t *testing.T) {
	var inp []Cassette
	sum := int64(0)
	ans := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, ans)
}

func TestChangeMoney11(t *testing.T) {
	var inp []Cassette
	sum := int64(100)
	ans := ChangeMoney(inp, sum)
	assert.Equal(t, 0, len(ans))
}

func TestChangeMoney12(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 100, 3, true),
		MakeCassette(10, 2000, 4, true),
		MakeCassette(1555, 200, 0, false),
		MakeCassette(1556, 1000, 4, true),
		MakeCassette(1557, 500, 4, true),
	}
	sum := int64(1200)
	ans := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, ans)
}

func TestChangeMoney13(t *testing.T) {
	inp := []Cassette{
		MakeCassette(11, 100, 3, false),
		MakeCassette(10, 5000, 10000000000, true),
		MakeCassette(1555, 200, 4, false),
		MakeCassette(1556, 1000, 4, false),

		MakeCassette(1557, 1000, 1000000000, true),
		MakeCassette(1558, 100, 1000000000, true),
		MakeCassette(1559, 500, 1000000000, false),
		MakeCassette(15510, 100, 4, false),
	}
	sum := int64(9999999999000)
	ans := ChangeMoney(inp, sum)
	checkAns(t, inp, sum, ans)
}

func BenchmarkChangeMoney(b *testing.B) {
	inp := []Cassette{
		MakeCassette(11, 100, 3, false),
		MakeCassette(10, 5000, 1000000000, true),
		MakeCassette(1555, 200, 4, false),
		MakeCassette(1556, 1000, 4, false),

		MakeCassette(1557, 1000, 1000000000, true),
		MakeCassette(1558, 100, 1000000000, true),
		MakeCassette(1559, 500, 1000000000, false),
		MakeCassette(15510, 100, 4, false),
	}
	sum := int64(999999999900)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChangeMoney(inp, sum)
	}
}
