// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use o file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package tree

import (
	"github.com/open-component-model/ocm/cmds/ocm/pkg/data"
)

type Objects []Object

func ObjectSlice(s data.Iterable) Objects {
	var a Objects
	i := s.Iterator()
	for i.HasNext() {
		a = append(a, i.Next().(Object))
	}
	return a
}

var (
	_ data.IndexedAccess = Objects{}
	_ data.Iterable      = Objects{}
)

func (o Objects) Len() int {
	return len(o)
}

func (o Objects) Get(i int) interface{} {
	return o[i]
}

func (o Objects) Iterator() data.Iterator {
	return data.NewIndexedIterator(o)
}

////////////////////////////////////////////////////////////////////////////////

type TreeObjects []*TreeObject

var (
	_ data.IndexedAccess = TreeObjects{}
	_ data.Iterable      = TreeObjects{}
)

func (o TreeObjects) Len() int {
	return len(o)
}

func (o TreeObjects) Get(i int) interface{} {
	return o[i]
}

func (o TreeObjects) Iterator() data.Iterator {
	return data.NewIndexedIterator(o)
}
