package urfave

import (
	"reflect"
	"testing"

	"github.com/go-micro/microwire/v5/cli"
)

const (
	FlagString = "string"
	FlagInt    = "int"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestParse(t *testing.T) {
	myCli := NewCLI(
		cli.CliName("test"),
		cli.CliVersion("v0.0.1"),
		cli.CliDescription("Test Description"),
		cli.CliUsage("Test Usage"),
	)

	var destString string
	var destInt int
	err := myCli.Add(
		cli.Name(FlagString),
		cli.Default("micro!1!1"),
		cli.EnvVars("STRINGFLAG"),
		cli.Usage("string flag usage"),
		cli.Destination(&destString),
	)
	expect(t, err, nil)

	err = myCli.Add(
		cli.Name(FlagInt),
		cli.EnvVars("INTFLAG"),
		cli.Usage("int flag usage"),
		cli.Destination(&destInt),
	)
	expect(t, err, nil)

	err = myCli.Parse(
		[]string{
			"testapp",
			"--string",
			"demo",
			"--int",
			"42",
		},
	)
	expect(t, err, nil)

	expect(t, destString, "demo")
	expect(t, destInt, 42)
}
