// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"net/http"
	"time"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	Webhook string `envconfig:"PLUGIN_Webhook"`
	Secret  string `envconfig:"PLUGIN_SECRET"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// write code here
	now := time.Now()
	sign, err := genSign(args.Secret, now.Unix())
	if err != nil {
		return err
	}

	body, err := NewBodyBuffer(now.Unix(), &sign, args)
	if err != nil {
		return err
	}

	_, err = http.Post(args.Webhook, "application/json", body)
	return err
}
