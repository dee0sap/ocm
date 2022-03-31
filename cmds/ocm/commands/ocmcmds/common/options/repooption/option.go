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

package repooption

import (
	"github.com/gardener/ocm/cmds/ocm/clictx"
	"github.com/gardener/ocm/cmds/ocm/commands/ocmcmds/common"
	"github.com/gardener/ocm/cmds/ocm/pkg/output"
	"github.com/gardener/ocm/pkg/oci"
	"github.com/gardener/ocm/pkg/runtime"

	"github.com/gardener/ocm/pkg/ocm"
	"github.com/spf13/pflag"
)

func From(o *output.Options) *Option {
	v := o.GetOptions((*Option)(nil))
	if v == nil {
		return nil
	}
	return v.(*Option)
}

type Option struct {
	Spec       string
	Repository ocm.Repository
}

var _ common.OptionCompleter = (*Option)(nil)

func (o *Option) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.Spec, "repo", "r", "", "repository name or spec")
}

func (o *Option) Complete(ctx clictx.Context) error {
	return nil
}

func (o *Option) CompleteWithSession(octx clictx.OCM, session ocm.Session) error {
	if o.Repository == nil {
		r, err := o.GetRepository(octx, session)
		if err != nil {
			return err
		}
		o.Repository = r
	}
	return nil
}

func (o *Option) GetRepository(ctx clictx.OCM, session ocm.Session) (ocm.Repository, error) {
	if o.Spec != "" {
		r, _, err := session.DetermineRepository(ctx.Context(), o.Spec, ctx.GetAlias)
		return r, err
	}
	return nil, nil
}

func (o *Option) Usage() string {
	s := `
If the <code>--repo</code> option is specified, the given names are interpreted
relative to the specified repository using the syntax

<center><code>&lt;component>[:&lt;version>]</code></center>

If no <code>--repo</code> option is specified the given names are interpreted 
as located OCM component version references:

<center><code>[&lt;repo type>::]&lt;host>[:&lt;port>][/&lt;base path>]//&lt;component>[:&lt;version>]</code></center>

Additionally there is a variant to denote common transport archives
and general repository specifications

<center><code>[&lt;repo type>::]&lt;filepath>|&lt;spec json>[//&lt;component>[:&lt;version>]]</code></center>

The <code>--repo</code> option takes an OCM repository specification:

<center><code>[&lt;repo type>::]&lt;configured name>|&lt;file path>|&lt;spec json></code></center>

For the *Common Transport Format* the types <code>directory</code>,
<code>tar</code> or <code>tgz</code> is possible.

Using the JSON variant any repository type supported by the 
linked library can be used:

Dedicated OCM repository types:
`

	types := runtime.KindNames(ocm.DefaultContext().RepositoryTypes())
	for _, t := range types {
		s += "- `" + t + "`\n"
	}

	s += `
OCI Repository types (using standard component repository to OCI mapping):
`
	types = runtime.KindNames(oci.DefaultContext().RepositoryTypes())
	for _, t := range types {
		s += "- `" + t + "`\n"
	}
	return s
}
