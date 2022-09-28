package urfave

import (
	"reflect"
	"testing"

	"github.com/go-micro/microwire/cli"
)

const (
	FlagString = "string"
	FlagInt    = "int"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestParse(t *testing.T) {
	myCmd := NewCLI(
		cli.CliName("test"),
		cli.CliVersion("v0.0.1"),
		cli.CliDescription("Test Description"),
		cli.CliUsage("Test Usage"),
	)

	myCmd.Add(cli.Name(FlagString), cli.Default("default string"), cli.EnvVars("STRINGFLAG"), cli.Usage("string flag usage"))
	myCmd.Add(cli.Name(FlagInt), cli.Default(0), cli.EnvVars("INTFLAG"), cli.Usage("int flag usage"))

	err := myCmd.Init(
		[]string{
			"testapp",
			"--string",
			"demo",
			"--int",
			"42",
		},
	)
	expect(t, err, nil)

	expect(t, myCmd.StringValue(FlagString), "demo")
	expect(t, myCmd.IntValue(FlagInt), 42)
}
