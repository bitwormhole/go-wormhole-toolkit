package autoconfig

import "testing"

func TestRunAutoConfig(t *testing.T) {

	args := &CommandArgs{}

	ctx, err := Run(args)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("autoconfig package dir: " + ctx.PackageDirectory.Path.Path())
}
