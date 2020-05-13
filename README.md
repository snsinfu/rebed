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

```
rebed [options] [<input>...]

 -b, --binsize <binsize>  Set bin size. [default: 1000]
 -m, --mode <mode>        Set binning mode (mean or sum). [default: sum]
 -g, --genome <genome>    Specify a text file containing chromosome sizes.
 -h, --help               Show this help message and exit.
```

rebed loads BedGraph files from command-line argument or standard input and
outputs binned data to standard output.

The following example rebins input.bdg with 1000-bp bins:

```console
rebed -b 1000 input.bdg
```

`rebed` detects chromosome size from input samples by default. You may
explicitly define chromosome size by supplying a text file containing
chromosome sizes to `-g` option:

```console
$ cat hg19.chrom.sizes
chr1	249250621
chr10	135534747
...
$ rebed -g hg19.chrom.sizes -b 100000 input.bdg
```

## Testing

```console
git clone https://github.com/snsinfu/rebed
cd rebed
make test
```

## License

MIT License.
