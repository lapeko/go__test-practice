package cache

import (
	"fmt"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand/v2"
	"sync"
	"testing"
)

type cacheMock struct {
	suite.Suite
	cache Cache[any]
	wg    sync.WaitGroup
}

func (suite *cacheMock) SetupSuite() {
	suite.cache = New[any]()
}

func (suite *cacheMock) TearDownTest() {
	suite.cache.Clear()
}

func (suite *cacheMock) stringifyInt(d int) string {
	return fmt.Sprintf("%d", d)
}

//func (suite *cacheMock) SetupTest() {}
//func (suite *cacheMock) TearDownSuite() {}

func (suite *cacheMock) TestWriteReadOnce() {
	key := "something"
	randInt := rand.IntN(1000)
	_, ok := suite.cache.Get(key)
	suite.False(ok, fmt.Sprintf("key: %s should not exist in clean/new cache", key))
	suite.cache.Put(key, randInt)
	got, ok := suite.cache.Get(key)
	suite.True(ok, fmt.Sprintf("key: %s should exist after inserting it into cache", key))
	suite.Equal(got, randInt, fmt.Sprintf("value: %d should be equal to: %d", got, randInt))
}

func (suite *cacheMock) TestConcurrentOps() {
	size := 100
	suite.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer suite.wg.Done()
			suite.cache.Put(suite.stringifyInt(i), i)
		}()
	}
	suite.wg.Wait()

	suite.wg.Add(size)
	for i := 0; i < size; i++ {
		key := suite.stringifyInt(i)
		go func() {
			defer suite.wg.Done()
			got, ok := suite.cache.Get(key)
			suite.True(ok, fmt.Sprintf("key: %s should exist after inserting it into cache", key))
			suite.Equal(got, i, fmt.Sprintf("value: %d should be equal to: %d", got, i))
		}()
	}
	suite.wg.Wait()

	suite.wg.Add(size)
	for i := 0; i < size; i++ {
		key := suite.stringifyInt(i)
		go func() {
			defer suite.wg.Done()
			suite.cache.Delete(key)
			got, ok := suite.cache.Get(key)
			suite.False(ok)
			suite.Nil(got)
		}()
	}
	suite.wg.Wait()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(cacheMock))
}
