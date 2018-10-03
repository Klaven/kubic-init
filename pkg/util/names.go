package util

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// The namespace used when no namespace is provided
	defaultNamespace = metav1.NamespaceSystem
)

func NewNamespacedName(name, namespace string) types.NamespacedName {
	return types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}
}

func NamespacedNameToString(ns types.NamespacedName) string {
	if len(ns.Namespace) > 0 {
		return fmt.Sprintf("%s/%s", ns.Namespace, ns.Name)
	}
	return ns.Name
}

// ParseId parses a Kubernetes resource name as Namespace and Name
func StringToNamespacedName(name string) types.NamespacedName {
	nname := ""
	nnamespace := ""

	res := strings.SplitN(name, "/", 2)
	if len(res) == 2 {
		nname, nnamespace = res[1], res[0]
	} else {
		nname, nnamespace = res[0], defaultNamespace
	}
	return types.NamespacedName{Name: nname, Namespace: nnamespace}
}

type ObjNamespacer interface {
	GetName() string
	GetNamespace() string
}

func NamespacedObjToNamespacedName(obj ObjNamespacer) types.NamespacedName {
	return types.NamespacedName{
		Name:      obj.GetName(),
		Namespace: obj.GetNamespace(),
	}
}

func NamaspacedObjToMeta(obj ObjNamespacer) metav1.ObjectMeta {
	ns := NamespacedObjToNamespacedName(obj)
	return metav1.ObjectMeta{
		Name:      ns.Name,
		Namespace: ns.Namespace,
	}
}

func NamespacedObjToString(obj ObjNamespacer) string {
	if len(obj.GetNamespace()) > 0 {
		return fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName())
	}
	return fmt.Sprintf("%s", obj.GetName())
}
