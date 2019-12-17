// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

// Package config defines data structure and logic for loading and
// refreshing configurations for Honeydipper

package gitinventory

import (
	"github.com/honeydipper/honeydipper/internal/inventory"
	"github.com/honeydipper/honeydipper/pkg/dipper"
)

// RepoInfo points to a git repo where config data can be read from.
type RepoInfo struct {
	Repo        string
	Branch      string
	Path        string
	Name        string
	Description string
}

// GitInventory is a inventory with git backend
type GitInventory struct {
	InitRepo          RepoInfo
	DataSet           *inventory.DataSet
	Loaded            map[RepoInfo]*Repo
	WorkingDir        string
	LastRunningConfig struct {
		DataSet *inventory.DataSet
		Loaded  map[RepoInfo]*Repo
	}
}

// New initialize a inventory with git backend
func New(initRepo *RepoInfo) *GitInventory {
	return &GitInventory{
		InitRepo: *initRepo,
	}
}

// Bootstrap loads the init repo and all referenced repos
func (g *GitInventory) Bootstrap() {
	if g.WorkingDir == "" {
		g.WorkingDir = "/tmp"
	}
	g.loadRepo(g.InitRepo)
	g.assemble()
}

func (g *GitInventory) isRepoLoaded(repoInfo RepoInfo) bool {
	_, ok := g.Loaded[repoInfo]
	return ok
}

func (g *GitInventory) loadRepo(repoInfo RepoInfo) {
	if !g.isRepoLoaded(repoInfo) {
		repo := newRepo(g, repoInfo)
		repo.load()
		if g.Loaded == nil {
			g.Loaded = map[RepoInfo]*Repo{}
		}
		g.Loaded[repoInfo] = repo
	}
}

func (g *GitInventory) assemble() {
	g.DataSet, g.Loaded = g.Loaded[g.InitRepo].assemble(&(inventory.DataSet{}), map[RepoInfo]*Repo{})

	// post loading tasks
	g.extendAllSystems()
	g.parseRegex()
}

func (g *GitInventory) extendAllSystems() {
	for _, system := range g.DataSet.Systems {
		system.Extend(g)
	}
}

func (g *GitInventory) parseRegex() {
	var processor func(key string, val interface{}) (interface{}, bool)

	processor = func(name string, val interface{}) (interface{}, bool) {
		switch v := val.(type) {
		case string:
			return dipper.RegexParser(name, val)
		case inventory.Rule:
			dipper.Recursive(&v.Do, processor)
		case inventory.Workflow:
			dipper.Recursive(v.Match, processor)
			dipper.Recursive(v.Steps, processor)
			dipper.Recursive(v.Threads, processor)
			dipper.Recursive(v.Else, processor)
			dipper.Recursive(v.Cases, processor)
		}
		return nil, false
	}

	dipper.Recursive(g.DataSet.Workflows, processor)
	dipper.Recursive(g.DataSet.Rules, processor)
	dipper.Recursive(g.DataSet.Contexts, dipper.RegexParser)
}
