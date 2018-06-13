//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BOSHBackupRestores returns a BOSHBackupRestoreInformer.
	BOSHBackupRestores() BOSHBackupRestoreInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BOSHBackupRestores returns a BOSHBackupRestoreInformer.
func (v *version) BOSHBackupRestores() BOSHBackupRestoreInformer {
	return &bOSHBackupRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
