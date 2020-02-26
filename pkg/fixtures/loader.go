package fixtures

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/mon"
)

type FixtureSet struct {
	Enabled  bool
	Writer   FixtureWriterFactory
	Fixtures []interface{}
}

type FixtureLoader struct {
	fixtureSets []*FixtureSet
}

func NewFixtureLoader(fixtureSets []*FixtureSet) *FixtureLoader {
	return &FixtureLoader{
		fixtureSets: fixtureSets,
	}
}

func (f *FixtureLoader) Load(config cfg.Config, logger mon.Logger) error {
	logger = logger.WithChannel("fixture_loader")

	if !config.IsSet("fixture_loader_enabled") {
		logger.Info("fixture loader is not configured")
		return nil
	}

	if !config.GetBool("fixture_loader_enabled") {
		logger.Info("fixture loader is not enabled")
		return nil
	}

	for _, fs := range f.fixtureSets {

		if !fs.Enabled {
			logger.Info("skipping disabled fixture set")
			continue
		}

		if fs.Writer == nil {
			return fmt.Errorf("fixture set is missing a writer")
		}

		writer := fs.Writer(config, logger)
		err := writer.WriteFixtures(fs)

		if err != nil {
			return fmt.Errorf("error during loading of fixture set: %w", err)
		}
	}

	return nil
}
