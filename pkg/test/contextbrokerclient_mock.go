// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package client

import (
	"context"
	"sync"

	"github.com/diwise/context-broker/pkg/ngsild"
	. "github.com/diwise/context-broker/pkg/ngsild/client"
	"github.com/diwise/context-broker/pkg/ngsild/types"
)

// Ensure, that ContextBrokerClientMock does implement ContextBrokerClient.
// If this is not the case, regenerate this file with moq.
var _ ContextBrokerClient = &ContextBrokerClientMock{}

// ContextBrokerClientMock is a mock implementation of ContextBrokerClient.
//
// 	func TestSomethingThatUsesContextBrokerClient(t *testing.T) {
//
// 		// make and configure a mocked ContextBrokerClient
// 		mockedContextBrokerClient := &ContextBrokerClientMock{
// 			CreateEntityFunc: func(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
// 				panic("mock out the CreateEntity method")
// 			},
// 			DeleteEntityFunc: func(ctx context.Context, entityID string) (*ngsild.DeleteEntityResult, error) {
// 				panic("mock out the DeleteEntity method")
// 			},
// 			MergeEntityFunc: func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error) {
// 				panic("mock out the MergeEntity method")
// 			},
// 			QueryEntitiesFunc: func(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
// 				panic("mock out the QueryEntities method")
// 			},
// 			QueryTemporalEvolutionOfEntitiesFunc: func(ctx context.Context, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.QueryTemporalEntitiesResult, error) {
// 				panic("mock out the QueryTemporalEvolutionOfEntities method")
// 			},
// 			RetrieveEntityFunc: func(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error) {
// 				panic("mock out the RetrieveEntity method")
// 			},
// 			RetrieveTemporalEvolutionOfEntityFunc: func(ctx context.Context, entityID string, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.RetrieveTemporalEvolutionOfEntityResult, error) {
// 				panic("mock out the RetrieveTemporalEvolutionOfEntity method")
// 			},
// 			UpdateEntityAttributesFunc: func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
// 				panic("mock out the UpdateEntityAttributes method")
// 			},
// 		}
//
// 		// use mockedContextBrokerClient in code that requires ContextBrokerClient
// 		// and then make assertions.
//
// 	}
type ContextBrokerClientMock struct {
	// CreateEntityFunc mocks the CreateEntity method.
	CreateEntityFunc func(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error)

	// DeleteEntityFunc mocks the DeleteEntity method.
	DeleteEntityFunc func(ctx context.Context, entityID string) (*ngsild.DeleteEntityResult, error)

	// MergeEntityFunc mocks the MergeEntity method.
	MergeEntityFunc func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error)

	// QueryEntitiesFunc mocks the QueryEntities method.
	QueryEntitiesFunc func(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error)

	// QueryTemporalEvolutionOfEntitiesFunc mocks the QueryTemporalEvolutionOfEntities method.
	QueryTemporalEvolutionOfEntitiesFunc func(ctx context.Context, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.QueryTemporalEntitiesResult, error)

	// RetrieveEntityFunc mocks the RetrieveEntity method.
	RetrieveEntityFunc func(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error)

	// RetrieveTemporalEvolutionOfEntityFunc mocks the RetrieveTemporalEvolutionOfEntity method.
	RetrieveTemporalEvolutionOfEntityFunc func(ctx context.Context, entityID string, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.RetrieveTemporalEvolutionOfEntityResult, error)

	// UpdateEntityAttributesFunc mocks the UpdateEntityAttributes method.
	UpdateEntityAttributesFunc func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateEntity holds details about calls to the CreateEntity method.
		CreateEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Entity is the entity argument value.
			Entity types.Entity
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// DeleteEntity holds details about calls to the DeleteEntity method.
		DeleteEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityID is the entityID argument value.
			EntityID string
		}
		// MergeEntity holds details about calls to the MergeEntity method.
		MergeEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
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
			// Headers is the headers argument value.
			Headers map[string][]string
			// Parameters is the parameters argument value.
			Parameters []RequestDecoratorFunc
		}
		// RetrieveEntity holds details about calls to the RetrieveEntity method.
		RetrieveEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityID is the entityID argument value.
			EntityID string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// RetrieveTemporalEvolutionOfEntity holds details about calls to the RetrieveTemporalEvolutionOfEntity method.
		RetrieveTemporalEvolutionOfEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityID is the entityID argument value.
			EntityID string
			// Headers is the headers argument value.
			Headers map[string][]string
			// Parameters is the parameters argument value.
			Parameters []RequestDecoratorFunc
		}
		// UpdateEntityAttributes holds details about calls to the UpdateEntityAttributes method.
		UpdateEntityAttributes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
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
	lockUpdateEntityAttributes            sync.RWMutex
}

// CreateEntity calls CreateEntityFunc.
func (mock *ContextBrokerClientMock) CreateEntity(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
	if mock.CreateEntityFunc == nil {
		panic("ContextBrokerClientMock.CreateEntityFunc: method is nil but ContextBrokerClient.CreateEntity was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Entity  types.Entity
		Headers map[string][]string
	}{
		Ctx:     ctx,
		Entity:  entity,
		Headers: headers,
	}
	mock.lockCreateEntity.Lock()
	mock.calls.CreateEntity = append(mock.calls.CreateEntity, callInfo)
	mock.lockCreateEntity.Unlock()
	return mock.CreateEntityFunc(ctx, entity, headers)
}

// CreateEntityCalls gets all the calls that were made to CreateEntity.
// Check the length with:
//     len(mockedContextBrokerClient.CreateEntityCalls())
func (mock *ContextBrokerClientMock) CreateEntityCalls() []struct {
	Ctx     context.Context
	Entity  types.Entity
	Headers map[string][]string
} {
	var calls []struct {
		Ctx     context.Context
		Entity  types.Entity
		Headers map[string][]string
	}
	mock.lockCreateEntity.RLock()
	calls = mock.calls.CreateEntity
	mock.lockCreateEntity.RUnlock()
	return calls
}

// DeleteEntity calls DeleteEntityFunc.
func (mock *ContextBrokerClientMock) DeleteEntity(ctx context.Context, entityID string) (*ngsild.DeleteEntityResult, error) {
	if mock.DeleteEntityFunc == nil {
		panic("ContextBrokerClientMock.DeleteEntityFunc: method is nil but ContextBrokerClient.DeleteEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
	}{
		Ctx:      ctx,
		EntityID: entityID,
	}
	mock.lockDeleteEntity.Lock()
	mock.calls.DeleteEntity = append(mock.calls.DeleteEntity, callInfo)
	mock.lockDeleteEntity.Unlock()
	return mock.DeleteEntityFunc(ctx, entityID)
}

// DeleteEntityCalls gets all the calls that were made to DeleteEntity.
// Check the length with:
//     len(mockedContextBrokerClient.DeleteEntityCalls())
func (mock *ContextBrokerClientMock) DeleteEntityCalls() []struct {
	Ctx      context.Context
	EntityID string
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
	}
	mock.lockDeleteEntity.RLock()
	calls = mock.calls.DeleteEntity
	mock.lockDeleteEntity.RUnlock()
	return calls
}

// MergeEntity calls MergeEntityFunc.
func (mock *ContextBrokerClientMock) MergeEntity(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.MergeEntityResult, error) {
	if mock.MergeEntityFunc == nil {
		panic("ContextBrokerClientMock.MergeEntityFunc: method is nil but ContextBrokerClient.MergeEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockMergeEntity.Lock()
	mock.calls.MergeEntity = append(mock.calls.MergeEntity, callInfo)
	mock.lockMergeEntity.Unlock()
	return mock.MergeEntityFunc(ctx, entityID, fragment, headers)
}

// MergeEntityCalls gets all the calls that were made to MergeEntity.
// Check the length with:
//     len(mockedContextBrokerClient.MergeEntityCalls())
func (mock *ContextBrokerClientMock) MergeEntityCalls() []struct {
	Ctx      context.Context
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
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
func (mock *ContextBrokerClientMock) QueryEntities(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
	if mock.QueryEntitiesFunc == nil {
		panic("ContextBrokerClientMock.QueryEntitiesFunc: method is nil but ContextBrokerClient.QueryEntities was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}{
		Ctx:              ctx,
		EntityTypes:      entityTypes,
		EntityAttributes: entityAttributes,
		Query:            query,
		Headers:          headers,
	}
	mock.lockQueryEntities.Lock()
	mock.calls.QueryEntities = append(mock.calls.QueryEntities, callInfo)
	mock.lockQueryEntities.Unlock()
	return mock.QueryEntitiesFunc(ctx, entityTypes, entityAttributes, query, headers)
}

// QueryEntitiesCalls gets all the calls that were made to QueryEntities.
// Check the length with:
//     len(mockedContextBrokerClient.QueryEntitiesCalls())
func (mock *ContextBrokerClientMock) QueryEntitiesCalls() []struct {
	Ctx              context.Context
	EntityTypes      []string
	EntityAttributes []string
	Query            string
	Headers          map[string][]string
} {
	var calls []struct {
		Ctx              context.Context
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
func (mock *ContextBrokerClientMock) QueryTemporalEvolutionOfEntities(ctx context.Context, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.QueryTemporalEntitiesResult, error) {
	if mock.QueryTemporalEvolutionOfEntitiesFunc == nil {
		panic("ContextBrokerClientMock.QueryTemporalEvolutionOfEntitiesFunc: method is nil but ContextBrokerClient.QueryTemporalEvolutionOfEntities was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Headers    map[string][]string
		Parameters []RequestDecoratorFunc
	}{
		Ctx:        ctx,
		Headers:    headers,
		Parameters: parameters,
	}
	mock.lockQueryTemporalEvolutionOfEntities.Lock()
	mock.calls.QueryTemporalEvolutionOfEntities = append(mock.calls.QueryTemporalEvolutionOfEntities, callInfo)
	mock.lockQueryTemporalEvolutionOfEntities.Unlock()
	return mock.QueryTemporalEvolutionOfEntitiesFunc(ctx, headers, parameters...)
}

// QueryTemporalEvolutionOfEntitiesCalls gets all the calls that were made to QueryTemporalEvolutionOfEntities.
// Check the length with:
//     len(mockedContextBrokerClient.QueryTemporalEvolutionOfEntitiesCalls())
func (mock *ContextBrokerClientMock) QueryTemporalEvolutionOfEntitiesCalls() []struct {
	Ctx        context.Context
	Headers    map[string][]string
	Parameters []RequestDecoratorFunc
} {
	var calls []struct {
		Ctx        context.Context
		Headers    map[string][]string
		Parameters []RequestDecoratorFunc
	}
	mock.lockQueryTemporalEvolutionOfEntities.RLock()
	calls = mock.calls.QueryTemporalEvolutionOfEntities
	mock.lockQueryTemporalEvolutionOfEntities.RUnlock()
	return calls
}

// RetrieveEntity calls RetrieveEntityFunc.
func (mock *ContextBrokerClientMock) RetrieveEntity(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error) {
	if mock.RetrieveEntityFunc == nil {
		panic("ContextBrokerClientMock.RetrieveEntityFunc: method is nil but ContextBrokerClient.RetrieveEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Headers:  headers,
	}
	mock.lockRetrieveEntity.Lock()
	mock.calls.RetrieveEntity = append(mock.calls.RetrieveEntity, callInfo)
	mock.lockRetrieveEntity.Unlock()
	return mock.RetrieveEntityFunc(ctx, entityID, headers)
}

// RetrieveEntityCalls gets all the calls that were made to RetrieveEntity.
// Check the length with:
//     len(mockedContextBrokerClient.RetrieveEntityCalls())
func (mock *ContextBrokerClientMock) RetrieveEntityCalls() []struct {
	Ctx      context.Context
	EntityID string
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
		Headers  map[string][]string
	}
	mock.lockRetrieveEntity.RLock()
	calls = mock.calls.RetrieveEntity
	mock.lockRetrieveEntity.RUnlock()
	return calls
}

// RetrieveTemporalEvolutionOfEntity calls RetrieveTemporalEvolutionOfEntityFunc.
func (mock *ContextBrokerClientMock) RetrieveTemporalEvolutionOfEntity(ctx context.Context, entityID string, headers map[string][]string, parameters ...RequestDecoratorFunc) (*ngsild.RetrieveTemporalEvolutionOfEntityResult, error) {
	if mock.RetrieveTemporalEvolutionOfEntityFunc == nil {
		panic("ContextBrokerClientMock.RetrieveTemporalEvolutionOfEntityFunc: method is nil but ContextBrokerClient.RetrieveTemporalEvolutionOfEntity was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		EntityID   string
		Headers    map[string][]string
		Parameters []RequestDecoratorFunc
	}{
		Ctx:        ctx,
		EntityID:   entityID,
		Headers:    headers,
		Parameters: parameters,
	}
	mock.lockRetrieveTemporalEvolutionOfEntity.Lock()
	mock.calls.RetrieveTemporalEvolutionOfEntity = append(mock.calls.RetrieveTemporalEvolutionOfEntity, callInfo)
	mock.lockRetrieveTemporalEvolutionOfEntity.Unlock()
	return mock.RetrieveTemporalEvolutionOfEntityFunc(ctx, entityID, headers, parameters...)
}

// RetrieveTemporalEvolutionOfEntityCalls gets all the calls that were made to RetrieveTemporalEvolutionOfEntity.
// Check the length with:
//     len(mockedContextBrokerClient.RetrieveTemporalEvolutionOfEntityCalls())
func (mock *ContextBrokerClientMock) RetrieveTemporalEvolutionOfEntityCalls() []struct {
	Ctx        context.Context
	EntityID   string
	Headers    map[string][]string
	Parameters []RequestDecoratorFunc
} {
	var calls []struct {
		Ctx        context.Context
		EntityID   string
		Headers    map[string][]string
		Parameters []RequestDecoratorFunc
	}
	mock.lockRetrieveTemporalEvolutionOfEntity.RLock()
	calls = mock.calls.RetrieveTemporalEvolutionOfEntity
	mock.lockRetrieveTemporalEvolutionOfEntity.RUnlock()
	return calls
}

// UpdateEntityAttributes calls UpdateEntityAttributesFunc.
func (mock *ContextBrokerClientMock) UpdateEntityAttributes(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
	if mock.UpdateEntityAttributesFunc == nil {
		panic("ContextBrokerClientMock.UpdateEntityAttributesFunc: method is nil but ContextBrokerClient.UpdateEntityAttributes was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockUpdateEntityAttributes.Lock()
	mock.calls.UpdateEntityAttributes = append(mock.calls.UpdateEntityAttributes, callInfo)
	mock.lockUpdateEntityAttributes.Unlock()
	return mock.UpdateEntityAttributesFunc(ctx, entityID, fragment, headers)
}

// UpdateEntityAttributesCalls gets all the calls that were made to UpdateEntityAttributes.
// Check the length with:
//     len(mockedContextBrokerClient.UpdateEntityAttributesCalls())
func (mock *ContextBrokerClientMock) UpdateEntityAttributesCalls() []struct {
	Ctx      context.Context
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}
	mock.lockUpdateEntityAttributes.RLock()
	calls = mock.calls.UpdateEntityAttributes
	mock.lockUpdateEntityAttributes.RUnlock()
	return calls
}
