package dt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	maxSize int
	lru     LRUCache
}

func (ts *TestSuite) SetupSuite() {
	fmt.Println("Setup Suite")
	ts.maxSize = 3
	// ts.lru = NewLRUCache(ts.maxSize) // if you want to tests to use same lru
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (ts *TestSuite) TestAdd() {

	ts.lru = NewLRUCache(ts.maxSize)

	ts.lru.AddEntry("Born Year", 1998)
	val, _ := ts.lru.CheckCache("Born Year")
	ts.Equal(CacheValue(1998), val)
	ts.Equal(1, ts.lru.GetCount())

	_, err := ts.lru.CheckCache("No key")
	ts.EqualError(err, "not found")
}

func (ts *TestSuite) TestAddMultiple() {
	ts.lru = NewLRUCache(ts.maxSize)

	ts.lru.AddEntry("Born Year", 1998)
	val, _ := ts.lru.CheckCache("Born Year")
	ts.Equal(CacheValue(1998), val)
	ts.Equal(1, ts.lru.GetCount())

	_, err := ts.lru.CheckCache("No key")
	ts.EqualError(err, "not found")

	ts.lru.AddEntry("Favorite Number", 52)
	val, _ = ts.lru.CheckCache("Favorite Number")
	ts.Equal(CacheValue(52), val)
	ts.Equal(2, ts.lru.GetCount())
}

func (ts *TestSuite) TestRemoveLeastRecentlyUsedEntry() {
	ts.lru.AddEntry("Key1", 1)
	ts.lru.AddEntry("Key2", 2)
	ts.lru.AddEntry("Key3", 3)

	_, _ = ts.lru.CheckCache("Key1") // move to head

	ts.lru.AddEntry("Key4", 4) // remove tail, 2

	val, err := ts.lru.CheckCache("Key2")
	ts.Equal(CacheValue(0), val)
	ts.EqualError(err, "not found")
	ts.Equal(3, ts.lru.GetCount())

}
