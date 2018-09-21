rebed
=====

rebed is a command-line utility for binning BedGraph data track.

- [Usage](#usage)
- [Testing](#testing)
- [License](#license)

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
