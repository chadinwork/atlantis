// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommandRequirementHandler)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	"reflect"
	"time"
)

type MockCommandRequirementHandler struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommandRequirementHandler(options ...pegomock.Option) *MockCommandRequirementHandler {
	mock := &MockCommandRequirementHandler{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommandRequirementHandler) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommandRequirementHandler) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommandRequirementHandler) ValidateApplyProject(repoDir string, ctx command.ProjectContext) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRequirementHandler().")
	}
	params := []pegomock.Param{repoDir, ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidateApplyProject", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommandRequirementHandler) ValidateProjectDependencies(_param0 command.ProjectContext) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRequirementHandler().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidateProjectDependencies", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommandRequirementHandler) ValidateImportProject(repoDir string, ctx command.ProjectContext) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRequirementHandler().")
	}
	params := []pegomock.Param{repoDir, ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidateImportProject", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommandRequirementHandler) ValidatePlanProject(repoDir string, ctx command.ProjectContext) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRequirementHandler().")
	}
	params := []pegomock.Param{repoDir, ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidatePlanProject", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommandRequirementHandler) VerifyWasCalledOnce() *VerifierMockCommandRequirementHandler {
	return &VerifierMockCommandRequirementHandler{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommandRequirementHandler) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCommandRequirementHandler {
	return &VerifierMockCommandRequirementHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommandRequirementHandler) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommandRequirementHandler {
	return &VerifierMockCommandRequirementHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommandRequirementHandler) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCommandRequirementHandler {
	return &VerifierMockCommandRequirementHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommandRequirementHandler struct {
	mock                   *MockCommandRequirementHandler
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommandRequirementHandler) ValidateApplyProject(repoDir string, ctx command.ProjectContext) *MockCommandRequirementHandler_ValidateApplyProject_OngoingVerification {
	params := []pegomock.Param{repoDir, ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ValidateApplyProject", params, verifier.timeout)
	return &MockCommandRequirementHandler_ValidateApplyProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommandRequirementHandler_ValidateApplyProject_OngoingVerification struct {
	mock              *MockCommandRequirementHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommandRequirementHandler_ValidateApplyProject_OngoingVerification) GetCapturedArguments() (string, command.ProjectContext) {
	repoDir, ctx := c.GetAllCapturedArguments()
	return repoDir[len(repoDir)-1], ctx[len(ctx)-1]
}

func (c *MockCommandRequirementHandler_ValidateApplyProject_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []command.ProjectContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(command.ProjectContext)
		}
	}
	return
}

func (verifier *VerifierMockCommandRequirementHandler) ValidateImportProject(repoDir string, ctx command.ProjectContext) *MockCommandRequirementHandler_ValidateImportProject_OngoingVerification {
	params := []pegomock.Param{repoDir, ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ValidateImportProject", params, verifier.timeout)
	return &MockCommandRequirementHandler_ValidateImportProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommandRequirementHandler_ValidateImportProject_OngoingVerification struct {
	mock              *MockCommandRequirementHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommandRequirementHandler_ValidateImportProject_OngoingVerification) GetCapturedArguments() (string, command.ProjectContext) {
	repoDir, ctx := c.GetAllCapturedArguments()
	return repoDir[len(repoDir)-1], ctx[len(ctx)-1]
}

func (c *MockCommandRequirementHandler_ValidateImportProject_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []command.ProjectContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(command.ProjectContext)
		}
	}
	return
}

func (verifier *VerifierMockCommandRequirementHandler) ValidatePlanProject(repoDir string, ctx command.ProjectContext) *MockCommandRequirementHandler_ValidatePlanProject_OngoingVerification {
	params := []pegomock.Param{repoDir, ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ValidatePlanProject", params, verifier.timeout)
	return &MockCommandRequirementHandler_ValidatePlanProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommandRequirementHandler_ValidatePlanProject_OngoingVerification struct {
	mock              *MockCommandRequirementHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommandRequirementHandler_ValidatePlanProject_OngoingVerification) GetCapturedArguments() (string, command.ProjectContext) {
	repoDir, ctx := c.GetAllCapturedArguments()
	return repoDir[len(repoDir)-1], ctx[len(ctx)-1]
}

func (c *MockCommandRequirementHandler_ValidatePlanProject_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []command.ProjectContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(command.ProjectContext)
		}
	}
	return
}
