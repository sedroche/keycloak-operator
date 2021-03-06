// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package realm

import (
	"github.com/integr8ly/keycloak-operator/pkg/apis/aerogear/v1alpha1"
	"sync"
)

var (
	lockHandlerMockAccepted        sync.RWMutex
	lockHandlerMockDeprovision     sync.RWMutex
	lockHandlerMockInitialise      sync.RWMutex
	lockHandlerMockPreflightChecks sync.RWMutex
	lockHandlerMockProvision       sync.RWMutex
	lockHandlerMockReconcile       sync.RWMutex
)

// HandlerMock is a mock implementation of Handler.
//
//     func TestSomethingThatUsesHandler(t *testing.T) {
//
//         // make and configure a mocked Handler
//         mockedHandler := &HandlerMock{
//             AcceptedFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the Accepted method")
//             },
//             DeprovisionFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the Deprovision method")
//             },
//             InitialiseFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the Initialise method")
//             },
//             PreflightChecksFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the PreflightChecks method")
//             },
//             ProvisionFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the Provision method")
//             },
//             ReconcileFunc: func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
// 	               panic("mock out the Reconcile method")
//             },
//         }
//
//         // use mockedHandler in code that requires Handler
//         // and then make assertions.
//
//     }
type HandlerMock struct {
	// AcceptedFunc mocks the Accepted method.
	AcceptedFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// DeprovisionFunc mocks the Deprovision method.
	DeprovisionFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// InitialiseFunc mocks the Initialise method.
	InitialiseFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// PreflightChecksFunc mocks the PreflightChecks method.
	PreflightChecksFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// ProvisionFunc mocks the Provision method.
	ProvisionFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// ReconcileFunc mocks the Reconcile method.
	ReconcileFunc func(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error)

	// calls tracks calls to the methods.
	calls struct {
		// Accepted holds details about calls to the Accepted method.
		Accepted []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
		// Deprovision holds details about calls to the Deprovision method.
		Deprovision []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
		// Initialise holds details about calls to the Initialise method.
		Initialise []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
		// PreflightChecks holds details about calls to the PreflightChecks method.
		PreflightChecks []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
		// Provision holds details about calls to the Provision method.
		Provision []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
		// Reconcile holds details about calls to the Reconcile method.
		Reconcile []struct {
			// Realm is the realm argument value.
			Realm *v1alpha1.KeycloakRealm
		}
	}
}

// Accepted calls AcceptedFunc.
func (mock *HandlerMock) Accepted(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.AcceptedFunc == nil {
		panic("HandlerMock.AcceptedFunc: method is nil but Handler.Accepted was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockAccepted.Lock()
	mock.calls.Accepted = append(mock.calls.Accepted, callInfo)
	lockHandlerMockAccepted.Unlock()
	return mock.AcceptedFunc(realm)
}

// AcceptedCalls gets all the calls that were made to Accepted.
// Check the length with:
//     len(mockedHandler.AcceptedCalls())
func (mock *HandlerMock) AcceptedCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockAccepted.RLock()
	calls = mock.calls.Accepted
	lockHandlerMockAccepted.RUnlock()
	return calls
}

// Deprovision calls DeprovisionFunc.
func (mock *HandlerMock) Deprovision(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.DeprovisionFunc == nil {
		panic("HandlerMock.DeprovisionFunc: method is nil but Handler.Deprovision was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockDeprovision.Lock()
	mock.calls.Deprovision = append(mock.calls.Deprovision, callInfo)
	lockHandlerMockDeprovision.Unlock()
	return mock.DeprovisionFunc(realm)
}

// DeprovisionCalls gets all the calls that were made to Deprovision.
// Check the length with:
//     len(mockedHandler.DeprovisionCalls())
func (mock *HandlerMock) DeprovisionCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockDeprovision.RLock()
	calls = mock.calls.Deprovision
	lockHandlerMockDeprovision.RUnlock()
	return calls
}

// Initialise calls InitialiseFunc.
func (mock *HandlerMock) Initialise(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.InitialiseFunc == nil {
		panic("HandlerMock.InitialiseFunc: method is nil but Handler.Initialise was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockInitialise.Lock()
	mock.calls.Initialise = append(mock.calls.Initialise, callInfo)
	lockHandlerMockInitialise.Unlock()
	return mock.InitialiseFunc(realm)
}

// InitialiseCalls gets all the calls that were made to Initialise.
// Check the length with:
//     len(mockedHandler.InitialiseCalls())
func (mock *HandlerMock) InitialiseCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockInitialise.RLock()
	calls = mock.calls.Initialise
	lockHandlerMockInitialise.RUnlock()
	return calls
}

// PreflightChecks calls PreflightChecksFunc.
func (mock *HandlerMock) PreflightChecks(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.PreflightChecksFunc == nil {
		panic("HandlerMock.PreflightChecksFunc: method is nil but Handler.PreflightChecks was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockPreflightChecks.Lock()
	mock.calls.PreflightChecks = append(mock.calls.PreflightChecks, callInfo)
	lockHandlerMockPreflightChecks.Unlock()
	return mock.PreflightChecksFunc(realm)
}

// PreflightChecksCalls gets all the calls that were made to PreflightChecks.
// Check the length with:
//     len(mockedHandler.PreflightChecksCalls())
func (mock *HandlerMock) PreflightChecksCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockPreflightChecks.RLock()
	calls = mock.calls.PreflightChecks
	lockHandlerMockPreflightChecks.RUnlock()
	return calls
}

// Provision calls ProvisionFunc.
func (mock *HandlerMock) Provision(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.ProvisionFunc == nil {
		panic("HandlerMock.ProvisionFunc: method is nil but Handler.Provision was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockProvision.Lock()
	mock.calls.Provision = append(mock.calls.Provision, callInfo)
	lockHandlerMockProvision.Unlock()
	return mock.ProvisionFunc(realm)
}

// ProvisionCalls gets all the calls that were made to Provision.
// Check the length with:
//     len(mockedHandler.ProvisionCalls())
func (mock *HandlerMock) ProvisionCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockProvision.RLock()
	calls = mock.calls.Provision
	lockHandlerMockProvision.RUnlock()
	return calls
}

// Reconcile calls ReconcileFunc.
func (mock *HandlerMock) Reconcile(realm *v1alpha1.KeycloakRealm) (*v1alpha1.KeycloakRealm, error) {
	if mock.ReconcileFunc == nil {
		panic("HandlerMock.ReconcileFunc: method is nil but Handler.Reconcile was just called")
	}
	callInfo := struct {
		Realm *v1alpha1.KeycloakRealm
	}{
		Realm: realm,
	}
	lockHandlerMockReconcile.Lock()
	mock.calls.Reconcile = append(mock.calls.Reconcile, callInfo)
	lockHandlerMockReconcile.Unlock()
	return mock.ReconcileFunc(realm)
}

// ReconcileCalls gets all the calls that were made to Reconcile.
// Check the length with:
//     len(mockedHandler.ReconcileCalls())
func (mock *HandlerMock) ReconcileCalls() []struct {
	Realm *v1alpha1.KeycloakRealm
} {
	var calls []struct {
		Realm *v1alpha1.KeycloakRealm
	}
	lockHandlerMockReconcile.RLock()
	calls = mock.calls.Reconcile
	lockHandlerMockReconcile.RUnlock()
	return calls
}
