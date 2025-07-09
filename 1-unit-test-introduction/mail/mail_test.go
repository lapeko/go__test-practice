package mail

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
)

type writerMock struct {
	bytes []byte
}

func (w *writerMock) Write(p []byte) (n int, err error) {
	w.bytes = append(w.bytes, p...)
	return len(p), nil
}

func TestMailService_Send(t *testing.T) {
	want := "to      subject    body\nName    Subject    Long Body\n"

	wm := &writerMock{}
	ms := New(wm)
	if err := ms.Send("Name", "Subject", "Long Body"); err != nil {
		t.Errorf("MailService.Send error\n\tgot error: %q\t\nwant: nil", err.Error())
	}

	got := string(wm.bytes)

	if got != want {
		t.Errorf("MailService.Send\n\twrote:\t%q\n\twant:\t%q", got, want)
	}
}

type mailServiceMock struct {
	sync.Mutex
	callCounter atomic.Int64
	sent        []string
	genError    func(mock *mailServiceMock) error
}

func (ms *mailServiceMock) Send(to, _, _ string) error {
	ms.callCounter.Add(1)
	if err := ms.genError(ms); err != nil {
		return err
	}
	ms.Lock()
	ms.sent = append(ms.sent, to)
	ms.Unlock()
	return nil
}

func TestEmailBatchSend(t *testing.T) {
	mockError := errors.New("nope")
	errorWant := fmt.Errorf("EmailBatchSend error: %w", mockError)

	msm := &mailServiceMock{genError: func(m *mailServiceMock) error {
		if m.callCounter.Load() == 2 {
			return mockError
		}
		return nil
	}}

	notSentGot, gotError := EmailBatchSend(msm, "", "", []string{"client1", "client2", "client3"})

	if !reflect.DeepEqual(gotError, errorWant) {
		t.Errorf("TestEmailBatchSend error\n\tgot:\t%q\n\twant:\t%q", gotError, errorWant)
	}

	if len(notSentGot) != 1 {
		t.Errorf("TestEmailBatchSend notSent\n\tgot:\t%v\n\twant:\t%v", len(notSentGot), 1)
	}
}
