package v1

import (
	"context"
	"runtime"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/opt"
	"github.com/tetratelabs/wazero/internal/integration_test/spectest"
	"github.com/tetratelabs/wazero/internal/platform"
)

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestWazevo(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	c := opt.NewRuntimeConfigOptimizingCompiler().WithCoreFeatures(api.CoreFeaturesV1)
	spectest.Run(t, Testcases, context.Background(), c)
}
