package notifier

import "sync"

type Notifier interface {
	Notify(userId, message string) error
}

type MultiNotifier interface {
	Add(n Notifier)
	NotifyAll(userId, message string) <-chan error
}

type multiNotifier struct {
	sync.Mutex
	notifiers []Notifier
}

func New() MultiNotifier {
	return &multiNotifier{}
}

func (mn *multiNotifier) Add(n Notifier) {
	mn.Lock()
	mn.notifiers = append(mn.notifiers, n)
	mn.Unlock()
}

func (mn *multiNotifier) NotifyAll(userId, message string) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)
		var wg sync.WaitGroup
		wg.Add(len(mn.notifiers))

		for _, notifier := range mn.notifiers {
			go func() {
				defer wg.Done()
				if err := notifier.Notify(userId, message); err != nil {
					errChan <- err
				}
			}()
		}
		wg.Wait()
	}()

	return errChan
}
