package v1beta1

import (
	"github.com/goharbor/harbor-operator/pkg/convert"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (e *Exporter) ConvertTo(dstRaw conversion.Hub) error {
	return convert.ConverterObject(e).To(dstRaw)
}

func (e *Exporter) ConvertFrom(srcRaw conversion.Hub) error {
	return convert.ConverterObject(e).From(srcRaw)
}
