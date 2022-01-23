// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package oci

import (
	"context"

	"github.com/gardener/ocm/pkg/oci/core"
	_ "github.com/gardener/ocm/pkg/oci/repositories"
)

const KIND_OCIARTEFACT = core.KIND_OCIARTEFACT

type Context = core.Context
type Repository = core.Repository
type RepositoryTypeScheme = core.RepositoryTypeScheme
type RepositorySpec = core.RepositorySpec
type GenericRepositorySpec = core.GenericRepositorySpec
type ArtefactAccess = core.ArtefactAccess
type ArtefactComposer = core.ArtefactComposer
type DataAccess = core.DataAccess

var DefaultContext = core.DefaultContext

func ForContext(ctx context.Context) Context {
	return core.ForContext(ctx)
}

func NewDefaultContext(ctx context.Context) Context {
	return core.NewContext(ctx, core.DefaultRepositoryTypeScheme)
}

func NewContext(ctx context.Context, reposcheme core.RepositoryTypeScheme) Context {
	return core.NewContext(ctx, reposcheme)
}