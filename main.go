package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
func fdGitProfile() *os.File{
  fd,err := os.Open(os.ExpandEnv("$HOME/.git-profile"));
  if err == nil {
    return fd;
  }else if errors.Is(err, os.ErrNotExist){
    log.Println(".git-profile folder doesn't exists, creating the folder");
    err:= os.Mkdir(os.ExpandEnv("$HOME/.git-profile"),0755);
    if err!=nil {
      log.Println("error creating the .git-profile directory");
      log.Fatal(err)
    }
  }else{
    log.Print(err);
    log.Fatalln("Error dealing with $HOME/.git-profile directory")
  }
  return fd;
}


func checkGitProfile(){
  fd := fdGitProfile();
  defer fd.Close();
}


func listProfiles() {
  fd:= fdGitProfile();
  defer fd.Close();
  // return only the first 100 profiles 
  files,err:= fd.Readdirnames(100);
  if err!=nil {
    log.Fatalln("Error reading files in .git-profile directory");
  }
  fmt.Println("Available profiles are ");
  for _,file := range files {
    fmt.Println(file)
  }
}

func addProfile(){
  fd:= fdGitProfile();
  defer fd.Close();

  username:= "";
  email:= "";
  fmt.Println("Enter the username of the account");
  fmt.Scanln(&username);
  fmt.Println("Enter the email of the account");
  fmt.Scanln(&email);

  fmt.Println("username and email entered are", username, email);
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
  args := os.Args[1:]
  if len(args) == 0 {
    log.Fatalln("Enter a command, available are list , add ,remove");
  }
  switch args[0]{
    case "use":
      log.Fatalln("Use isn't supported yet");
    case "list":
      listProfiles();
    case "add":
      addProfile();
    case "remove":
      log.Fatalln("Remove isn't supported yet");
  } 

}
