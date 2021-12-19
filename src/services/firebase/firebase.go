package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

const (
	SectionKey = "firebase"
)

type Firebase struct {
	App       *firebase.App
	Firestore *firestore.Client
	Settings  *Settings
	Context   *context.Context
}

func Create(settings map[string]interface{}) (app *Firebase, err error) {
	app = &Firebase{}

	if app.Settings, err = readSettings(SectionKey, settings); err != nil {
		return nil, fmt.Errorf("could read the firebase settings: %w", err)
	}

	if err = app.Settings.Credentials.load(); err != nil {
		return nil, fmt.Errorf("could not load credentials: %w", err)
	}

	sa := option.WithCredentialsFile(app.Settings.Credentials.Secret)
	ctx := context.Background()
	app.Context = &ctx
	app.App, err = firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase service: %w", err)
	}

	app.Firestore, err = app.App.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("could read the firebase settings: %w", err)
	}

	return app, nil
}

func (app *Firebase) Dispose() (err error) {
	if err = app.Firestore.Close(); err != nil {
		return fmt.Errorf("could not dispose of firestore: %w", err)
	}

	return
}
