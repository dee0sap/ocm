package valuemergehandler

import (
	"github.com/open-component-model/ocm/pkg/contexts/ocm/cpi"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/valuemergehandler/hpi"
)

func Merge(ctx cpi.Context, m *Specification, hint hpi.Hint, local Value, inbound *Value) (bool, error) {
	return hpi.Merge(ctx, m, hint, local, inbound)
}

func LabelHint(name string, optversion ...string) hpi.Hint {
	return hpi.LabelHint(name, optversion...)
}
