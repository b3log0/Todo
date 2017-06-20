package main
import (
	"runtime"
	// "fmt"
	"os/exec"
)

const (
	task_mark1 = " "
	task_mark2 = "*"
	done_mark0 = "-"
	undone_mac = "\u2610" 
	done_mac = "\u2611" 
	undone_win = "[ ]" 
	done_win = "[*]" 
)

func getDoneMark() string{
	switch runtime.GOOS{
	case "windows":
		return done_win
	case "mac"://需要确认具体值
		return done_mac
	}
	return ""
}

func getUnDoneMark() string{
	switch runtime.GOOS{
	case "windows":
		return undone_win
	case "mac"://需要确认具体值
		return undone_mac
	}
	return ""
}

func LocateDirectory(filepath string) string{
	switch runtime.GOOS{
	case "windows":
		return undone_win
	case "mac"://需要确认具体值
		return undone_mac
	}
	return ""
}

func main(){    
    exec.Command("cd D:\\CommonTools\\maven_repo\\").Run()
    // c := exec.Command("cmd", "/C", "cd", "D:/CommonTools/maven_repo/")
    // if err := c.Run(); err != nil { 
    //     fmt.Println("Error: ", err)
    // }   
}