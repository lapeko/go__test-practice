package mail

import (
	"fmt"
	"io"
	"sync"
	"text/tabwriter"
)

type MailService interface {
	Send(to, subject, body string) error
}

type mailService struct {
	writer io.Writer
}

func New(writer io.Writer) MailService {
	return &mailService{writer}
}

func (ms *mailService) Send(to, subject, body string) error {
	tw := tabwriter.NewWriter(ms.writer, 0, 0, 4, ' ', 0)
	if _, err := tw.Write([]byte("to\tsubject\tbody\n")); err != nil {
		return err
	}
	if _, err := tw.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", to, subject, body))); err != nil {
		return err
	}
	return tw.Flush()
}

func EmailBatchSend(ms MailService, subject, body string, to []string) (notSent []string, err error) {
	var wg sync.WaitGroup
	mutex := sync.Mutex{}
	wg.Add(len(to))

	for _, receiver := range to {
		go func() {
			defer wg.Done()
			if e := ms.Send(receiver, subject, body); e != nil {
				mutex.Lock()
				if err == nil {
					err = fmt.Errorf("EmailBatchSend error: %w", e)
				}
				notSent = append(notSent, receiver)
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	return
}
