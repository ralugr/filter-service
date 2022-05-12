# Filter Service
This is the repository for the message filter service

- Build with Go version 1.18
- Uses [chi router](github.com/go-chi/chi)
- Uses [sqlite](github.com/mattn/go-sqlite3)

## Steps to download
#### Cloning with submodules
>git clone --recurse-submodules https://github.com/ralugr/filter-service
#### If you already cloned the repo without submodules
> git submodule update --init --recursive

## Pulling upstream changes from the project remote
By default, the git pull command recursively fetches submodules changes, however, it does not update the submodules.
> git pull

To finalize the update, you need to run git submodule update
> git submodule update --init --recursive




brew install sqlite3
brew install sqlite-utils
go get github.com/mattn/go-sqlite3



Further improvements:
- Use a logging framework (if needed)
- Control log level
- 
