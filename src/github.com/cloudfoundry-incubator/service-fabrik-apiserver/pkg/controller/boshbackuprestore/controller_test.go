
//TODO copyright header


package boshbackuprestore_test

import (
	"time"

	. "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup/v1alpha1"
	. "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/backup/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("BOSHBackupRestore controller", func() {
	var instance BOSHBackupRestore
	var expectedKey string
	var client BOSHBackupRestoreInterface
	var before chan struct{}
	var after chan struct{}

	BeforeEach(func() {
		instance = BOSHBackupRestore{}
		instance.Name = "instance-1"
		expectedKey = "boshbackuprestore-controller-test-handler/instance-1"
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when creating a new object", func() {
		It("invoke the reconcile method", func() {
			client = cs.BackupV1alpha1().BOSHBackupRestores("boshbackuprestore-controller-test-handler")
			before = make(chan struct{})
			after = make(chan struct{})

			actualKey := ""
			var actualErr error = nil

			// Setup test callbacks to be called when the message is reconciled
			controller.BeforeReconcile = func(key string) {
				defer close(before)
				actualKey = key
			}
			controller.AfterReconcile = func(key string, err error) {
				defer close(after)
				actualKey = key
				actualErr = err
			}

			// Create an instance
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			// Verify reconcile function is called against the correct key
			select {
			case <-before:
				Expect(actualKey).To(Equal(expectedKey))
				Expect(actualErr).ShouldNot(HaveOccurred())
			case <-time.After(time.Second * 2):
				Fail("reconcile never called")
			}

			select {
			case <-after:
				Expect(actualKey).To(Equal(expectedKey))
				Expect(actualErr).ShouldNot(HaveOccurred())
			case <-time.After(time.Second * 2):
				Fail("reconcile never finished")
			}
		})
	})
})
