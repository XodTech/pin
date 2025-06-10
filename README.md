# Pin - Small CLI Utility For Pinning Notes As Files

## Introduction

Pin is a small CLI utility that helps you add notes as files.
It supports multiple file name aliases which you can configure yourself.

## Getting Started

1.  Install Pin using the [installation instructions](#installation-from-source) below.
2.  Configure custom aliases using the [configuration instructions](#configuration) bellow
3. Once installed and configured, you can start using `pin`!

## Installation from source

### On UNIX-Based OS:

- Move into the directory where you cloned this repository
- Ensure you have the Go compiler installed on your machine. You can download it from the official Go website or from your distribution's official repository if you haven't already. For example:
  - On Ubuntu/Debian: `sudo apt install golang-go`
  - On Red Hat/CentOS: `sudo yum install golang`
  - On macOS (with Homebrew): `brew install go`
  - On Arch Linux and Arch based distributions: `sudo pacman -S go`
- Run the following command: `go build`
- Then, you need to manually move the `pin` executable binary to your `/usr/bin` directory. Run the command: `sudo mv pin /usr/bin/`
<!-- Add instruction how to add binary file to a different place and then add an alias pointing to it -->

### Windows:

* Ensure you have the Go compiler installed on your machine. You can download it from the official Go website if you haven't already
* Run the `install.bat` script. This script will install the necessary dependencies and configure your environment
**Important Note:** Never run `.bat` scripts from untrusted sources, as they can potentially harm your system. Only run scripts from sources you trust, and make sure you understand what the script is doing before executing it.

## Configuration

You can customize the behavior of the `pin` command by adding aliases to the `aliases.conf` file,which can be found:

- On UNIX-Based OS: `~/.config/pin/aliases.conf`
- On Windows: `C:\Users\%USERNAME%\pin\aliases.conf`

### Here's an example of how to do this:

```conf
test = .test
```

With this configuration, when you run the command `pin test Some text`, it will create a file called `.test` with the content `Some text`. If you run the command again, it will append the provided text to the end of the file.

**Important Note:** If the specified key does not exist, the app will default to using the file name `.pin`.

Some key points to keep in mind when configuring custom aliases:

* The format for adding an alias is `alias = filename`
* You can use any valid filename as the value for the alias
* You can add multiple aliases to the `aliases.conf` file to customize the behavior of the `pin` command for different use cases

## Contributing

To contribute to this repository:

* Fork the repository
* Clone the forked repository to your local machine
* Create a new branch for your contribution
* Make changes, commit, and push to your fork
* Create a pull request to the original repository

Please:
* Follow the existing coding style
* Test new features
* Use commit messages that are meaningful and consistent in style with existing ones

## License
Pin is released under the terms of the [MIT License](LICENSE)
