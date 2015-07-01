# Go Books App

A play Go web app using Mysql and config file parsing, JSON serialization.

This app was put together from various nice pieces from a number of posts
and docs I found.

## Setup

1. Create a MySQL database.
2. Execute the script(s) in the `db` directory in the database.

## Configuration

See `conf/example.conf.sample`

## Testing

    go test github.com/scottjbarr/gobooksapp

## Start Server (Development)

Assuming you have a `conf/local.conf`

    go run config.go models.go handlers.go server.go -config ./conf/local.conf

## Start Server (Compiled)

    gobooksapp -config ./conf/your-config.conf

## License

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENSE.md](LICENSE.md)
