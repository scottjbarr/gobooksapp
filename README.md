# Go Books App

A play Go web app using Mysql and config file parsing, DB access, and a the
basics of a simple HTTP app.

Pieces of this were found all over the place.

## Setup

1. Create a MySQL database.
2. Execute the script(s) in the `db` directory in the database.

## Configuration

See `conf/example.conf.sample`

## Starting The Server (Development)

Assuming you have a `conf/local.conf`

    go run config.go models.go handlers.go server.go --config ./conf/local.conf

## License

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENSE.md](LICENSE.md)
