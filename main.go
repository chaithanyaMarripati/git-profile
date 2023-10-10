package main

import "os"
import "log"

func checkGitProfile(){
  fd,err := os.Open("~/.git-profile");
  defer fd.Close();
  if err!= nil{
    log.Println("creating the git profile spaces");
    err := os.Mkdir(os.ExpandEnv("$HOME/.git-profile"),0755);
    if err !=nil {
      log.Fatalln(err);
    }
  }
}

func main()  {
  // check if the command is run as a sudo user
  uid := os.Getuid()
  if uid!=0 {
    log.Fatalln("Please run this as a sudo user");
  }
  log.Println("sudo user please move on");
  // check if the .git-profile is present, if not create it 
  checkGitProfile();
}
