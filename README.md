# Go Social API
This project is based on [OtavioGallego/DevBook](https://github.com/OtavioGallego/DevBook) and serves functionalities for a social network, such as users and posts management, as well as a user following system.

## Usage

You should make a copy of the `example.env` file, rename it into `.env` and provide the required info, as detailed below:

- `DB_USER` is the username for your database;
- `DB_PASSWORD` is the password for your database;
- `DB_NAME` is the name for your database;
- `API_PORT` is the port on which the project should run;
- `SECRET_KEY` is the seed for password hash generation.

You shold also run `sql/sql.sql` on your database in order to create the necessary environment. Optionally, running the `sql/data.sql` file will populate the database with sample data for testing.

## Dependencies

Download all project dependencies with the `go get` command.

## Compiling and running

You may compile the project by running `go build` from the root directory, which will generate a binary file called `api`.

It is also possible to invoke Go and run the project with `go run main.go`.
