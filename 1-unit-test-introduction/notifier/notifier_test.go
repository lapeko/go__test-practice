package notifier

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

type mockNotifier struct {
	mock.Mock
}

func (m *mockNotifier) Notify(userId, message string) error {
	args := m.Called(userId, message)
	return args.Error(0)
}

func TestMultiNotifier(t *testing.T) {
	mn := New()
	notifier1 := new(mockNotifier)
	notifier2 := new(mockNotifier)

	mn.Add(notifier1)
	mn.Add(notifier2)

	userId := fmt.Sprintf("%d", rand.IntN(100))
	body := "message"
	err := errors.New("nope")

	notifier1.On("Notify", userId, body).Return(nil).Once()
	notifier2.On("Notify", userId, body).Return(err).Once()

	errCh := mn.NotifyAll(userId, "message")

	var fails []error
	for err := range errCh {
		fails = append(fails, err)
	}

	require.Equal(t, len(fails), 1)
	require.Equal(t, fails[0], err)

	notifier1.AssertExpectations(t)
	notifier2.AssertExpectations(t)
}
