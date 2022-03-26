# go-workspaces

## Working with replace

Clone the golang.org/x/sync into ./local folder

```
git clone https://github.com/golang/sync
```

## Working with workspaces in Golang 1.18

The better organization of the code is to have the workspace file in a root directory and the other modules in subdirectories.

1. From the root of the workspace

```
go work init ./test
```

This creates the file go.work and add the line use ./test

```
go work use ./sync
```

2. Running the test program

Now we can run the go run command outside the module, in the workspace, specifying the module name:

```
go run test1.18
```

or the package location

```
go run ./test
```

3. If the go.work doesn't contain use ./sync, the test program will use the sync package installed in GOPATH

4. We can build the test program from the module directory and it will be compiled in the workspace context, because the tool search go.work file in the parent directories.
