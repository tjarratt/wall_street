package wall_street_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"wall_street"
	. "wall_street/testhelpers"
)

var _ = Describe("Wall Street", func() {
	Describe("Readline", func() {
		It("reads from a pipe and returns a string", func() {
			SimulatePipes("The return of the Archons", func(r io.Reader) {
				wall_street.SetReadPipe(r)

				readline := wall_street.Readline("Tonight on The Outer Limits")
				Expect(readline).To(Equal("The return of the Archons"))
			})
		})

		Describe("echoing output back to the terminal", func() {
			It("writes to a pipe", func() {
				output := SimulatePipes("Where No Man Has Gone Before", func(r io.Reader) {
					wall_street.SetReadPipe(r)
					wall_street.Readline("Tonight on The Outer Limits")
				})

				Expect(len(output)).To(Equal(1))

				out := output[0]
				Expect(out).To(Equal("Where No Man Has Gone Before"))
			})

			It("can be suppressed", func() {
				wall_street.DisableEcho()
				defer func() {
					wall_street.EnableEcho()
				}()

				output := CaptureSTDOUT(func() {
					SimulatePipes("Where No Man Has Gone Before", func(r io.Reader) {
						wall_street.SetReadPipe(r)
						wall_street.Readline("Tonight on The Outer Limits")
					})
				})

				Expect(output).To(Equal([]string{}))
			})
		})
	})
})
