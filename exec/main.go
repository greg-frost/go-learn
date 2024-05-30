package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println(" \n[ ВЫПОЛНЕНИЕ ПРОЦЕССОВ ]\n ")

	/* Команда date */

	// Без флагов
	fmt.Println("> date")
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dateOut))

	// С несуществующим флагом
	fmt.Println("> date -x")
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		fmt.Print("Ошибка: ")
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("ошибка выполнения:", err)
		case *exec.ExitError:
			fmt.Println("код завершения =", e.ExitCode())
		default:
			log.Fatal(err)
		}
	}
	fmt.Println()

	/* Команда grep */

	fmt.Println("> grep hello")
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))

	/* Команда ls */

	// Bash
	fmt.Println("> ls -a -l -h")
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(lsOut[:250]))
	fmt.Println("...")
	fmt.Println()

	// Syscall
	fmt.Println("> ls (syscall)")
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		log.Fatal(lookErr)
	}
	execErr := syscall.Exec(
		binary,
		[]string{"ls", "-a", "-l", "-h"},
		os.Environ(),
	)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
