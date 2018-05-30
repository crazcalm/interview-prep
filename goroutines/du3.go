/*
du2.go takes too long to finish. There's no reason why all the calls to walkDir can't be
done concurrently, thereby exploiting parallelism in the disk system.

The third version of du, creates a new goroutine for each call to walkDir. It uses a
sync.WaitGroup to count the number of calls to walkDir that are still active, and a
closer goroutine to close the filesizes channel when the counter drops to zero.

Since this program creates many thousands of goroutines at its peak, we have to
change dirents to use a counting semaphore to prevent it from opening too many
files at once.
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

//dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //acquire token
	defer func() { <-sema }() //release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

//walkDir recursively walks the file tree rooted at dir
//and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "Show verbose progress messages")

func main() {
	//Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//Traverse the each root of the file tree in parallel
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	//Prints the result periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop //fileSizes was closed. Break label loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}

	}
	printDiskUsage(nfiles, nbytes)
}
