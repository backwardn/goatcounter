// Copyright © 2019 Martin Tournoij <martin@arp242.net>
// This file is part of GoatCounter and published under the terms of the EUPL
// v1.2, which can be found in the LICENSE file or at http://eupl12.zgo.at

package main

import (
	"fmt"
	"os"
)

var usage = map[string]string{
	"": `Usage: goatcounter [command] [flags]

Commands:
help           Show help.
serve          Serve just one domain. This is probably what you want if you're
               looking to self-host GoatCounter. Requires creating a site with
               "create" first.
init           Create a new site and user; only needed for "serve".
migrate        Run database migrations.

Advanced commands:
dev            Run a development server.
saas           Run a "SaaS" production server.
reindex        Re-create the cached statistics (*_stats tables) from the hits.
               This is generally rarely needed and mostly a development tool.

See "help <command>" for more details for the command.`,

	"serve":   "TODO",
	"init":    "TODO",
	"migrate": "TODO",
	"dev":     "TODO",
	"saas":    "TODO",
	"reindex": "TODO",
}

const usage_flags = `
Global flags for all commands:

-domain        Base domain followed by comma and list of static domains. You
               need to have at least one static domain. Must include port
               number if non-standard.
               Default: goatcounter.localhost:8081,static.goatcounter.localhost:8081

-listen        Address to listen on. Default: localhost:8081

-db            Database connection string. Use "sqlite://<dbfile>" for SQLite,
               or "postgres://<connect string>" for PostgreSQL
               Default: sqlite://db/goatcounter.sqlite3

-smtp          SMTP server for sending login emails and errors (if
               -emailerrors is enabled).
               Default is blank, meaning nothing is sent.

-emailerrors   Email errors to this address; requires -smtp. Default: not set.

-stripe        Stripe keys; needed for billing. It needs the secret,
               publishable, and webhook (sk_*, pk_*, whsec_*) keys as
               colon-separated

-debug         Modules to debug, comma-separated or 'all' for all modules.

-plan          Plan for new installations; default: business.

-certdir       Directory to store ACME-generated certificates for custom
               domains. Default: empty.

-cpuprofile    Path to write CPU and memory pprof profiles.
-memprofile

Flags for reindex:

-since         Reindex only statistics since this date instead of all.
                As year-month-day.

Flags for migrate:

-all           Run all migrations that haven't been run yet.

Any positional flags will be interpreted as migration names to run.
`

func die(code int, msg string, args ...interface{}) {
	out := os.Stdout
	if code > 0 {
		out = os.Stderr
	}

	if code > 0 {
		msg = "%s: " + msg
	}

	fmt.Fprintf(out, msg+"\n", append([]interface{}{os.Args[0]}, args...)...)
	os.Exit(code)
}
