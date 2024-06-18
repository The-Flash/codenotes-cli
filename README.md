# CodeNotes CLI

Command-Line Interface for CodeNotes VSCode Extension


# Note Types

1. Folder notes
2. Line notes


## Initializing a project

```sh
codenotes init
```

```sh
codenotes init --directory /path/to/your/project

```

# Notes API


### Adding a folder note

```sh
codenotes add note --message "This is a folder note" --directory /path/to/directory
```

### Adding a file note

```sh
codenotes add note --message "This is a file note" --file /path/to/file
```

### Adding a line note

```sh
codenotes add note --line --message "This is a line note" --line-no 10 --file /path/to/file
```


## Adding a sticky line note

```sh
codenotes add note --message "This is a line note" --line-no 10 --sticky --directory /path/to/file
```


### Deleting a note

```sh
codenotes delete note 01ARZ3NDEKTSV4RRFFQ69G5FAV
```

### Updating a folder note

```sh
codenotes update note 01ARZ3NDEKTSV4RRFFQ69G5FAV --message "This is an updated note"
```

### Getting a single note

```sh
codenotes get note 01ARZ3NDEKTSV4RRFFQ69G5FAV
```


# Reminders API - TODO

## Putting a reminder on a note
