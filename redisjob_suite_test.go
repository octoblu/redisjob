package redisjob_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRedisjob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redisjob Suite")
}
