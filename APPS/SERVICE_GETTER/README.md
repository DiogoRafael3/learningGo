This is the first app developed in the course.

This app allows you to get public IPs and server names from urls using the net package.

To run this application either run

go run main.go <COMMAND_NAME> --<FLAG> <WEB_ADRESS>

or

./serv_get <COMMAND_NAME> --<FLAG> <WEB_ADRESS>

after performing a go build.
The difference is ./serv_get skips compile time because that was done in go build.

To check out both commands, just do ./serv_get and they will be output in the command line.
