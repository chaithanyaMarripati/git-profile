package main

import "fmt"
import "os"

func main()  {
  // check if the command is run as a sudo user
  uid := os.Getuid()
  if uid!=0 {
    fmt.Println("Please run this as a sudo user");
    os.Exit(1)
  }
  fmt.Println("Running as sudo user, moving on");
}
