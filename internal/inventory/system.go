// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package inventory

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/imdario/mergo"
)

// Copy performs a deep copy of the given system
func (s *System) Copy() *System {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(*s)
	if err != nil {
		panic(err)
	}

	var scopy System
	err = dec.Decode(&scopy)
	if err != nil {
		panic(err)
	}
	return &scopy
}

// MergeSystem merges the given source system with current system
func (s *System) MergeSystem(src *System) {
	for name, trigger := range src.Triggers {
		if s.Triggers == nil {
			s.Triggers = map[string]Trigger{}
		}

		exist, ok := s.Triggers[name]
		if ok {
			err := mergo.Merge(&exist, trigger, mergo.WithOverride, mergo.WithAppendSlice)
			if err != nil {
				panic(err)
			}

			// mergo.Merge won't merge value with different type
			if exist.Description == "" {
				exist.Description = trigger.Description
			}
			if exist.Meta == nil {
				exist.Meta = trigger.Meta
			}
		} else {
			exist = trigger
		}

		s.Triggers[name] = exist
	}

	for name, function := range src.Functions {
		if s.Functions == nil {
			s.Functions = map[string]Function{}
		}

		exist, ok := s.Functions[name]
		if ok {
			err := mergo.Merge(&exist, function, mergo.WithOverride, mergo.WithAppendSlice)
			if err != nil {
				panic(err)
			}

			// mergo.Merge won't merge value with different type
			if exist.Description == "" {
				exist.Description = function.Description
			}
			if exist.Meta == nil {
				exist.Meta = function.Meta
			}
		} else {
			exist = function
		}

		s.Functions[name] = exist
	}

	err := mergo.Merge(&s.Data, src.Data, mergo.WithOverride, mergo.WithAppendSlice)
	if err != nil {
		panic(err)
	}

	s.Extends = append(s.Extends, src.Extends...)

	if src.Description != "" {
		s.Description = src.Description
	}
	if src.Meta != nil {
		s.Meta = src.Meta
	}
}

// AddSubSystem adds a subsystem to the current system
func (s *System) AddSubSystem(subKey string, src *System) {
	for name, trigger := range src.Triggers {
		if s.Triggers == nil {
			s.Triggers = map[string]Trigger{}
		}
		s.Triggers[subKey+"."+name] = trigger
	}

	for name, function := range src.Functions {
		if s.Functions == nil {
			s.Functions = map[string]Function{}
		}
		s.Functions[subKey+"."+name] = function
	}

	if s.Data == nil {
		s.Data = map[string]interface{}{}
	}
	s.Data[subKey] = src.Data
}

// Extend loads all dependent and subsystems of the system
func (s *System) Extend(inv Inventory) {
	if s.extended {
		return
	}

	var merged = &System{
		extended: true,
	}

	for _, extend := range s.Extends {
		parts := strings.Split(extend, "=")
		var base, subKey string
		base = strings.TrimSpace(parts[0])
		if len(parts) >= 2 {
			subKey = base
			base = strings.TrimSpace(parts[1])
		}

		baseSys := inv.Root().GetSystem(base)
		if baseSys == nil {
			panic(fmt.Errorf("system not found %s", base))
		}

		baseSys.Extend(inv)
		baseCopy := baseSys.Copy()

		if subKey != "" {
			merged.AddSubSystem(subKey, baseCopy)
		} else {
			merged.MergeSystem(baseCopy)
		}
	}

	merged.MergeSystem(s)
	*s = *merged
}
