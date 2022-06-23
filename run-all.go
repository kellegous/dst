package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const dockerImageName = "kellegous/dst"

func isInDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}

func hadDockerImage(name string) bool {
	c := exec.Command(
		"docker",
		"inspect",
		name)
	return c.Run() == nil
}

type Flags struct {
	InDocker bool
	Root     string
	Verbose  bool
}

func (f *Flags) Register(fs *flag.FlagSet) {
	fs.BoolVar(
		&f.InDocker,
		"in-docker",
		isInDocker(),
		"run in-docker mode")
	fs.StringVar(
		&f.Root,
		"root",
		"",
		"the project's root directory")
	fs.BoolVar(
		&f.Verbose,
		"verbose",
		false,
		"print out more stuff")
}

func ParseFlags() (*Flags, error) {
	var flags Flags
	flags.Register(flag.CommandLine)
	flag.Parse()

	var err error
	flags.Root, err = getRoot(flags.Root)
	if err != nil {
		return nil, err
	}

	return &flags, nil
}

func buildDockerImage(
	root string,
	name string,
	arch string,
	verbose bool,
) error {
	c := exec.Command(
		"docker",
		"build",
		"-t", name,
		fmt.Sprintf("--build-arg=ARCH=%s", arch),
		".")
	c.Dir = root
	if verbose {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}
	return c.Run()
}

func RunOnHost(flags *Flags) error {
	if err := buildDockerImage(
		flags.Root,
		dockerImageName,
		runtime.GOARCH,
		flags.Verbose,
	); err != nil {
		return err
	}

	c := exec.Command(
		"docker",
		"run",
		"-ti",
		"--rm",
		"-v", fmt.Sprintf("%s:/data", flags.Root),
		"-w", "/data",
		dockerImageName,
		"go", "run", "run-all.go")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = flags.Root
	return c.Run()
}

type Runner struct {
	Name     string
	Commands []*exec.Cmd
}

func (r *Runner) Run() error {
	fmt.Println(r.Name)
	for _, cmd := range r.Commands {
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	fmt.Println()
	return nil
}

func newRunner(name string, cmd ...*exec.Cmd) Runner {
	return Runner{
		Name:     name,
		Commands: cmd,
	}
}

func withOutput(c *exec.Cmd) *exec.Cmd {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func RunInDocker(flags *Flags) error {
	runners := []Runner{
		newRunner(
			"Ruby",
			withOutput(exec.Command("ruby", "main.rb")),
		),
		newRunner(
			"Python",
			withOutput(exec.Command("python3", "main.py")),
		),
		newRunner(
			"JavaScript",
			withOutput(exec.Command("node", "main.js")),
		),
		newRunner(
			"Java",
			exec.Command("javac", "Main.java"),
			withOutput(exec.Command("java", "-classpath", ".", "Main")),
		),
		newRunner(
			"Go",
			withOutput(exec.Command("go", "run", "main.go")),
		),
		newRunner(
			"PHP",
			withOutput(exec.Command("php", "main.php")),
		),
		// newRunner(
		// 	"Rust",
		// 	exec.Command("cargo", "build"),
		// 	withOutput(exec.Command("./target/debug/dst")),
		// ),
		// newRunner(
		// 	"C++",
		// 	exec.Command("make", "clean", "ALL"),
		// 	withOutput(exec.Command("./main")),
		// ),
	}

	for _, runner := range runners {
		if err := runner.Run(); err != nil {
			return err
		}
	}

	return nil
}

func getRoot(fromFlag string) (string, error) {
	if fromFlag != "" {
		return fromFlag, nil
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to determine caller")
	}

	return filepath.Abs(filepath.Dir(file))
}

func main() {
	flags, err := ParseFlags()
	if err != nil {
		log.Panic(err)
	}

	if isInDocker() {
		if err := RunInDocker(flags); err != nil {
			log.Panic(err)
		}
	} else {
		if err := RunOnHost(flags); err != nil {
			log.Panic(err)
		}
	}
}
