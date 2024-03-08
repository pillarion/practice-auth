// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mock

//go:generate minimock -i github.com/pillarion/practice-auth/internal/core/port/repository/user.Repo -o repo_minimock.go -n RepoMock -p mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
)

// RepoMock implements user.Repo
type RepoMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mRepoMockDelete

	funcInsert          func(ctx context.Context, user *model.Info) (i1 int64, err error)
	inspectFuncInsert   func(ctx context.Context, user *model.Info)
	afterInsertCounter  uint64
	beforeInsertCounter uint64
	InsertMock          mRepoMockInsert

	funcSelect          func(ctx context.Context, id int64) (up1 *model.User, err error)
	inspectFuncSelect   func(ctx context.Context, id int64)
	afterSelectCounter  uint64
	beforeSelectCounter uint64
	SelectMock          mRepoMockSelect

	funcUpdate          func(ctx context.Context, user *model.Info) (err error)
	inspectFuncUpdate   func(ctx context.Context, user *model.Info)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mRepoMockUpdate
}

// NewRepoMock returns a mock for user.Repo
func NewRepoMock(t minimock.Tester) *RepoMock {
	m := &RepoMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mRepoMockDelete{mock: m}
	m.DeleteMock.callArgs = []*RepoMockDeleteParams{}

	m.InsertMock = mRepoMockInsert{mock: m}
	m.InsertMock.callArgs = []*RepoMockInsertParams{}

	m.SelectMock = mRepoMockSelect{mock: m}
	m.SelectMock.callArgs = []*RepoMockSelectParams{}

	m.UpdateMock = mRepoMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*RepoMockUpdateParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRepoMockDelete struct {
	mock               *RepoMock
	defaultExpectation *RepoMockDeleteExpectation
	expectations       []*RepoMockDeleteExpectation

	callArgs []*RepoMockDeleteParams
	mutex    sync.RWMutex
}

// RepoMockDeleteExpectation specifies expectation struct of the Repo.Delete
type RepoMockDeleteExpectation struct {
	mock    *RepoMock
	params  *RepoMockDeleteParams
	results *RepoMockDeleteResults
	Counter uint64
}

// RepoMockDeleteParams contains parameters of the Repo.Delete
type RepoMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// RepoMockDeleteResults contains results of the Repo.Delete
type RepoMockDeleteResults struct {
	err error
}

// Expect sets up expected params for Repo.Delete
func (mmDelete *mRepoMockDelete) Expect(ctx context.Context, id int64) *mRepoMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepoMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &RepoMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &RepoMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Repo.Delete
func (mmDelete *mRepoMockDelete) Inspect(f func(ctx context.Context, id int64)) *mRepoMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for RepoMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Repo.Delete
func (mmDelete *mRepoMockDelete) Return(err error) *RepoMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepoMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &RepoMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &RepoMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the Repo.Delete method
func (mmDelete *mRepoMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *RepoMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Repo.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Repo.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Repo.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mRepoMockDelete) When(ctx context.Context, id int64) *RepoMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("RepoMock.Delete mock is already set by Set")
	}

	expectation := &RepoMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &RepoMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Repo.Delete return parameters for the expectation previously defined by the When method
func (e *RepoMockDeleteExpectation) Then(err error) *RepoMock {
	e.results = &RepoMockDeleteResults{err}
	return e.mock
}

// Delete implements user.Repo
func (mmDelete *RepoMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := RepoMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := RepoMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("RepoMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the RepoMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to RepoMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished RepoMock.Delete invocations
func (mmDelete *RepoMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of RepoMock.Delete invocations
func (mmDelete *RepoMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mRepoMockDelete) Calls() []*RepoMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*RepoMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockDeleteDone() bool {
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
func (m *RepoMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Delete")
		} else {
			m.t.Errorf("Expected call to RepoMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Delete")
	}
}

type mRepoMockInsert struct {
	mock               *RepoMock
	defaultExpectation *RepoMockInsertExpectation
	expectations       []*RepoMockInsertExpectation

	callArgs []*RepoMockInsertParams
	mutex    sync.RWMutex
}

// RepoMockInsertExpectation specifies expectation struct of the Repo.Insert
type RepoMockInsertExpectation struct {
	mock    *RepoMock
	params  *RepoMockInsertParams
	results *RepoMockInsertResults
	Counter uint64
}

// RepoMockInsertParams contains parameters of the Repo.Insert
type RepoMockInsertParams struct {
	ctx  context.Context
	user *model.Info
}

// RepoMockInsertResults contains results of the Repo.Insert
type RepoMockInsertResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for Repo.Insert
func (mmInsert *mRepoMockInsert) Expect(ctx context.Context, user *model.Info) *mRepoMockInsert {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	if mmInsert.defaultExpectation == nil {
		mmInsert.defaultExpectation = &RepoMockInsertExpectation{}
	}

	mmInsert.defaultExpectation.params = &RepoMockInsertParams{ctx, user}
	for _, e := range mmInsert.expectations {
		if minimock.Equal(e.params, mmInsert.defaultExpectation.params) {
			mmInsert.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInsert.defaultExpectation.params)
		}
	}

	return mmInsert
}

// Inspect accepts an inspector function that has same arguments as the Repo.Insert
func (mmInsert *mRepoMockInsert) Inspect(f func(ctx context.Context, user *model.Info)) *mRepoMockInsert {
	if mmInsert.mock.inspectFuncInsert != nil {
		mmInsert.mock.t.Fatalf("Inspect function is already set for RepoMock.Insert")
	}

	mmInsert.mock.inspectFuncInsert = f

	return mmInsert
}

// Return sets up results that will be returned by Repo.Insert
func (mmInsert *mRepoMockInsert) Return(i1 int64, err error) *RepoMock {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	if mmInsert.defaultExpectation == nil {
		mmInsert.defaultExpectation = &RepoMockInsertExpectation{mock: mmInsert.mock}
	}
	mmInsert.defaultExpectation.results = &RepoMockInsertResults{i1, err}
	return mmInsert.mock
}

// Set uses given function f to mock the Repo.Insert method
func (mmInsert *mRepoMockInsert) Set(f func(ctx context.Context, user *model.Info) (i1 int64, err error)) *RepoMock {
	if mmInsert.defaultExpectation != nil {
		mmInsert.mock.t.Fatalf("Default expectation is already set for the Repo.Insert method")
	}

	if len(mmInsert.expectations) > 0 {
		mmInsert.mock.t.Fatalf("Some expectations are already set for the Repo.Insert method")
	}

	mmInsert.mock.funcInsert = f
	return mmInsert.mock
}

// When sets expectation for the Repo.Insert which will trigger the result defined by the following
// Then helper
func (mmInsert *mRepoMockInsert) When(ctx context.Context, user *model.Info) *RepoMockInsertExpectation {
	if mmInsert.mock.funcInsert != nil {
		mmInsert.mock.t.Fatalf("RepoMock.Insert mock is already set by Set")
	}

	expectation := &RepoMockInsertExpectation{
		mock:   mmInsert.mock,
		params: &RepoMockInsertParams{ctx, user},
	}
	mmInsert.expectations = append(mmInsert.expectations, expectation)
	return expectation
}

// Then sets up Repo.Insert return parameters for the expectation previously defined by the When method
func (e *RepoMockInsertExpectation) Then(i1 int64, err error) *RepoMock {
	e.results = &RepoMockInsertResults{i1, err}
	return e.mock
}

// Insert implements user.Repo
func (mmInsert *RepoMock) Insert(ctx context.Context, user *model.Info) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmInsert.beforeInsertCounter, 1)
	defer mm_atomic.AddUint64(&mmInsert.afterInsertCounter, 1)

	if mmInsert.inspectFuncInsert != nil {
		mmInsert.inspectFuncInsert(ctx, user)
	}

	mm_params := RepoMockInsertParams{ctx, user}

	// Record call args
	mmInsert.InsertMock.mutex.Lock()
	mmInsert.InsertMock.callArgs = append(mmInsert.InsertMock.callArgs, &mm_params)
	mmInsert.InsertMock.mutex.Unlock()

	for _, e := range mmInsert.InsertMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmInsert.InsertMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInsert.InsertMock.defaultExpectation.Counter, 1)
		mm_want := mmInsert.InsertMock.defaultExpectation.params
		mm_got := RepoMockInsertParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInsert.t.Errorf("RepoMock.Insert got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInsert.InsertMock.defaultExpectation.results
		if mm_results == nil {
			mmInsert.t.Fatal("No results are set for the RepoMock.Insert")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmInsert.funcInsert != nil {
		return mmInsert.funcInsert(ctx, user)
	}
	mmInsert.t.Fatalf("Unexpected call to RepoMock.Insert. %v %v", ctx, user)
	return
}

// InsertAfterCounter returns a count of finished RepoMock.Insert invocations
func (mmInsert *RepoMock) InsertAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsert.afterInsertCounter)
}

// InsertBeforeCounter returns a count of RepoMock.Insert invocations
func (mmInsert *RepoMock) InsertBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsert.beforeInsertCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Insert.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInsert *mRepoMockInsert) Calls() []*RepoMockInsertParams {
	mmInsert.mutex.RLock()

	argCopy := make([]*RepoMockInsertParams, len(mmInsert.callArgs))
	copy(argCopy, mmInsert.callArgs)

	mmInsert.mutex.RUnlock()

	return argCopy
}

// MinimockInsertDone returns true if the count of the Insert invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockInsertDone() bool {
	for _, e := range m.InsertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsert != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertInspect logs each unmet expectation
func (m *RepoMock) MinimockInsertInspect() {
	for _, e := range m.InsertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Insert with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		if m.InsertMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Insert")
		} else {
			m.t.Errorf("Expected call to RepoMock.Insert with params: %#v", *m.InsertMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsert != nil && mm_atomic.LoadUint64(&m.afterInsertCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Insert")
	}
}

type mRepoMockSelect struct {
	mock               *RepoMock
	defaultExpectation *RepoMockSelectExpectation
	expectations       []*RepoMockSelectExpectation

	callArgs []*RepoMockSelectParams
	mutex    sync.RWMutex
}

// RepoMockSelectExpectation specifies expectation struct of the Repo.Select
type RepoMockSelectExpectation struct {
	mock    *RepoMock
	params  *RepoMockSelectParams
	results *RepoMockSelectResults
	Counter uint64
}

// RepoMockSelectParams contains parameters of the Repo.Select
type RepoMockSelectParams struct {
	ctx context.Context
	id  int64
}

// RepoMockSelectResults contains results of the Repo.Select
type RepoMockSelectResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for Repo.Select
func (mmSelect *mRepoMockSelect) Expect(ctx context.Context, id int64) *mRepoMockSelect {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("RepoMock.Select mock is already set by Set")
	}

	if mmSelect.defaultExpectation == nil {
		mmSelect.defaultExpectation = &RepoMockSelectExpectation{}
	}

	mmSelect.defaultExpectation.params = &RepoMockSelectParams{ctx, id}
	for _, e := range mmSelect.expectations {
		if minimock.Equal(e.params, mmSelect.defaultExpectation.params) {
			mmSelect.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSelect.defaultExpectation.params)
		}
	}

	return mmSelect
}

// Inspect accepts an inspector function that has same arguments as the Repo.Select
func (mmSelect *mRepoMockSelect) Inspect(f func(ctx context.Context, id int64)) *mRepoMockSelect {
	if mmSelect.mock.inspectFuncSelect != nil {
		mmSelect.mock.t.Fatalf("Inspect function is already set for RepoMock.Select")
	}

	mmSelect.mock.inspectFuncSelect = f

	return mmSelect
}

// Return sets up results that will be returned by Repo.Select
func (mmSelect *mRepoMockSelect) Return(up1 *model.User, err error) *RepoMock {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("RepoMock.Select mock is already set by Set")
	}

	if mmSelect.defaultExpectation == nil {
		mmSelect.defaultExpectation = &RepoMockSelectExpectation{mock: mmSelect.mock}
	}
	mmSelect.defaultExpectation.results = &RepoMockSelectResults{up1, err}
	return mmSelect.mock
}

// Set uses given function f to mock the Repo.Select method
func (mmSelect *mRepoMockSelect) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *RepoMock {
	if mmSelect.defaultExpectation != nil {
		mmSelect.mock.t.Fatalf("Default expectation is already set for the Repo.Select method")
	}

	if len(mmSelect.expectations) > 0 {
		mmSelect.mock.t.Fatalf("Some expectations are already set for the Repo.Select method")
	}

	mmSelect.mock.funcSelect = f
	return mmSelect.mock
}

// When sets expectation for the Repo.Select which will trigger the result defined by the following
// Then helper
func (mmSelect *mRepoMockSelect) When(ctx context.Context, id int64) *RepoMockSelectExpectation {
	if mmSelect.mock.funcSelect != nil {
		mmSelect.mock.t.Fatalf("RepoMock.Select mock is already set by Set")
	}

	expectation := &RepoMockSelectExpectation{
		mock:   mmSelect.mock,
		params: &RepoMockSelectParams{ctx, id},
	}
	mmSelect.expectations = append(mmSelect.expectations, expectation)
	return expectation
}

// Then sets up Repo.Select return parameters for the expectation previously defined by the When method
func (e *RepoMockSelectExpectation) Then(up1 *model.User, err error) *RepoMock {
	e.results = &RepoMockSelectResults{up1, err}
	return e.mock
}

// Select implements user.Repo
func (mmSelect *RepoMock) Select(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmSelect.beforeSelectCounter, 1)
	defer mm_atomic.AddUint64(&mmSelect.afterSelectCounter, 1)

	if mmSelect.inspectFuncSelect != nil {
		mmSelect.inspectFuncSelect(ctx, id)
	}

	mm_params := RepoMockSelectParams{ctx, id}

	// Record call args
	mmSelect.SelectMock.mutex.Lock()
	mmSelect.SelectMock.callArgs = append(mmSelect.SelectMock.callArgs, &mm_params)
	mmSelect.SelectMock.mutex.Unlock()

	for _, e := range mmSelect.SelectMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmSelect.SelectMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSelect.SelectMock.defaultExpectation.Counter, 1)
		mm_want := mmSelect.SelectMock.defaultExpectation.params
		mm_got := RepoMockSelectParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSelect.t.Errorf("RepoMock.Select got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSelect.SelectMock.defaultExpectation.results
		if mm_results == nil {
			mmSelect.t.Fatal("No results are set for the RepoMock.Select")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmSelect.funcSelect != nil {
		return mmSelect.funcSelect(ctx, id)
	}
	mmSelect.t.Fatalf("Unexpected call to RepoMock.Select. %v %v", ctx, id)
	return
}

// SelectAfterCounter returns a count of finished RepoMock.Select invocations
func (mmSelect *RepoMock) SelectAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSelect.afterSelectCounter)
}

// SelectBeforeCounter returns a count of RepoMock.Select invocations
func (mmSelect *RepoMock) SelectBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSelect.beforeSelectCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Select.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSelect *mRepoMockSelect) Calls() []*RepoMockSelectParams {
	mmSelect.mutex.RLock()

	argCopy := make([]*RepoMockSelectParams, len(mmSelect.callArgs))
	copy(argCopy, mmSelect.callArgs)

	mmSelect.mutex.RUnlock()

	return argCopy
}

// MinimockSelectDone returns true if the count of the Select invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockSelectDone() bool {
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
func (m *RepoMock) MinimockSelectInspect() {
	for _, e := range m.SelectMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Select with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SelectMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		if m.SelectMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Select")
		} else {
			m.t.Errorf("Expected call to RepoMock.Select with params: %#v", *m.SelectMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSelect != nil && mm_atomic.LoadUint64(&m.afterSelectCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Select")
	}
}

type mRepoMockUpdate struct {
	mock               *RepoMock
	defaultExpectation *RepoMockUpdateExpectation
	expectations       []*RepoMockUpdateExpectation

	callArgs []*RepoMockUpdateParams
	mutex    sync.RWMutex
}

// RepoMockUpdateExpectation specifies expectation struct of the Repo.Update
type RepoMockUpdateExpectation struct {
	mock    *RepoMock
	params  *RepoMockUpdateParams
	results *RepoMockUpdateResults
	Counter uint64
}

// RepoMockUpdateParams contains parameters of the Repo.Update
type RepoMockUpdateParams struct {
	ctx  context.Context
	user *model.Info
}

// RepoMockUpdateResults contains results of the Repo.Update
type RepoMockUpdateResults struct {
	err error
}

// Expect sets up expected params for Repo.Update
func (mmUpdate *mRepoMockUpdate) Expect(ctx context.Context, user *model.Info) *mRepoMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepoMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepoMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &RepoMockUpdateParams{ctx, user}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the Repo.Update
func (mmUpdate *mRepoMockUpdate) Inspect(f func(ctx context.Context, user *model.Info)) *mRepoMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for RepoMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by Repo.Update
func (mmUpdate *mRepoMockUpdate) Return(err error) *RepoMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepoMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &RepoMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &RepoMockUpdateResults{err}
	return mmUpdate.mock
}

// Set uses given function f to mock the Repo.Update method
func (mmUpdate *mRepoMockUpdate) Set(f func(ctx context.Context, user *model.Info) (err error)) *RepoMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the Repo.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the Repo.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the Repo.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mRepoMockUpdate) When(ctx context.Context, user *model.Info) *RepoMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("RepoMock.Update mock is already set by Set")
	}

	expectation := &RepoMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &RepoMockUpdateParams{ctx, user},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up Repo.Update return parameters for the expectation previously defined by the When method
func (e *RepoMockUpdateExpectation) Then(err error) *RepoMock {
	e.results = &RepoMockUpdateResults{err}
	return e.mock
}

// Update implements user.Repo
func (mmUpdate *RepoMock) Update(ctx context.Context, user *model.Info) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(ctx, user)
	}

	mm_params := RepoMockUpdateParams{ctx, user}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, &mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := RepoMockUpdateParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("RepoMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the RepoMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(ctx, user)
	}
	mmUpdate.t.Fatalf("Unexpected call to RepoMock.Update. %v %v", ctx, user)
	return
}

// UpdateAfterCounter returns a count of finished RepoMock.Update invocations
func (mmUpdate *RepoMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of RepoMock.Update invocations
func (mmUpdate *RepoMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to RepoMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mRepoMockUpdate) Calls() []*RepoMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*RepoMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *RepoMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *RepoMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoMock.Update")
		} else {
			m.t.Errorf("Expected call to RepoMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to RepoMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepoMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockDeleteInspect()

			m.MinimockInsertInspect()

			m.MinimockSelectInspect()

			m.MinimockUpdateInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockInsertDone() &&
		m.MinimockSelectDone() &&
		m.MinimockUpdateDone()
}
