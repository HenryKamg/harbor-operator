package v1beta1

import (
	"github.com/goharbor/harbor-operator/pkg/convert"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (n *NotarySigner) ConvertTo(dstRaw conversion.Hub) error {
	return convert.ConverterObject(n).To(dstRaw)
}

func (n *NotarySigner) ConvertFrom(srcRaw conversion.Hub) error {
	return convert.ConverterObject(n).From(srcRaw)
}
