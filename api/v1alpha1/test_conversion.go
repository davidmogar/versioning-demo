package v1alpha1

import (
	"fmt"
	"github.com/davidmogar/versioning-demo/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"strings"
)

func (src *Test) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Test) // This is the hub

	dst.ObjectMeta = src.ObjectMeta

	components := strings.Split(src.Spec.NamespacedName, "/")
	if len(components) != 2 {
		return fmt.Errorf("invalid namespaced name '%s'", src.Spec.NamespacedName)
	}

	dst.Spec.NamespacedName = v1beta1.NamespacedName{
		Name:      components[1],
		Namespace: components[0],
	}

	return nil
}

func (dst *Test) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Test) // This is the hub

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.Foo = "Foo" // Set it here or force a default in v1alpha1 as this field was removed in the new version
	dst.Spec.NamespacedName = fmt.Sprintf("%s/%s", src.Spec.NamespacedName.Namespace, src.Spec.NamespacedName.Name)

	return nil
}
