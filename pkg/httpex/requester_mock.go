package httpex

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i gitlab.ozon.dev/capcom6/homework-2/pkg/httpex.Requester -o ./pkg/httpex/requester_mock.go -n RequesterMock

import (
	"io"
	"net/http"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RequesterMock implements Requester
type RequesterMock struct {
	t minimock.Tester

	funcPost          func(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	inspectFuncPost   func(url string, contentType string, body io.Reader)
	afterPostCounter  uint64
	beforePostCounter uint64
	PostMock          mRequesterMockPost
}

// NewRequesterMock returns a mock for Requester
func NewRequesterMock(t minimock.Tester) *RequesterMock {
	m := &RequesterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.PostMock = mRequesterMockPost{mock: m}
	m.PostMock.callArgs = []*RequesterMockPostParams{}

	return m
}

type mRequesterMockPost struct {
	mock               *RequesterMock
	defaultExpectation *RequesterMockPostExpectation
	expectations       []*RequesterMockPostExpectation

	callArgs []*RequesterMockPostParams
	mutex    sync.RWMutex
}

// RequesterMockPostExpectation specifies expectation struct of the Requester.Post
type RequesterMockPostExpectation struct {
	mock    *RequesterMock
	params  *RequesterMockPostParams
	results *RequesterMockPostResults
	Counter uint64
}

// RequesterMockPostParams contains parameters of the Requester.Post
type RequesterMockPostParams struct {
	url         string
	contentType string
	body        io.Reader
}

// RequesterMockPostResults contains results of the Requester.Post
type RequesterMockPostResults struct {
	resp *http.Response
	err  error
}

// Expect sets up expected params for Requester.Post
func (mmPost *mRequesterMockPost) Expect(url string, contentType string, body io.Reader) *mRequesterMockPost {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("RequesterMock.Post mock is already set by Set")
	}

	if mmPost.defaultExpectation == nil {
		mmPost.defaultExpectation = &RequesterMockPostExpectation{}
	}

	mmPost.defaultExpectation.params = &RequesterMockPostParams{url, contentType, body}
	for _, e := range mmPost.expectations {
		if minimock.Equal(e.params, mmPost.defaultExpectation.params) {
			mmPost.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPost.defaultExpectation.params)
		}
	}

	return mmPost
}

// Inspect accepts an inspector function that has same arguments as the Requester.Post
func (mmPost *mRequesterMockPost) Inspect(f func(url string, contentType string, body io.Reader)) *mRequesterMockPost {
	if mmPost.mock.inspectFuncPost != nil {
		mmPost.mock.t.Fatalf("Inspect function is already set for RequesterMock.Post")
	}

	mmPost.mock.inspectFuncPost = f

	return mmPost
}

// Return sets up results that will be returned by Requester.Post
func (mmPost *mRequesterMockPost) Return(resp *http.Response, err error) *RequesterMock {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("RequesterMock.Post mock is already set by Set")
	}

	if mmPost.defaultExpectation == nil {
		mmPost.defaultExpectation = &RequesterMockPostExpectation{mock: mmPost.mock}
	}
	mmPost.defaultExpectation.results = &RequesterMockPostResults{resp, err}
	return mmPost.mock
}

//Set uses given function f to mock the Requester.Post method
func (mmPost *mRequesterMockPost) Set(f func(url string, contentType string, body io.Reader) (resp *http.Response, err error)) *RequesterMock {
	if mmPost.defaultExpectation != nil {
		mmPost.mock.t.Fatalf("Default expectation is already set for the Requester.Post method")
	}

	if len(mmPost.expectations) > 0 {
		mmPost.mock.t.Fatalf("Some expectations are already set for the Requester.Post method")
	}

	mmPost.mock.funcPost = f
	return mmPost.mock
}

// When sets expectation for the Requester.Post which will trigger the result defined by the following
// Then helper
func (mmPost *mRequesterMockPost) When(url string, contentType string, body io.Reader) *RequesterMockPostExpectation {
	if mmPost.mock.funcPost != nil {
		mmPost.mock.t.Fatalf("RequesterMock.Post mock is already set by Set")
	}

	expectation := &RequesterMockPostExpectation{
		mock:   mmPost.mock,
		params: &RequesterMockPostParams{url, contentType, body},
	}
	mmPost.expectations = append(mmPost.expectations, expectation)
	return expectation
}

// Then sets up Requester.Post return parameters for the expectation previously defined by the When method
func (e *RequesterMockPostExpectation) Then(resp *http.Response, err error) *RequesterMock {
	e.results = &RequesterMockPostResults{resp, err}
	return e.mock
}

// Post implements Requester
func (mmPost *RequesterMock) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	mm_atomic.AddUint64(&mmPost.beforePostCounter, 1)
	defer mm_atomic.AddUint64(&mmPost.afterPostCounter, 1)

	if mmPost.inspectFuncPost != nil {
		mmPost.inspectFuncPost(url, contentType, body)
	}

	mm_params := &RequesterMockPostParams{url, contentType, body}

	// Record call args
	mmPost.PostMock.mutex.Lock()
	mmPost.PostMock.callArgs = append(mmPost.PostMock.callArgs, mm_params)
	mmPost.PostMock.mutex.Unlock()

	for _, e := range mmPost.PostMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.resp, e.results.err
		}
	}

	if mmPost.PostMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPost.PostMock.defaultExpectation.Counter, 1)
		mm_want := mmPost.PostMock.defaultExpectation.params
		mm_got := RequesterMockPostParams{url, contentType, body}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPost.t.Errorf("RequesterMock.Post got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPost.PostMock.defaultExpectation.results
		if mm_results == nil {
			mmPost.t.Fatal("No results are set for the RequesterMock.Post")
		}
		return (*mm_results).resp, (*mm_results).err
	}
	if mmPost.funcPost != nil {
		return mmPost.funcPost(url, contentType, body)
	}
	mmPost.t.Fatalf("Unexpected call to RequesterMock.Post. %v %v %v", url, contentType, body)
	return
}

// PostAfterCounter returns a count of finished RequesterMock.Post invocations
func (mmPost *RequesterMock) PostAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPost.afterPostCounter)
}

// PostBeforeCounter returns a count of RequesterMock.Post invocations
func (mmPost *RequesterMock) PostBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPost.beforePostCounter)
}

// Calls returns a list of arguments used in each call to RequesterMock.Post.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPost *mRequesterMockPost) Calls() []*RequesterMockPostParams {
	mmPost.mutex.RLock()

	argCopy := make([]*RequesterMockPostParams, len(mmPost.callArgs))
	copy(argCopy, mmPost.callArgs)

	mmPost.mutex.RUnlock()

	return argCopy
}

// MinimockPostDone returns true if the count of the Post invocations corresponds
// the number of defined expectations
func (m *RequesterMock) MinimockPostDone() bool {
	for _, e := range m.PostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPost != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		return false
	}
	return true
}

// MinimockPostInspect logs each unmet expectation
func (m *RequesterMock) MinimockPostInspect() {
	for _, e := range m.PostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequesterMock.Post with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		if m.PostMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequesterMock.Post")
		} else {
			m.t.Errorf("Expected call to RequesterMock.Post with params: %#v", *m.PostMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPost != nil && mm_atomic.LoadUint64(&m.afterPostCounter) < 1 {
		m.t.Error("Expected call to RequesterMock.Post")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RequesterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockPostInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RequesterMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RequesterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockPostDone()
}
