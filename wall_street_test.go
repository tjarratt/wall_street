package wall_street_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"wall_street"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wall Street", func() {
	Describe("Readline", func() {
		It("reads from a pipe and returns a string", func() {
			simulatePipes("The return of the Archons", func(r io.Reader) {
				wall_street.SetReadPipe(r)

				readline := wall_street.Readline("Tonight on The Outer Limits")
				Expect(readline).To(Equal("The return of the Archons"))
			})
		})

		Describe("echoing output back to the terminal", func() {
			It("writes to a pipe", func() {
				output := simulatePipes("Where No Man Has Gone Before", func(r io.Reader) {
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

				output := captureSTDOUT(func() {
					simulatePipes("Where No Man Has Gone Before", func(r io.Reader) {
						wall_street.SetReadPipe(r)
						wall_street.Readline("Tonight on The Outer Limits")
					})
				})

				Expect(output).To(Equal([]string{}))
			})
		})
	})
})

func simulatePipes(input string, block func(r io.Reader)) []string {
	originalPipe := os.Stdin
	reader, writer, err := os.Pipe()
	Expect(err).NotTo(HaveOccurred())

	os.Stdin = reader
	defer func() {
		os.Stdin = originalPipe
	}()

	go func() {
		defer writer.Close()
		writer.Write([]byte(input))
	}()

	return captureSTDOUT(func() { block(reader) })
}

func captureSTDOUT(block func()) []string {
	originalPipe := os.Stdout
	reader, writer, err := os.Pipe()
	Expect(err).ToNot(HaveOccurred())
	os.Stdout = writer
	wall_street.SetWritePipe(writer)

	defer func() {
		os.Stdout = originalPipe
		wall_street.SetWritePipe(originalPipe)
	}()

	block()
	writer.Close()

	var buf bytes.Buffer
	io.Copy(&buf, reader)
	if len(buf.String()) == 0 {
		return []string{}
	}

	return strings.Split(buf.String(), "\n")
}
