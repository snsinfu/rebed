rebed
=====

rebed is a command-line utility for binning BedGraph data track.

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [License](#license)

## Installation

For linux, download a single-binary release and put it in your bin directory:

```console
curl -Lo ~/bin/rebed https://github.com/snsinfu/rebed/releases/download/v0.1/rebed-linux-x86_64
```

Or build your own binary:

```console
go get https://github.com/snsinfu/rebed
cd ${GOPATH}/src/github.com/snsinfu/rebed
make
```

## Usage

rebed loads a BedGraph file specified via command-line argument or standard
input and outputs binned data to standard output. The bin size can be specified
via an -b/--binsize option.

The following example rebins input.bdg with 1000-bp bins:

```console
rebed -b 1000 input.bdg
```

## Testing

```console
git clone https://github.com/snsinfu/rebed
cd rebed
make test
```

## License

MIT License.
