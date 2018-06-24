This is a simple demonstration of using TDD to create DB capability with Golang.

It uses Sqlite as default DB because it can be used as an in-memory DB, making TDD easy, or simple file-based.

The unit tests demonstrate full CRUD operations and the use of json marshalling, which is how the array of recipients in the email TO field are persisted.

There is a main() in email.go that covers less functionality; but, with one line, allows you to toggle between an in-memory and a file-based demonstration.  You can uncomment the line and put in the file location.  If output to file, you'll see it add one row each time you run it, then go back to one row after you've reached 5 rows, verifying it is persisting to a file.

You'll need to install two dependencies to run in your IDE.

## Dependencies

```
go get github.com/stretchr/testify
go get github.com/mattn/go-sqlite3
```

