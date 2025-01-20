# Version Upgrader

Simple CLI to upgrade versions according to selectable options.

I made this solution because sometimes in my work, we commit changes and forget to update the version, being forced to create a new commit just for this.

So, that's a simple way to avoid this mistake (check my <a href="#use-case">use case</a>).

![{0CB440E0-DDFC-40A7-8F3F-88E478C178FC}](https://github.com/user-attachments/assets/3a21dc1b-ef5f-4ab9-be8c-118f686572df)


## Usage

Clone the application:

```bash
git clone https://github.com/brunopstephan/version-upgrader.git
```

Install dependencies:

```bash
go mod download
```

Build depending for your ARCH and OS using ENV's:

Linux:
```bash
env GOOS="linux" GOARCH="amd64" go build -o "upgrader" cmd/app/main.go 
```

Windows:
```bash
env GOOS="windows" GOARCH="amd64" go build -o "upgrader.exe" cmd/app/main.go 
```

Build using Makefile:

```bash
make all
```

Now, just execute the binary passing the flags

- file: mandatory, path to the file to be updated.
- version_path: optional, path for the version number inside the JSON file, it must be passed like "foo.bar.version_number", default is "version"

<h2 id="use-case">Use case</h2>

My target application is an Google Chrome Extension code in Node.js, so i just use [lefthook](https://github.com/evilmartians/lefthook) to set the pre-commit hook and configured to call a node script that's trigger the binary.



