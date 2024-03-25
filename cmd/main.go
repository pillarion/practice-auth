package main

import (
	"context"

	"github.com/pillarion/practice-auth/internal/app"
	"github.com/pillarion/practice-auth/internal/core/tools/logger"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.FatalOnError("failed to init app", err)
	}

	err = a.Run()
	if err != nil {
		logger.FatalOnError("failed to run app", err)
	}
}
