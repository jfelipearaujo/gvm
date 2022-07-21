# GVM - Go Version Manager for Windows

This application can manage multiple golang installations.

If you already have Go installed, just run the command bellow:

```
go install github.com/jfelipearaujo/gvm@latest
```

## ATTENTION:

If you use OSX or Linux, you SHOULD NOT be using this application.

How to use:

Before anything else, you must download the correct version (x64/x86) for your windows, see the [release](https://github.com/jfelipearaujo/gvm/releases) page.

Download the zip file, extract all the content in some folder.

Example:

```
C:\Users\{userName}\.gvm\bin\
```

When finished, add the installation path into your PATH environment variable.

If all goes well, open a new prompt and type the command below:

```
gvm
```

You should see the following output:

```
Usage: gvm <command>

A Go Lang Version Manager

Flags:
  -h, --help       Show context-sensitive help.
      --version    Print version information and quit

Commands:
  go-path      Set the GOPATH environment variable
  install      Install (or reinstall) a valid version of Go Lang
  list         List installed all versions of Go Lang
  uninstall    Uninstall a version of Go Lang
  use          Use an installed version of Go Lang

Run "gvm <command> --help" for more information on a command.

gvm: error: expected one of "go-path",  "install",  "list",  "uninstall",  "use"
exit status 1
```

## Install a version

```
gvm install <version>
```

## Uninstall a version

```
gvm uninstall <version>
```

## List all versions

```
gvm list
```

## Use a version

```
gvm use <version>
```

## Setup GOPATH environment variable

This will define the GOPATH as the default one (C:\{userName}\go)

```
gvm go-path
```

Or

```
gvm go-path "C:\MyFolder\go"
```

If the selected folder does not have the necessary sub-folders (bin, pkg and src) they will be created.
