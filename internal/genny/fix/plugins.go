package fix

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gobuffalo/cli/internal/genny/plugins/install"

	cmdPlugins "github.com/gobuffalo/cli/internal/cmd/plugins"
	"github.com/gobuffalo/cli/internal/plugins"
	"github.com/gobuffalo/cli/internal/plugins/plugdeps"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/meta"
)

// CleanPluginCache cleans the plugins cache folder by removing it
func CleanPluginCache(opts *Options) ([]string, error) {
	fmt.Println("~~~ Cleaning plugins cache ~~~")
	os.RemoveAll(plugins.CachePath)
	return nil, nil
}

// ReinstallPlugins installs latest versions of the plugins
func ReinstallPlugins(opts *Options) ([]string, error) {
	plugs, err := plugdeps.List(opts.App)
	if err != nil && !errors.Is(err, plugdeps.ErrMissingConfig) {
		return nil, err
	}

	run := genny.WetRunner(context.Background())
	gg, err := install.New(&install.Options{
		App:     opts.App,
		Plugins: plugs.List(),
	})
	if err != nil {
		return nil, err
	}

	run.WithGroup(gg)

	fmt.Println("~~~ Reinstalling plugins ~~~")
	return nil, run.Run()
}

// RemoveOldPlugins removes old and deprecated plugins
func RemoveOldPlugins(opts *Options) ([]string, error) {
	fmt.Println("~~~ Removing old plugins ~~~")

	run := genny.WetRunner(context.Background())
	app := meta.New(".")
	plugs, err := plugdeps.List(app)
	if err != nil && !errors.Is(err, plugdeps.ErrMissingConfig) {
		return nil, err
	}

	plugs.Remove(plugdeps.Plugin{
		Binary: "buffalo-pop",
	})

	fmt.Println("~~~ Removing github.com/gobuffalo/buffalo-pop plugin ~~~")
	run.WithRun(cmdPlugins.NewEncodePluginsRunner(app, plugs))
	return nil, run.Run()
}