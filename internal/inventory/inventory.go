// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package inventory

import (
	"fmt"
)

// Inventory is the top level configuration container
type Inventory interface {
	Root() Scope
}

// ScopedWorkflow
type ScopedWorkflow struct {
	workflow *Workflow
	scope    Scope
}

// Scope provides configuration assets
type Scope interface {
	Parent() Scope
	// GetCollapsedRulesByDriver(string) []*CollapsedRule
	// GetTriggeredRules(event string, data interface{}) []*ScopedWorkflow
	GetAsset(assetType string, name string) interface{}

	GetSystem(name string) *System
	GetWorkflow(name string) *ScopedWorkflow
	GetContext(name string) map[string]interface{}
	GetDriverMeta(name string) map[string]interface{}
	GetDriverData(name string) map[string]interface{}
}

// ScopeProvider methods to be implemented to provide the configuration assets
type ScopeProvider interface {
	GetLocalSystem(name string) *System
	GetLocalWorkflow(name string) *ScopedWorkflow
	GetLocalContext(name string) map[string]interface{}
	GetLocalDriverMeta(name string) map[string]interface{}
	GetLocalDriverData(name string) map[string]interface{}
	Load()
}

// ScopeBase is the basic implementation of the Scope
type ScopeBase struct {
	InitPolicy string

	AllowWorkflows []string
	AllowContexts  []string
	AllowSystems   []string
	AllowDrivers   []string
	BlockWorkflows []string
	BlockContexts  []string
	BlockSystems   []string
	BlockDrivers   []string

	LoadWorkflows bool
	LoadContexts  bool
	LoadSystems   bool
	LoadDrivers   bool

	parent Scope
}

// SetParent sets the parent scope for the policy
func (p *ScopeBase) SetParent(parent Scope) {
	p.parent = parent
}

func (p *ScopeBase) getLocalAsset(assetType string, name string) interface{} {
	var self, asset interface{} = p, nil
	scope := self.(ScopeProvider)

	switch assetType {
	case "system":
		asset = scope.GetLocalSystem(name)
	case "workflow":
		asset = scope.GetLocalWorkflow(name)
	case "context":
		asset = scope.GetLocalContext(name)
	case "driverMeta":
		asset = scope.GetLocalDriverMeta(name)
	case "driverData":
		asset = scope.GetLocalDriverData(name)
	default:
		panic(fmt.Errorf("unknown asset type: %s", assetType))
	}

	return asset
}

func (p *ScopeBase) isBlocked(assetType string, name string) bool {
	if p.InitPolicy == "allow" {
		var blocked []string
		switch assetType {
		case "system":
			blocked = p.BlockSystems
		case "workflow":
			blocked = p.BlockWorkflows
		case "context":
			blocked = p.BlockContexts
		case "driverMeta":
			blocked = p.BlockDrivers
		case "driverData":
			blocked = p.BlockDrivers
		default:
			panic(fmt.Errorf("unknown asset type: %s", assetType))
		}
		for _, n := range blocked {
			if n == name {
				return true
			}
		}
		return false
	} else {
		var allowed []string
		switch assetType {
		case "system":
			allowed = p.AllowSystems
		case "workflow":
			allowed = p.AllowWorkflows
		case "context":
			allowed = p.AllowContexts
		case "driverMeta":
			allowed = p.AllowDrivers
		case "driverData":
			allowed = p.AllowDrivers
		default:
			panic(fmt.Errorf("unknown asset type: %s", assetType))
		}
		for _, n := range allowed {
			if n == name {
				return false
			}
		}
		return true
	}
}

// GetAsset gets the configuration asset with given name and type
func (p *ScopeBase) GetAsset(assetType string, name string) interface{} {
	asset := p.getLocalAsset(assetType, name)

	if asset == nil && !p.isBlocked(assetType, name) && p.parent != nil {
		return p.parent.GetAsset(assetType, name)
	}

	return asset
}
