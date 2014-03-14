package testhelpers

import(
	. "github.com/onsi/gomega"
	"bytes"
  "io"
	"os"
  "strings"
  "wall_street"
)

func SimulatePipes(input string, block func(r io.Reader)) []string {
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

	return CaptureSTDOUT(func() { block(reader) })
}

func CaptureSTDOUT(block func()) []string {
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
