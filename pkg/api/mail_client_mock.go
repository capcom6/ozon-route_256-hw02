package api

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i gitlab.ozon.dev/capcom6/homework-2/pkg/api.MailAggregatorClient -o ./pkg/api/mail_client_mock.go -n MailAggregatorClientMock

import (
	context "context"
	sync "sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	grpc "google.golang.org/grpc"
)

// MailAggregatorClientMock implements MailAggregatorClient
type MailAggregatorClientMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption) (ep1 *Empty, err error)
	inspectFuncCreate   func(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mMailAggregatorClientMockCreate

	funcDelete          func(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption) (mp1 *Mailboxes, err error)
	inspectFuncDelete   func(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mMailAggregatorClientMockDelete

	funcPull          func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Messages, err error)
	inspectFuncPull   func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption)
	afterPullCounter  uint64
	beforePullCounter uint64
	PullMock          mMailAggregatorClientMockPull

	funcSelect          func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Mailboxes, err error)
	inspectFuncSelect   func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption)
	afterSelectCounter  uint64
	beforeSelectCounter uint64
	SelectMock          mMailAggregatorClientMockSelect
}

// NewMailAggregatorClientMock returns a mock for MailAggregatorClient
func NewMailAggregatorClientMock(t minimock.Tester) *MailAggregatorClientMock {
	m := &MailAggregatorClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mMailAggregatorClientMockCreate{mock: m}
	m.CreateMock.callArgs = []*MailAggregatorClientMockCreateParams{}

	m.DeleteMock = mMailAggregatorClientMockDelete{mock: m}
	m.DeleteMock.callArgs = []*MailAggregatorClientMockDeleteParams{}

	m.PullMock = mMailAggregatorClientMockPull{mock: m}
	m.PullMock.callArgs = []*MailAggregatorClientMockPullParams{}

	m.SelectMock = mMailAggregatorClientMockSelect{mock: m}
	m.SelectMock.callArgs = []*MailAggregatorClientMockSelectParams{}

	return m
}

type mMailAggregatorClientMockCreate struct {
	mock               *MailAggregatorClientMock
	defaultExpectation *MailAggregatorClientMockCreateExpectation
	expectations       []*MailAggregatorClientMockCreateExpectation

	callArgs []*MailAggregatorClientMockCreateParams
	mutex    sync.RWMutex
}

// MailAggregatorClientMockCreateExpectation specifies expectation struct of the MailAggregatorClient.Create
type MailAggregatorClientMockCreateExpectation struct {
	mock    *MailAggregatorClientMock
	params  *MailAggregatorClientMockCreateParams
	results *MailAggregatorClientMockCreateResults
	Counter uint64
}

// MailAggregatorClientMockCreateParams contains parameters of the MailAggregatorClient.Create
type MailAggregatorClientMockCreateParams struct {
	ctx  context.Context
	in   *MailboxCreate
	opts []grpc.CallOption
}

// MailAggregatorClientMockCreateResults contains results of the MailAggregatorClient.Create
type MailAggregatorClientMockCreateResults struct {
	ep1 *Empty
	err error
}

// Expect sets up expected params for MailAggregatorClient.Create
func (mmCreate *mMailAggregatorClientMockCreate) Expect(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption) *mMailAggregatorClientMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MailAggregatorClientMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MailAggregatorClientMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &MailAggregatorClientMockCreateParams{ctx, in, opts}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the MailAggregatorClient.Create
func (mmCreate *mMailAggregatorClientMockCreate) Inspect(f func(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption)) *mMailAggregatorClientMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for MailAggregatorClientMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by MailAggregatorClient.Create
func (mmCreate *mMailAggregatorClientMockCreate) Return(ep1 *Empty, err error) *MailAggregatorClientMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MailAggregatorClientMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MailAggregatorClientMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &MailAggregatorClientMockCreateResults{ep1, err}
	return mmCreate.mock
}

//Set uses given function f to mock the MailAggregatorClient.Create method
func (mmCreate *mMailAggregatorClientMockCreate) Set(f func(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption) (ep1 *Empty, err error)) *MailAggregatorClientMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the MailAggregatorClient.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the MailAggregatorClient.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the MailAggregatorClient.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mMailAggregatorClientMockCreate) When(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption) *MailAggregatorClientMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MailAggregatorClientMock.Create mock is already set by Set")
	}

	expectation := &MailAggregatorClientMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &MailAggregatorClientMockCreateParams{ctx, in, opts},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up MailAggregatorClient.Create return parameters for the expectation previously defined by the When method
func (e *MailAggregatorClientMockCreateExpectation) Then(ep1 *Empty, err error) *MailAggregatorClientMock {
	e.results = &MailAggregatorClientMockCreateResults{ep1, err}
	return e.mock
}

// Create implements MailAggregatorClient
func (mmCreate *MailAggregatorClientMock) Create(ctx context.Context, in *MailboxCreate, opts ...grpc.CallOption) (ep1 *Empty, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, in, opts...)
	}

	mm_params := &MailAggregatorClientMockCreateParams{ctx, in, opts}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ep1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := MailAggregatorClientMockCreateParams{ctx, in, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("MailAggregatorClientMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the MailAggregatorClientMock.Create")
		}
		return (*mm_results).ep1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, in, opts...)
	}
	mmCreate.t.Fatalf("Unexpected call to MailAggregatorClientMock.Create. %v %v %v", ctx, in, opts)
	return
}

// CreateAfterCounter returns a count of finished MailAggregatorClientMock.Create invocations
func (mmCreate *MailAggregatorClientMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of MailAggregatorClientMock.Create invocations
func (mmCreate *MailAggregatorClientMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to MailAggregatorClientMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mMailAggregatorClientMockCreate) Calls() []*MailAggregatorClientMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*MailAggregatorClientMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *MailAggregatorClientMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *MailAggregatorClientMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MailAggregatorClientMock.Create")
		} else {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to MailAggregatorClientMock.Create")
	}
}

type mMailAggregatorClientMockDelete struct {
	mock               *MailAggregatorClientMock
	defaultExpectation *MailAggregatorClientMockDeleteExpectation
	expectations       []*MailAggregatorClientMockDeleteExpectation

	callArgs []*MailAggregatorClientMockDeleteParams
	mutex    sync.RWMutex
}

// MailAggregatorClientMockDeleteExpectation specifies expectation struct of the MailAggregatorClient.Delete
type MailAggregatorClientMockDeleteExpectation struct {
	mock    *MailAggregatorClientMock
	params  *MailAggregatorClientMockDeleteParams
	results *MailAggregatorClientMockDeleteResults
	Counter uint64
}

// MailAggregatorClientMockDeleteParams contains parameters of the MailAggregatorClient.Delete
type MailAggregatorClientMockDeleteParams struct {
	ctx  context.Context
	in   *MailboxDelete
	opts []grpc.CallOption
}

// MailAggregatorClientMockDeleteResults contains results of the MailAggregatorClient.Delete
type MailAggregatorClientMockDeleteResults struct {
	mp1 *Mailboxes
	err error
}

// Expect sets up expected params for MailAggregatorClient.Delete
func (mmDelete *mMailAggregatorClientMockDelete) Expect(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption) *mMailAggregatorClientMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MailAggregatorClientMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &MailAggregatorClientMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &MailAggregatorClientMockDeleteParams{ctx, in, opts}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the MailAggregatorClient.Delete
func (mmDelete *mMailAggregatorClientMockDelete) Inspect(f func(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption)) *mMailAggregatorClientMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for MailAggregatorClientMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by MailAggregatorClient.Delete
func (mmDelete *mMailAggregatorClientMockDelete) Return(mp1 *Mailboxes, err error) *MailAggregatorClientMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MailAggregatorClientMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &MailAggregatorClientMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &MailAggregatorClientMockDeleteResults{mp1, err}
	return mmDelete.mock
}

//Set uses given function f to mock the MailAggregatorClient.Delete method
func (mmDelete *mMailAggregatorClientMockDelete) Set(f func(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption) (mp1 *Mailboxes, err error)) *MailAggregatorClientMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the MailAggregatorClient.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the MailAggregatorClient.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the MailAggregatorClient.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mMailAggregatorClientMockDelete) When(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption) *MailAggregatorClientMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("MailAggregatorClientMock.Delete mock is already set by Set")
	}

	expectation := &MailAggregatorClientMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &MailAggregatorClientMockDeleteParams{ctx, in, opts},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up MailAggregatorClient.Delete return parameters for the expectation previously defined by the When method
func (e *MailAggregatorClientMockDeleteExpectation) Then(mp1 *Mailboxes, err error) *MailAggregatorClientMock {
	e.results = &MailAggregatorClientMockDeleteResults{mp1, err}
	return e.mock
}

// Delete implements MailAggregatorClient
func (mmDelete *MailAggregatorClientMock) Delete(ctx context.Context, in *MailboxDelete, opts ...grpc.CallOption) (mp1 *Mailboxes, err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, in, opts...)
	}

	mm_params := &MailAggregatorClientMockDeleteParams{ctx, in, opts}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.mp1, e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := MailAggregatorClientMockDeleteParams{ctx, in, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("MailAggregatorClientMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the MailAggregatorClientMock.Delete")
		}
		return (*mm_results).mp1, (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, in, opts...)
	}
	mmDelete.t.Fatalf("Unexpected call to MailAggregatorClientMock.Delete. %v %v %v", ctx, in, opts)
	return
}

// DeleteAfterCounter returns a count of finished MailAggregatorClientMock.Delete invocations
func (mmDelete *MailAggregatorClientMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of MailAggregatorClientMock.Delete invocations
func (mmDelete *MailAggregatorClientMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to MailAggregatorClientMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mMailAggregatorClientMockDelete) Calls() []*MailAggregatorClientMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*MailAggregatorClientMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *MailAggregatorClientMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *MailAggregatorClientMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MailAggregatorClientMock.Delete")
		} else {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to MailAggregatorClientMock.Delete")
	}
}

type mMailAggregatorClientMockPull struct {
	mock               *MailAggregatorClientMock
	defaultExpectation *MailAggregatorClientMockPullExpectation
	expectations       []*MailAggregatorClientMockPullExpectation

	callArgs []*MailAggregatorClientMockPullParams
	mutex    sync.RWMutex
}

// MailAggregatorClientMockPullExpectation specifies expectation struct of the MailAggregatorClient.Pull
type MailAggregatorClientMockPullExpectation struct {
	mock    *MailAggregatorClientMock
	params  *MailAggregatorClientMockPullParams
	results *MailAggregatorClientMockPullResults
	Counter uint64
}

// MailAggregatorClientMockPullParams contains parameters of the MailAggregatorClient.Pull
type MailAggregatorClientMockPullParams struct {
	ctx  context.Context
	in   *MailboxGet
	opts []grpc.CallOption
}

// MailAggregatorClientMockPullResults contains results of the MailAggregatorClient.Pull
type MailAggregatorClientMockPullResults struct {
	mp1 *Messages
	err error
}

// Expect sets up expected params for MailAggregatorClient.Pull
func (mmPull *mMailAggregatorClientMockPull) Expect(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) *mMailAggregatorClientMockPull {
	if mmPull.mock.funcPull != nil {
		mmPull.mock.t.Fatalf("MailAggregatorClientMock.Pull mock is already set by Set")
	}

	if mmPull.defaultExpectation == nil {
		mmPull.defaultExpectation = &MailAggregatorClientMockPullExpectation{}
	}

	mmPull.defaultExpectation.params = &MailAggregatorClientMockPullParams{ctx, in, opts}
	for _, e := range mmPull.expectations {
		if minimock.Equal(e.params, mmPull.defaultExpectation.params) {
			mmPull.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPull.defaultExpectation.params)
		}
	}

	return mmPull
}

// Inspect accepts an inspector function that has same arguments as the MailAggregatorClient.Pull
func (mmPull *mMailAggregatorClientMockPull) Inspect(f func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption)) *mMailAggregatorClientMockPull {
	if mmPull.mock.inspectFuncPull != nil {
		mmPull.mock.t.Fatalf("Inspect function is already set for MailAggregatorClientMock.Pull")
	}

	mmPull.mock.inspectFuncPull = f

	return mmPull
}

// Return sets up results that will be returned by MailAggregatorClient.Pull
func (mmPull *mMailAggregatorClientMockPull) Return(mp1 *Messages, err error) *MailAggregatorClientMock {
	if mmPull.mock.funcPull != nil {
		mmPull.mock.t.Fatalf("MailAggregatorClientMock.Pull mock is already set by Set")
	}

	if mmPull.defaultExpectation == nil {
		mmPull.defaultExpectation = &MailAggregatorClientMockPullExpectation{mock: mmPull.mock}
	}
	mmPull.defaultExpectation.results = &MailAggregatorClientMockPullResults{mp1, err}
	return mmPull.mock
}

//Set uses given function f to mock the MailAggregatorClient.Pull method
func (mmPull *mMailAggregatorClientMockPull) Set(f func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Messages, err error)) *MailAggregatorClientMock {
	if mmPull.defaultExpectation != nil {
		mmPull.mock.t.Fatalf("Default expectation is already set for the MailAggregatorClient.Pull method")
	}

	if len(mmPull.expectations) > 0 {
		mmPull.mock.t.Fatalf("Some expectations are already set for the MailAggregatorClient.Pull method")
	}

	mmPull.mock.funcPull = f
	return mmPull.mock
}

// When sets expectation for the MailAggregatorClient.Pull which will trigger the result defined by the following
// Then helper
func (mmPull *mMailAggregatorClientMockPull) When(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) *MailAggregatorClientMockPullExpectation {
	if mmPull.mock.funcPull != nil {
		mmPull.mock.t.Fatalf("MailAggregatorClientMock.Pull mock is already set by Set")
	}

	expectation := &MailAggregatorClientMockPullExpectation{
		mock:   mmPull.mock,
		params: &MailAggregatorClientMockPullParams{ctx, in, opts},
	}
	mmPull.expectations = append(mmPull.expectations, expectation)
	return expectation
}

// Then sets up MailAggregatorClient.Pull return parameters for the expectation previously defined by the When method
func (e *MailAggregatorClientMockPullExpectation) Then(mp1 *Messages, err error) *MailAggregatorClientMock {
	e.results = &MailAggregatorClientMockPullResults{mp1, err}
	return e.mock
}

// Pull implements MailAggregatorClient
func (mmPull *MailAggregatorClientMock) Pull(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Messages, err error) {
	mm_atomic.AddUint64(&mmPull.beforePullCounter, 1)
	defer mm_atomic.AddUint64(&mmPull.afterPullCounter, 1)

	if mmPull.inspectFuncPull != nil {
		mmPull.inspectFuncPull(ctx, in, opts...)
	}

	mm_params := &MailAggregatorClientMockPullParams{ctx, in, opts}

	// Record call args
	mmPull.PullMock.mutex.Lock()
	mmPull.PullMock.callArgs = append(mmPull.PullMock.callArgs, mm_params)
	mmPull.PullMock.mutex.Unlock()

	for _, e := range mmPull.PullMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.mp1, e.results.err
		}
	}

	if mmPull.PullMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPull.PullMock.defaultExpectation.Counter, 1)
		mm_want := mmPull.PullMock.defaultExpectation.params
		mm_got := MailAggregatorClientMockPullParams{ctx, in, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPull.t.Errorf("MailAggregatorClientMock.Pull got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPull.PullMock.defaultExpectation.results
		if mm_results == nil {
			mmPull.t.Fatal("No results are set for the MailAggregatorClientMock.Pull")
		}
		return (*mm_results).mp1, (*mm_results).err
	}
	if mmPull.funcPull != nil {
		return mmPull.funcPull(ctx, in, opts...)
	}
	mmPull.t.Fatalf("Unexpected call to MailAggregatorClientMock.Pull. %v %v %v", ctx, in, opts)
	return
}

// PullAfterCounter returns a count of finished MailAggregatorClientMock.Pull invocations
func (mmPull *MailAggregatorClientMock) PullAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPull.afterPullCounter)
}

// PullBeforeCounter returns a count of MailAggregatorClientMock.Pull invocations
func (mmPull *MailAggregatorClientMock) PullBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPull.beforePullCounter)
}

// Calls returns a list of arguments used in each call to MailAggregatorClientMock.Pull.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPull *mMailAggregatorClientMockPull) Calls() []*MailAggregatorClientMockPullParams {
	mmPull.mutex.RLock()

	argCopy := make([]*MailAggregatorClientMockPullParams, len(mmPull.callArgs))
	copy(argCopy, mmPull.callArgs)

	mmPull.mutex.RUnlock()

	return argCopy
}

// MinimockPullDone returns true if the count of the Pull invocations corresponds
// the number of defined expectations
func (m *MailAggregatorClientMock) MinimockPullDone() bool {
	for _, e := range m.PullMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PullMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPullCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPull != nil && mm_atomic.LoadUint64(&m.afterPullCounter) < 1 {
		return false
	}
	return true
}

// MinimockPullInspect logs each unmet expectation
func (m *MailAggregatorClientMock) MinimockPullInspect() {
	for _, e := range m.PullMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Pull with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PullMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPullCounter) < 1 {
		if m.PullMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MailAggregatorClientMock.Pull")
		} else {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Pull with params: %#v", *m.PullMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPull != nil && mm_atomic.LoadUint64(&m.afterPullCounter) < 1 {
		m.t.Error("Expected call to MailAggregatorClientMock.Pull")
	}
}

type mMailAggregatorClientMockSelect struct {
	mock               *MailAggregatorClientMock
	defaultExpectation *MailAggregatorClientMockSelectExpectation
	expectations       []*MailAggregatorClientMockSelectExpectation

	callArgs []*MailAggregatorClientMockSelectParams
	mutex    sync.RWMutex
}

// MailAggregatorClientMockSelectExpectation specifies expectation struct of the MailAggregatorClient.Select
type MailAggregatorClientMockSelectExpectation struct {
	mock    *MailAggregatorClientMock
	params  *MailAggregatorClientMockSelectParams
	results *MailAggregatorClientMockSelectResults
	Counter uint64
}

// MailAggregatorClientMockSelectParams contains parameters of the MailAggregatorClient.Select
type MailAggregatorClientMockSelectParams struct {
	ctx  context.Context
	in   *MailboxGet
	opts []grpc.CallOption
}

// MailAggregatorClientMockSelectResults contains results of the MailAggregatorClient.Select
type MailAggregatorClientMockSelectResults struct {
	mp1 *Mailboxes
	err error
}

// Expect sets up expected params for MailAggregatorClient.Select
func (mmSelect *mMailAggregatorClientMockSelect) Expect(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) *mMailAggregatorClientMockSelect {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("MailAggregatorClientMock.Select mock is already set by Set")
	}

	if mmSelect.defaultExpectation == nil {
		mmSelect.defaultExpectation = &MailAggregatorClientMockSelectExpectation{}
	}

	mmSelect.defaultExpectation.params = &MailAggregatorClientMockSelectParams{ctx, in, opts}
	for _, e := range mmSelect.expectations {
		if minimock.Equal(e.params, mmSelect.defaultExpectation.params) {
			mmSelect.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSelect.defaultExpectation.params)
		}
	}

	return mmSelect
}

// Inspect accepts an inspector function that has same arguments as the MailAggregatorClient.Select
func (mmSelect *mMailAggregatorClientMockSelect) Inspect(f func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption)) *mMailAggregatorClientMockSelect {
	if mmSelect.mock.inspectFuncSelect != nil {
		mmSelect.mock.t.Fatalf("Inspect function is already set for MailAggregatorClientMock.Select")
	}

	mmSelect.mock.inspectFuncSelect = f

	return mmSelect
}

// Return sets up results that will be returned by MailAggregatorClient.Select
func (mmSelect *mMailAggregatorClientMockSelect) Return(mp1 *Mailboxes, err error) *MailAggregatorClientMock {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("MailAggregatorClientMock.Select mock is already set by Set")
	}

	if mmSelect.defaultExpectation == nil {
		mmSelect.defaultExpectation = &MailAggregatorClientMockSelectExpectation{mock: mmSelect.mock}
	}
	mmSelect.defaultExpectation.results = &MailAggregatorClientMockSelectResults{mp1, err}
	return mmSelect.mock
}

//Set uses given function f to mock the MailAggregatorClient.Select method
func (mmSelect *mMailAggregatorClientMockSelect) Set(f func(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Mailboxes, err error)) *MailAggregatorClientMock {
	if mmSelect.defaultExpectation != nil {
		mmSelect.mock.t.Fatalf("Default expectation is already set for the MailAggregatorClient.Select method")
	}

	if len(mmSelect.expectations) > 0 {
		mmSelect.mock.t.Fatalf("Some expectations are already set for the MailAggregatorClient.Select method")
	}

	mmSelect.mock.funcSelect = f
	return mmSelect.mock
}

// When sets expectation for the MailAggregatorClient.Select which will trigger the result defined by the following
// Then helper
func (mmSelect *mMailAggregatorClientMockSelect) When(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) *MailAggregatorClientMockSelectExpectation {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("MailAggregatorClientMock.Select mock is already set by Set")
	}

	expectation := &MailAggregatorClientMockSelectExpectation{
		mock:   mmSelect.mock,
		params: &MailAggregatorClientMockSelectParams{ctx, in, opts},
	}
	mmSelect.expectations = append(mmSelect.expectations, expectation)
	return expectation
}

// Then sets up MailAggregatorClient.Select return parameters for the expectation previously defined by the When method
func (e *MailAggregatorClientMockSelectExpectation) Then(mp1 *Mailboxes, err error) *MailAggregatorClientMock {
	e.results = &MailAggregatorClientMockSelectResults{mp1, err}
	return e.mock
}

// Select implements MailAggregatorClient
func (mmSelect *MailAggregatorClientMock) Select(ctx context.Context, in *MailboxGet, opts ...grpc.CallOption) (mp1 *Mailboxes, err error) {
	mm_atomic.AddUint64(&mmSelect.beforeSelectCounter, 1)
	defer mm_atomic.AddUint64(&mmSelect.afterSelectCounter, 1)

	if mmSelect.inspectFuncSelect != nil {
		mmSelect.inspectFuncSelect(ctx, in, opts...)
	}

	mm_params := &MailAggregatorClientMockSelectParams{ctx, in, opts}

	// Record call args
	mmSelect.SelectMock.mutex.Lock()
	mmSelect.SelectMock.callArgs = append(mmSelect.SelectMock.callArgs, mm_params)
	mmSelect.SelectMock.mutex.Unlock()

	for _, e := range mmSelect.SelectMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.mp1, e.results.err
		}
	}

	if mmSelect.SelectMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSelect.SelectMock.defaultExpectation.Counter, 1)
		mm_want := mmSelect.SelectMock.defaultExpectation.params
		mm_got := MailAggregatorClientMockSelectParams{ctx, in, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSelect.t.Errorf("MailAggregatorClientMock.Select got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSelect.SelectMock.defaultExpectation.results
		if mm_results == nil {
			mmSelect.t.Fatal("No results are set for the MailAggregatorClientMock.Select")
		}
		return (*mm_results).mp1, (*mm_results).err
	}
	if mmSelect.funcSelect != nil {
		return mmSelect.funcSelect(ctx, in, opts...)
	}
	mmSelect.t.Fatalf("Unexpected call to MailAggregatorClientMock.Select. %v %v %v", ctx, in, opts)
	return
}

// SelectAfterCounter returns a count of finished MailAggregatorClientMock.Select invocations
func (mmSelect *MailAggregatorClientMock) SelectAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSelect.afterSelectCounter)
}

// SelectBeforeCounter returns a count of MailAggregatorClientMock.Select invocations
func (mmSelect *MailAggregatorClientMock) SelectBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSelect.beforeSelectCounter)
}

// Calls returns a list of arguments used in each call to MailAggregatorClientMock.Select.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSelect *mMailAggregatorClientMockSelect) Calls() []*MailAggregatorClientMockSelectParams {
	mmSelect.mutex.RLock()

	argCopy := make([]*MailAggregatorClientMockSelectParams, len(mmSelect.callArgs))
	copy(argCopy, mmSelect.callArgs)

	mmSelect.mutex.RUnlock()

	return argCopy
}

// MinimockSelectDone returns true if the count of the Select invocations corresponds
// the number of defined expectations
func (m *MailAggregatorClientMock) MinimockSelectDone() bool {
	for _, e := range m.SelectMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SelectMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSelect != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		return false
	}
	return true
}

// MinimockSelectInspect logs each unmet expectation
func (m *MailAggregatorClientMock) MinimockSelectInspect() {
	for _, e := range m.SelectMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Select with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SelectMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		if m.SelectMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MailAggregatorClientMock.Select")
		} else {
			m.t.Errorf("Expected call to MailAggregatorClientMock.Select with params: %#v", *m.SelectMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSelect != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		m.t.Error("Expected call to MailAggregatorClientMock.Select")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MailAggregatorClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockPullInspect()

		m.MinimockSelectInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MailAggregatorClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MailAggregatorClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockPullDone() &&
		m.MinimockSelectDone()
}