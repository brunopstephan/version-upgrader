# Version Upgrader

Simple CLI to upgrade versions accordint to selectable options.

![{F9ACFE9C-FF1F-4B0C-9F3A-8E78DB83A515}](https://github.com/user-attachments/assets/6fe27c67-d730-4e1b-9410-d3800ca9bf09)

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

Now, just execute the binary inside passing the flags

- file: mandatory, path to the file to be updated.
- version_path, optional, path for the version number inside the JSON file, it must be passed like "foo.bar.version_number" where the "version_number", default is "version"
