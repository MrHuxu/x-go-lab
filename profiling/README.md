# Profiling

    ├── README.md
    ├── scripts
    │   ├── FlameGraph    // util scripts from https://github.com/brendangregg/FlameGraph
    │   ├── ab.sh         // make pressure test requests
    │   ├── diff.sh       // diff two .folded files and generate .svg file
    │   └── gen.sh        // generate .perf, .folded, .svg files for a service
    └── server            // a small web server
        ├── go.mod
        ├── go.sum
        └── main.go

## Usage

### Launch Server

    cd server && go run main.go

### Generate profiling files

Open one console and enter `scripts` directory, then execute `gen.sh` to generate profiling files:

    # $1: host and port of profiling server
    # $2: filename
    # $3: profiling duration

    ./gen.sh http://localhost:6060 cpu1 10

Open another console and execute `ab.sh` multi times to make pressure requests:

    cd scripts && ./ab.sh

After the `gen.sh` is done, `cpu1.perf`, `cpu1.folded` and `cpu1.svg` will be generated.

### Diff profiling files

Open one console and enter `scripts` directory, then execute `diff.sh` to diff profiling files:

    # $1: filename 1
    # $2: filename 2

    ./diff.sh cpu1 cpu2

After the `diff.sh` is done, `diff_cpu1_cpu2.svg` will be generated.