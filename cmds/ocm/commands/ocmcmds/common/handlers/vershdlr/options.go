// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package vershdlr

import (
	"github.com/Masterminds/semver/v3"

	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/lookupoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/versionconstraintsoption"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/options"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/utils"
)

type Option interface {
	ApplyToHandler(handler *TypeHandler)
}

type Options []Option

func (o Options) ApplyToHandler(handler *TypeHandler) {
	for _, e := range o {
		e.ApplyToHandler(handler)
	}
}

func OptionsFor(o options.OptionSetProvider) Options {
	var hopts []Option
	if constr := versionconstraintsoption.From(o); constr != nil {
		if len(constr.Constraints) > 0 {
			hopts = append(hopts, WithVersionConstraints(constr.Constraints))
		}
		if constr.Latest {
			hopts = append(hopts, LatestOnly())
		}
	}
	if lookup := lookupoption.From(o); lookup != nil {
		hopts = append(hopts, Resolver(lookup))
	}
	return hopts
}

////////////////////////////////////////////////////////////////////////////////

type constraints struct {
	constraints []*semver.Constraints
}

func (o constraints) ApplyToHandler(handler *TypeHandler) {
	handler.constraints = o.constraints
}

func WithVersionConstraints(c []*semver.Constraints) Option {
	return constraints{c}
}

////////////////////////////////////////////////////////////////////////////////

type latestonly struct {
	flag bool
}

func (o latestonly) ApplyToHandler(handler *TypeHandler) {
	handler.latest = o.flag
}

func LatestOnly(b ...bool) Option {
	return latestonly{utils.OptionalDefaultedBool(true, b...)}
}

////////////////////////////////////////////////////////////////////////////////

type resolver struct {
	resolver ocm.ComponentVersionResolver
}

func (o resolver) ApplyToHandler(handler *TypeHandler) {
	handler.resolver = o.resolver
}

func Resolver(r ocm.ComponentVersionResolver) Option {
	return resolver{r}
}

////////////////////////////////////////////////////////////////////////////////

type repository struct {
	repository ocm.Repository
}

func (o repository) ApplyToHandler(handler *TypeHandler) {
	handler.repobase = o.repository
}

func Repository(r ocm.Repository) Option {
	return repository{r}
}