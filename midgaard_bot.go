/*
midgaard_bot, a Telegram bot which sets a bridge to Midgaard Merc MUD
Copyright (C) 2017 by Javier Sancho Fernandez <jsf at jsancho dot org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"log"

	"github.com/jessevdk/go-flags"
)

var config struct {
	Token string `short:"t" long:"token" description:"Telegram API Token" required:"true"`
}

func main() {
	_, err := flags.Parse(&config)
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = initSessions(ctx)
	if err != nil {
		log.Panic(err)
	}

	err = initTelegramWorkers(config.Token, ctx)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			break
		}
	}
}
