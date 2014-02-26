package wall_street_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestWall_street(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wall_street Suite")
}
