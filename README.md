# mkdeskf

A simple CLI tool that creates `.desktop` files for Linux applications and links without having to write them by hand.

## What is a `.desktop` file?

`.desktop` files tell your Linux desktop environment about an application — its name, icon, executable, and how to launch it. They live in `~/.local/share/applications/` and show up in your app launcher.

## Installation

### Using `go install` (recommended)

Requires Go 1.18 or later.

```sh
go install github.com/your-username/mkdeskf@latest
```

The binary is placed in `~/go/bin/`. Make sure that directory is on your `$PATH`:

```sh
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Build from source

```sh
git clone https://github.com/your-username/mkdeskf
cd mkdeskf
go build -o mkdeskf .
mv mkdeskf ~/.local/bin/
```

## Usage

Run the tool and follow the prompts:

```sh
mkdeskf
```

You will be asked to choose a type:

```
1: Application
2: Link
```

### Application

Creates a `.desktop` file that launches an executable.

| Prompt | Description |
|---|---|
| Are your files in this directory? | If yes, the current directory is used automatically |
| Directory | Path to the app directory. Use `$wd` for the current directory |
| Executable file name | The binary or script to run (relative to directory) |
| Application Name | The name shown in your launcher |
| Icon | Path to the icon file. Use `$wd` for the current directory. Optional |
| Version | App version (e.g. `1.0`). Optional |
| Comment | Short description shown as a tooltip. Optional |
| Categories | Semicolon-separated XDG categories (e.g. `Utility;`). Optional |

### Link

Creates a `.desktop` file that opens a URL in your default browser. It shows up in your app launcher alongside your other applications.

| Prompt | Description |
|---|---|
| Link Name | The name shown in your launcher |
| Icon | Icon name or path. Optional |
| Version | Version string. Optional |
| Comment | Short description shown as a tooltip. Optional |
| URL | The URL to open. `https://` is added automatically if you omit it |

### `$wd` shorthand

In any path field (directory, icon, save location), you can type `$wd` instead of the full path to the current working directory.

For example, if you run `mkdeskf` from `/home/user/apps/mytool`, entering `$wd/icon.png` becomes `/home/user/apps/mytool/icon.png`.

### Output location

By default, the `.desktop` file is saved to `~/.local/share/applications/`. You can enter a custom path when prompted, and `$wd` works there too.

## Examples

### Application

```
$ mkdeskf
1: Application
2: Link
Your choice: 1
Are your files in this directory? [Y/N]: y
Current directory: /home/user/apps/mytool
Executable file name: mytool
Application Name: My Tool
Icon ($wd = current directory, optional): $wd/icon.png
Version (optional): 1.0
Comment (optional): A useful tool
Categories (optional, e.g. Utility;): Utility;

Are these correct?
Application Name: My Tool
Executable: /home/user/apps/mytool/mytool
Icon: /home/user/apps/mytool/icon.png
Version: 1.0
Comment: A useful tool
Categories: Utility;
[Y/N]: y
Default location: /home/user/.local/share/applications/
Proceed with this location? [Y/N]: y
Done!
```

### Link

```
$ mkdeskf
1: Application
2: Link
Your choice: 2
Link Name: GitHub
Icon (optional):
Version (optional):
Comment (optional): Open GitHub in the browser
URL: github.com

Are these correct?
Link Name: GitHub
Comment: Open GitHub in the browser
URL: https://github.com
[Y/N]: y
Default location: /home/user/.local/share/applications/
Proceed with this location? [Y/N]: y
Done!
```

## License

See [LICENSE](LICENSE).
