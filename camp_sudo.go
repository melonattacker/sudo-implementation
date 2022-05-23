package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"bufio"
	"syscall"
	"flag"
)

var SUDOERS_PATH string = "/etc/camp_sudoers"

func usernameToUID(username string, uid *int) bool {
	userInfo, err := user.Lookup(username) // getpwnamを呼んでユーザー名からuidをひく
	if err != nil {
        fmt.Println("no such user", err)
		return false
    }
	*uid, _ = strconv.Atoi(userInfo.Uid)
	return true
} 

func isUserSudoers(uid int) bool {
	fp, err := os.Open(SUDOERS_PATH)
    if err != nil {
        fmt.Println(err)
		return false
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    for scanner.Scan() {
		// 一行ずつ処理
        line := scanner.Text() 
		var target_uid int = 0
		usernameToUID(line, &target_uid)
		if target_uid == uid {
			return true
		}
    }

    if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return false
    }
	return false
}



func main() {
	// コマンドライン引数
	var u = flag.String("u", "root", "string flag")
	var cmd = flag.String("cmd", "", "string flag")
	flag.Parse()

	// 現在のユーザーのidを取得
	ruid := syscall.Getuid()
	euid := syscall.Geteuid()
	fmt.Println(ruid, euid)
	// sudoersに該当のユーザーが存在するかチェック
	if isUserSudoers(ruid) != true {
		fmt.Println("error: No matching entry found in sudoers file %s", SUDOERS_PATH)
		os.Exit(1)
	}
	// ユーザーを指定された場合はそのuidを取得する
	var change int = 0
	if usernameToUID(*u, &change) != true {
		fmt.Println("error: Specified user not found: -u %s", u)
		os.Exit(1)
	}
	// ruid, euid, suidをセットする
	err := syscall.Setresuid(change, change, change)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 与えられたコマンドを実行する
	err = syscall.Exec(*cmd, []string{""}, os.Environ()) // Execはexecve()を実行する
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}