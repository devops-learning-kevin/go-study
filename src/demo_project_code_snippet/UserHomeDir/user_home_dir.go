package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"
	"syscall"
)

// UserHomeDir returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%.
// On Plan 9, it returns the $home environment variable.
func UserHomeDir() (string, error) {
	env, enverr := "HOME", "$HOME"
	switch runtime.GOOS {
	case "windows":
		env, enverr = "USERPROFILE", "%userprofile%"
	case "plan9":
		env, enverr = "home", "$home"
	}
	if v := Getenv1(env); v != "" {
		return v, nil
	}
	// On some geese the home directory is not always defined.
	switch runtime.GOOS {
	case "android":
		return "/sdcard", nil
	case "ios":
		return "/", nil
	}
	return "", errors.New(enverr + " is not defined")
}

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
func Getenv1(key string) string {
	Getenv(key)
	v, _ := syscall.Getenv(key)
	return v
}

// Interface is the interface required of test loggers.
// The os package will invoke the interface's methods to indicate that
// it is inspecting the given environment variables or files.
// Multiple goroutines may call these methods simultaneously.
type Interface interface {
	Getenv(key string)
	Stat(file string)
	Open(file string)
	Chdir(dir string)
}

// logger is the current logger Interface.
// We use an atomic.Value in case test startup
// is racing with goroutines started during init.
// That must not cause a race detector failure,
// although it will still result in limited visibility
// into exactly what those goroutines do.
var logger atomic.Value

// SetLogger sets the test logger implementation for the current process.
// It must be called only once, at process startup.
func SetLogger(impl Interface) {
	if logger.Load() != nil {
		panic("testlog: SetLogger must be called only once")
	}
	logger.Store(&impl)
}

// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
func Logger() Interface {
	impl := logger.Load()
	if impl == nil {
		return nil
	}
	return *impl.(*Interface)
}

// Getenv calls Logger().Getenv, if a logger has been set.
func Getenv(name string) {
	if log := Logger(); log != nil {
		log.Getenv(name)
	}
}

// Open calls Logger().Open, if a logger has been set.
func Open(name string) {
	if log := Logger(); log != nil {
		log.Open(name)
	}
}

// Stat calls Logger().Stat, if a logger has been set.
func Stat(name string) {
	if log := Logger(); log != nil {
		log.Stat(name)
	}
}

func main() {
	path, _ := UserHomeDir()
	fmt.Println("User path = \"" + path + "\"")
}
