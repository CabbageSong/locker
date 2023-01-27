package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    fmt.Printf("Running %v \n", os.Args[1:])
    Path := "./xxx"
    must(os.Chdir(Path))
    cmd := exec.Command(os.Args[1])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Chroot:       ".",
        Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
        Unshareflags: syscall.CLONE_NEWNS,
    }
    must(cmd.Run())

    must(syscall.Sethostname([]byte("container0")))
    must(syscall.Mount("proc", "proc", "proc", 0, ""))

    must(cmd.Run())

    must(syscall.Unmount("proc", 0))
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
