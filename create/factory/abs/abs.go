package abs

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write()
}

type IFactory interface {
	getReader() Reader
	getWriter() Writer
}

type GifReader struct {
}

func (pr *GifReader) read() {
	fmt.Println("Gif reader")
}

type GifWriter struct {
}

func (pw *GifWriter) write() {
	fmt.Println("Gif writer")
}

type GifFactory struct {
}

func (pf *GifFactory) getReader() Reader {
	return &GifReader{}
}

func (pf *GifFactory) getWriter() Writer {
	return &GifWriter{}
}

type PngReader struct {
}

func (pr *PngReader) read() {
	fmt.Println("png reader")
}

type PngWriter struct {
}

func (pw *PngWriter) write() {
	fmt.Println("png writer")
}

type PngFactory struct {
}

func (pf *PngFactory) getReader() Reader {
	return &PngReader{}
}

func (pf *PngFactory) getWriter() Writer {
	return &PngWriter{}
}

func Test() {
	factory := &PngFactory{}
	factory.getReader().read()
	factory.getWriter().write()
}
