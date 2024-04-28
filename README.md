# CodeNotes CLI

Command-Line Interface for CodeNotes VSCode Extension

# Usage

```
cargo run
```

# Note Types

1. Folder notes
2. Line notes


## Initializing a project

```sh
codenotes init
```

```sh
codenotes init --path /path/to/your/project

```

# Notes API

## Folder Notes API

### Adding a folder note

```sh
codenotes notes folder add --note "This is a folder note" --directory /path/to/directory
```


### Deleting a folder note

```sh
codenotes notes folder delete --note-id 01ARZ3NDEKTSV4RRFFQ69G5FAV
```

### Updating a  folder note

```sh
codenotes notes folder update --note-id 01ARZ3NDEKTSV4RRFFQ69G5FAV --note "This is an updated note"
```

### Getting a single note

```sh
codenotes notes folder get --note-id 01ARZ3NDEKTSV4RRFFQ69G5FAV
```

## Line Notes API

### Adding a line note

```sh
codenotes notes line add --note "This is a line note" --number 10 --directory /path/to/file
```

## Adding a sticky line note

```sh
codenotes notes line add --note "This is a line note" --number 10 --sticky --directory /path/to/file
```

# Reminders API

## Putting a reminder on a folder note

```
```

## Putting a reminder on a line note