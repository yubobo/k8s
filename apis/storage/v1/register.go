package v1

import "github.com/ericchiang/k8s"

func init() {
	k8s.Register("storage.k8s.io", "v1", "storageclasses", false, &StorageClass{})

	k8s.RegisterList("storage.k8s.io", "v1", "storageclasses", false, &StorageClassList{})
}
