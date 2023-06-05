// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cim

import (
	"context"
	"github.com/diwise/context-broker/pkg/ngsild"
	"github.com/diwise/context-broker/pkg/ngsild/types"
	"sync"
)

// Ensure, that ContextInformationManagerMock does implement ContextInformationManager.
// If this is not the case, regenerate this file with moq.
var _ ContextInformationManager = &ContextInformationManagerMock{}

// ContextInformationManagerMock is a mock implementation of ContextInformationManager.
//
//	func TestSomethingThatUsesContextInformationManager(t *testing.T) {
//
//		// make and configure a mocked ContextInformationManager
//		mockedContextInformationManager := &ContextInformationManagerMock{
//			CreateEntityFunc: func(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
//				panic("mock out the CreateEntity method")
//			},
//			DeleteEntityFunc: func(ctx context.Context, tenant string, entityID string) (*ngsild.DeleteEntityResult, error) {
//				panic("mock out the DeleteEntity method")
//			},
//			MergeEntityFunc: func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error) {
//				panic("mock out the MergeEntity method")
//			},
//			QueryEntitiesFunc: func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
//				panic("mock out the QueryEntities method")
//			},
//			QueryTemporalEvolutionOfEntitiesFunc: func(ctx context.Context, tenant string, entityIDs []string, entityTypes []string, params TemporalQueryParams, headers map[string][]string) (*ngsild.QueryTemporalEntitiesResult, error) {
//				panic("mock out the QueryTemporalEvolutionOfEntities method")
//			},
//			RetrieveEntityFunc: func(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error) {
//				panic("mock out the RetrieveEntity method")
//			},
//			RetrieveTemporalEvolutionOfEntityFunc: func(ctx context.Context, tenant string, entityID string, params TemporalQueryParams, headers map[string][]string) (types.EntityTemporal, error) {
//				panic("mock out the RetrieveTemporalEvolutionOfEntity method")
//			},
//			RetrieveTypesFunc: func(ctx context.Context, tenant string, headers map[string][]string) ([]string, error) {
//				panic("mock out the RetrieveTypes method")
//			},
//			StartFunc: func() error {
//				panic("mock out the Start method")
//			},
//			StopFunc: func() error {
//				panic("mock out the Stop method")
//			},
//			UpdateEntityAttributesFunc: func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
//				panic("mock out the UpdateEntityAttributes method")
//			},
//		}
//
//		// use mockedContextInformationManager in code that requires ContextInformationManager
//		// and then make assertions.
//
//	}
type ContextInformationManagerMock struct {
	// CreateEntityFunc mocks the CreateEntity method.
	CreateEntityFunc func(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error)

	// DeleteEntityFunc mocks the DeleteEntity method.
	DeleteEntityFunc func(ctx context.Context, tenant string, entityID string) (*ngsild.DeleteEntityResult, error)

	// MergeEntityFunc mocks the MergeEntity method.
	MergeEntityFunc func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error)

	// QueryEntitiesFunc mocks the QueryEntities method.
	QueryEntitiesFunc func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error)

	// QueryTemporalEvolutionOfEntitiesFunc mocks the QueryTemporalEvolutionOfEntities method.
	QueryTemporalEvolutionOfEntitiesFunc func(ctx context.Context, tenant string, entityIDs []string, entityTypes []string, params TemporalQueryParams, headers map[string][]string) (*ngsild.QueryTemporalEntitiesResult, error)

	// RetrieveEntityFunc mocks the RetrieveEntity method.
	RetrieveEntityFunc func(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error)

	// RetrieveTemporalEvolutionOfEntityFunc mocks the RetrieveTemporalEvolutionOfEntity method.
	RetrieveTemporalEvolutionOfEntityFunc func(ctx context.Context, tenant string, entityID string, params TemporalQueryParams, headers map[string][]string) (types.EntityTemporal, error)

	// RetrieveTypesFunc mocks the RetrieveTypes method.
	RetrieveTypesFunc func(ctx context.Context, tenant string, headers map[string][]string) ([]string, error)

	// StartFunc mocks the Start method.
	StartFunc func() error

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// UpdateEntityAttributesFunc mocks the UpdateEntityAttributes method.
	UpdateEntityAttributesFunc func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateEntity holds details about calls to the CreateEntity method.
		CreateEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// Entity is the entity argument value.
			Entity types.Entity
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// DeleteEntity holds details about calls to the DeleteEntity method.
		DeleteEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
		}
		// MergeEntity holds details about calls to the MergeEntity method.
		MergeEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Fragment is the fragment argument value.
			Fragment types.EntityFragment
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// QueryEntities holds details about calls to the QueryEntities method.
		QueryEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityTypes is the entityTypes argument value.
			EntityTypes []string
			// EntityAttributes is the entityAttributes argument value.
			EntityAttributes []string
			// Query is the query argument value.
			Query string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// QueryTemporalEvolutionOfEntities holds details about calls to the QueryTemporalEvolutionOfEntities method.
		QueryTemporalEvolutionOfEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityIDs is the entityIDs argument value.
			EntityIDs []string
			// EntityTypes is the entityTypes argument value.
			EntityTypes []string
			// Params is the params argument value.
			Params TemporalQueryParams
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// RetrieveEntity holds details about calls to the RetrieveEntity method.
		RetrieveEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// RetrieveTemporalEvolutionOfEntity holds details about calls to the RetrieveTemporalEvolutionOfEntity method.
		RetrieveTemporalEvolutionOfEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Params is the params argument value.
			Params TemporalQueryParams
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// RetrieveTypes holds details about calls to the RetrieveTypes method.
		RetrieveTypes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// Start holds details about calls to the Start method.
		Start []struct {
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
		// UpdateEntityAttributes holds details about calls to the UpdateEntityAttributes method.
		UpdateEntityAttributes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Fragment is the fragment argument value.
			Fragment types.EntityFragment
			// Headers is the headers argument value.
			Headers map[string][]string
		}
	}
	lockCreateEntity                      sync.RWMutex
	lockDeleteEntity                      sync.RWMutex
	lockMergeEntity                       sync.RWMutex
	lockQueryEntities                     sync.RWMutex
	lockQueryTemporalEvolutionOfEntities  sync.RWMutex
	lockRetrieveEntity                    sync.RWMutex
	lockRetrieveTemporalEvolutionOfEntity sync.RWMutex
	lockRetrieveTypes                     sync.RWMutex
	lockStart                             sync.RWMutex
	lockStop                              sync.RWMutex
	lockUpdateEntityAttributes            sync.RWMutex
}

// CreateEntity calls CreateEntityFunc.
func (mock *ContextInformationManagerMock) CreateEntity(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
	if mock.CreateEntityFunc == nil {
		panic("ContextInformationManagerMock.CreateEntityFunc: method is nil but ContextInformationManager.CreateEntity was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenant  string
		Entity  types.Entity
		Headers map[string][]string
	}{
		Ctx:     ctx,
		Tenant:  tenant,
		Entity:  entity,
		Headers: headers,
	}
	mock.lockCreateEntity.Lock()
	mock.calls.CreateEntity = append(mock.calls.CreateEntity, callInfo)
	mock.lockCreateEntity.Unlock()
	return mock.CreateEntityFunc(ctx, tenant, entity, headers)
}

// CreateEntityCalls gets all the calls that were made to CreateEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.CreateEntityCalls())
func (mock *ContextInformationManagerMock) CreateEntityCalls() []struct {
	Ctx     context.Context
	Tenant  string
	Entity  types.Entity
	Headers map[string][]string
} {
	var calls []struct {
		Ctx     context.Context
		Tenant  string
		Entity  types.Entity
		Headers map[string][]string
	}
	mock.lockCreateEntity.RLock()
	calls = mock.calls.CreateEntity
	mock.lockCreateEntity.RUnlock()
	return calls
}

// DeleteEntity calls DeleteEntityFunc.
func (mock *ContextInformationManagerMock) DeleteEntity(ctx context.Context, tenant string, entityID string) (*ngsild.DeleteEntityResult, error) {
	if mock.DeleteEntityFunc == nil {
		panic("ContextInformationManagerMock.DeleteEntityFunc: method is nil but ContextInformationManager.DeleteEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
	}
	mock.lockDeleteEntity.Lock()
	mock.calls.DeleteEntity = append(mock.calls.DeleteEntity, callInfo)
	mock.lockDeleteEntity.Unlock()
	return mock.DeleteEntityFunc(ctx, tenant, entityID)
}

// DeleteEntityCalls gets all the calls that were made to DeleteEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.DeleteEntityCalls())
func (mock *ContextInformationManagerMock) DeleteEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
	}
	mock.lockDeleteEntity.RLock()
	calls = mock.calls.DeleteEntity
	mock.lockDeleteEntity.RUnlock()
	return calls
}

// MergeEntity calls MergeEntityFunc.
func (mock *ContextInformationManagerMock) MergeEntity(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error) {
	if mock.MergeEntityFunc == nil {
		panic("ContextInformationManagerMock.MergeEntityFunc: method is nil but ContextInformationManager.MergeEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockMergeEntity.Lock()
	mock.calls.MergeEntity = append(mock.calls.MergeEntity, callInfo)
	mock.lockMergeEntity.Unlock()
	return mock.MergeEntityFunc(ctx, tenant, entityID, fragment, headers)
}

// MergeEntityCalls gets all the calls that were made to MergeEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.MergeEntityCalls())
func (mock *ContextInformationManagerMock) MergeEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}
	mock.lockMergeEntity.RLock()
	calls = mock.calls.MergeEntity
	mock.lockMergeEntity.RUnlock()
	return calls
}

// QueryEntities calls QueryEntitiesFunc.
func (mock *ContextInformationManagerMock) QueryEntities(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
	if mock.QueryEntitiesFunc == nil {
		panic("ContextInformationManagerMock.QueryEntitiesFunc: method is nil but ContextInformationManager.QueryEntities was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}{
		Ctx:              ctx,
		Tenant:           tenant,
		EntityTypes:      entityTypes,
		EntityAttributes: entityAttributes,
		Query:            query,
		Headers:          headers,
	}
	mock.lockQueryEntities.Lock()
	mock.calls.QueryEntities = append(mock.calls.QueryEntities, callInfo)
	mock.lockQueryEntities.Unlock()
	return mock.QueryEntitiesFunc(ctx, tenant, entityTypes, entityAttributes, query, headers)
}

// QueryEntitiesCalls gets all the calls that were made to QueryEntities.
// Check the length with:
//
//	len(mockedContextInformationManager.QueryEntitiesCalls())
func (mock *ContextInformationManagerMock) QueryEntitiesCalls() []struct {
	Ctx              context.Context
	Tenant           string
	EntityTypes      []string
	EntityAttributes []string
	Query            string
	Headers          map[string][]string
} {
	var calls []struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}
	mock.lockQueryEntities.RLock()
	calls = mock.calls.QueryEntities
	mock.lockQueryEntities.RUnlock()
	return calls
}

// QueryTemporalEvolutionOfEntities calls QueryTemporalEvolutionOfEntitiesFunc.
func (mock *ContextInformationManagerMock) QueryTemporalEvolutionOfEntities(ctx context.Context, tenant string, entityIDs []string, entityTypes []string, params TemporalQueryParams, headers map[string][]string) (*ngsild.QueryTemporalEntitiesResult, error) {
	if mock.QueryTemporalEvolutionOfEntitiesFunc == nil {
		panic("ContextInformationManagerMock.QueryTemporalEvolutionOfEntitiesFunc: method is nil but ContextInformationManager.QueryTemporalEvolutionOfEntities was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Tenant      string
		EntityIDs   []string
		EntityTypes []string
		Params      TemporalQueryParams
		Headers     map[string][]string
	}{
		Ctx:         ctx,
		Tenant:      tenant,
		EntityIDs:   entityIDs,
		EntityTypes: entityTypes,
		Params:      params,
		Headers:     headers,
	}
	mock.lockQueryTemporalEvolutionOfEntities.Lock()
	mock.calls.QueryTemporalEvolutionOfEntities = append(mock.calls.QueryTemporalEvolutionOfEntities, callInfo)
	mock.lockQueryTemporalEvolutionOfEntities.Unlock()
	return mock.QueryTemporalEvolutionOfEntitiesFunc(ctx, tenant, entityIDs, entityTypes, params, headers)
}

// QueryTemporalEvolutionOfEntitiesCalls gets all the calls that were made to QueryTemporalEvolutionOfEntities.
// Check the length with:
//
//	len(mockedContextInformationManager.QueryTemporalEvolutionOfEntitiesCalls())
func (mock *ContextInformationManagerMock) QueryTemporalEvolutionOfEntitiesCalls() []struct {
	Ctx         context.Context
	Tenant      string
	EntityIDs   []string
	EntityTypes []string
	Params      TemporalQueryParams
	Headers     map[string][]string
} {
	var calls []struct {
		Ctx         context.Context
		Tenant      string
		EntityIDs   []string
		EntityTypes []string
		Params      TemporalQueryParams
		Headers     map[string][]string
	}
	mock.lockQueryTemporalEvolutionOfEntities.RLock()
	calls = mock.calls.QueryTemporalEvolutionOfEntities
	mock.lockQueryTemporalEvolutionOfEntities.RUnlock()
	return calls
}

// RetrieveEntity calls RetrieveEntityFunc.
func (mock *ContextInformationManagerMock) RetrieveEntity(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error) {
	if mock.RetrieveEntityFunc == nil {
		panic("ContextInformationManagerMock.RetrieveEntityFunc: method is nil but ContextInformationManager.RetrieveEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Headers:  headers,
	}
	mock.lockRetrieveEntity.Lock()
	mock.calls.RetrieveEntity = append(mock.calls.RetrieveEntity, callInfo)
	mock.lockRetrieveEntity.Unlock()
	return mock.RetrieveEntityFunc(ctx, tenant, entityID, headers)
}

// RetrieveEntityCalls gets all the calls that were made to RetrieveEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.RetrieveEntityCalls())
func (mock *ContextInformationManagerMock) RetrieveEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Headers  map[string][]string
	}
	mock.lockRetrieveEntity.RLock()
	calls = mock.calls.RetrieveEntity
	mock.lockRetrieveEntity.RUnlock()
	return calls
}

// RetrieveTemporalEvolutionOfEntity calls RetrieveTemporalEvolutionOfEntityFunc.
func (mock *ContextInformationManagerMock) RetrieveTemporalEvolutionOfEntity(ctx context.Context, tenant string, entityID string, params TemporalQueryParams, headers map[string][]string) (types.EntityTemporal, error) {
	if mock.RetrieveTemporalEvolutionOfEntityFunc == nil {
		panic("ContextInformationManagerMock.RetrieveTemporalEvolutionOfEntityFunc: method is nil but ContextInformationManager.RetrieveTemporalEvolutionOfEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Params   TemporalQueryParams
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Params:   params,
		Headers:  headers,
	}
	mock.lockRetrieveTemporalEvolutionOfEntity.Lock()
	mock.calls.RetrieveTemporalEvolutionOfEntity = append(mock.calls.RetrieveTemporalEvolutionOfEntity, callInfo)
	mock.lockRetrieveTemporalEvolutionOfEntity.Unlock()
	return mock.RetrieveTemporalEvolutionOfEntityFunc(ctx, tenant, entityID, params, headers)
}

// RetrieveTemporalEvolutionOfEntityCalls gets all the calls that were made to RetrieveTemporalEvolutionOfEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.RetrieveTemporalEvolutionOfEntityCalls())
func (mock *ContextInformationManagerMock) RetrieveTemporalEvolutionOfEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Params   TemporalQueryParams
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Params   TemporalQueryParams
		Headers  map[string][]string
	}
	mock.lockRetrieveTemporalEvolutionOfEntity.RLock()
	calls = mock.calls.RetrieveTemporalEvolutionOfEntity
	mock.lockRetrieveTemporalEvolutionOfEntity.RUnlock()
	return calls
}

// RetrieveTypes calls RetrieveTypesFunc.
func (mock *ContextInformationManagerMock) RetrieveTypes(ctx context.Context, tenant string, headers map[string][]string) ([]string, error) {
	if mock.RetrieveTypesFunc == nil {
		panic("ContextInformationManagerMock.RetrieveTypesFunc: method is nil but ContextInformationManager.RetrieveTypes was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenant  string
		Headers map[string][]string
	}{
		Ctx:     ctx,
		Tenant:  tenant,
		Headers: headers,
	}
	mock.lockRetrieveTypes.Lock()
	mock.calls.RetrieveTypes = append(mock.calls.RetrieveTypes, callInfo)
	mock.lockRetrieveTypes.Unlock()
	return mock.RetrieveTypesFunc(ctx, tenant, headers)
}

// RetrieveTypesCalls gets all the calls that were made to RetrieveTypes.
// Check the length with:
//
//	len(mockedContextInformationManager.RetrieveTypesCalls())
func (mock *ContextInformationManagerMock) RetrieveTypesCalls() []struct {
	Ctx     context.Context
	Tenant  string
	Headers map[string][]string
} {
	var calls []struct {
		Ctx     context.Context
		Tenant  string
		Headers map[string][]string
	}
	mock.lockRetrieveTypes.RLock()
	calls = mock.calls.RetrieveTypes
	mock.lockRetrieveTypes.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *ContextInformationManagerMock) Start() error {
	if mock.StartFunc == nil {
		panic("ContextInformationManagerMock.StartFunc: method is nil but ContextInformationManager.Start was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	return mock.StartFunc()
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//
//	len(mockedContextInformationManager.StartCalls())
func (mock *ContextInformationManagerMock) StartCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *ContextInformationManagerMock) Stop() error {
	if mock.StopFunc == nil {
		panic("ContextInformationManagerMock.StopFunc: method is nil but ContextInformationManager.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//
//	len(mockedContextInformationManager.StopCalls())
func (mock *ContextInformationManagerMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}

// UpdateEntityAttributes calls UpdateEntityAttributesFunc.
func (mock *ContextInformationManagerMock) UpdateEntityAttributes(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
	if mock.UpdateEntityAttributesFunc == nil {
		panic("ContextInformationManagerMock.UpdateEntityAttributesFunc: method is nil but ContextInformationManager.UpdateEntityAttributes was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockUpdateEntityAttributes.Lock()
	mock.calls.UpdateEntityAttributes = append(mock.calls.UpdateEntityAttributes, callInfo)
	mock.lockUpdateEntityAttributes.Unlock()
	return mock.UpdateEntityAttributesFunc(ctx, tenant, entityID, fragment, headers)
}

// UpdateEntityAttributesCalls gets all the calls that were made to UpdateEntityAttributes.
// Check the length with:
//
//	len(mockedContextInformationManager.UpdateEntityAttributesCalls())
func (mock *ContextInformationManagerMock) UpdateEntityAttributesCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}
	mock.lockUpdateEntityAttributes.RLock()
	calls = mock.calls.UpdateEntityAttributes
	mock.lockUpdateEntityAttributes.RUnlock()
	return calls
}
