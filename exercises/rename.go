package main

import "os"

func main() {
  // source and destionation name
  src := "hello.txt"
  dst := "golang.txt"

  // rename file
  os.Rename(src, dst)
}