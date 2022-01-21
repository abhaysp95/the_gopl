 package main

 import (
	 "bufio"
	 "fmt"
	 "os"
	 "strconv"
	 "time"
 )

 func main() {
	 input := bufio.NewScanner(os.Stdin)
	 for input.Scan() {
		 /** This is just a demo, for what ? See the not below for this if
		 * block */
		 if in := input.Text(); in != "q" {
			 x, err := strconv.ParseInt(in, 0, 64)
			 if err != nil {
				 fmt.Fprintln(os.Stderr, err)
			 } else {
				 start := time.Now()
				 fmt.Println(popCount(&x), time.Since(start))
			 }
		 }
	 }
 }

func popCount(x *int64) int {
	c := 0
	for *x != 0 {
		if *x & 1 == 1 {
			c++
		}
		*x >>= 1
	}
	return c
}

/** Note:
Thus it is often necessary to declare f before the condition so that it is accessible after :

	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	f.ReadByte()
	f.Close()

You may be tempted to avoid declaring f and err in the outer block by moving the calls to ReadByte and Close inside an else block:

	if f, err := os.Open(fname); err != nil {
		return err
	} else {
		// f and err are visible here too
		f.ReadByte()
		f.Close()
	}

but normal practice in Go is to deal with the error in the if block and then return, so that the successful execution path is not intended
*/
