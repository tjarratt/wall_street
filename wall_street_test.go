package wall_street_test

import (
	"io"
	"wall_street"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wall Street", func() {
	Describe("Readline", func() {

		It("it reads from stdin and returns a string", func() {
			simulateSTDIN("The return of the Archons", func(r io.Reader) {
				readline := wall_street.Readline("Tonight on The Outer Limits")
				Expect(readline).To(Equal("The return of the Archons"))
			})
		})
	})
})

func simulateSTDIN(input string, block func(r io.Reader)) {
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		writer.Write([]byte(input))
	}()

	block(reader)
}
