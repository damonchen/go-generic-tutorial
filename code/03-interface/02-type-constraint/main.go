package main

type File interface {
  Read(p []byte) (n int, err error)
  Write(p []byte) (n int, err error)
  Close() error
}

type ReaderFiler interface {
  File,
  io.ReaderCloser(),
}

func read[T ReaderFiler](f T) (n int, err error) {
  defer f.Close()

  b := [4096]int
  return f.Read(b)
}

func main() {
	fp, err := os.Open("/tmp/a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n, err := read(fp)
	fmt.Println(n, err)
}
