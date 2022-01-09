# animatch

**animatch** is an anime video file tagging tool.

It uses the provided file name as reference. File names are cleaned, tokenized and normalized 
before comparing them to the (also normalized) list of anime titles provided by AniDB.

The main command is:
```shell
animatch tag [file|directory]
````
which appends an `[anidb-12345]` tag to the filename.

The Plex Hama.bundle plugin is then able to extract such tags on its own and propertly fetch the 
corresponding meta data. Another great tagging tool that might help is FileBot which can have 
either the `--q` flag added in a command line interface environment or instead the Quary Expression 
set in the FileBot Node user interface. 

Any of those two properties may be set to the following value in order to extract the anime id 
and directly query the AniDB database:

```
{(fn =~ /\[anidb-(\d+)\]/)[0][1]}
```

Use the following command to add auto completions for your terminal:

```shell
source < $(animatch completion $(basename $SHELL))
```

## Installation

Requirements:

- [Go compiler](https://go.dev/dl/)


### Installation with the Go compiler:

```shell
go install github.com/jxsl13/animatch
```

### Building from source with the Go compiler

Download the source code with either git or directly from the website as archive 
and unpack it to some directory. Usually the directory has the same name is the 
project, `animatch` in this case.

```shell
cd animatch
go build .
```

if you want to install from the source code to your `GOPATH` directory, you may 
also execute the following command.

```shell
cd animatch
go install .
```

## FileBot renaming script

Initially you may tag your files
```shell
animatch tag [file|directory]
```

FileBot bash script
```shell
filebot -rename input-directory \
--output output-directory \
--action test \
--order Airdate \
--conflict auto \
--db AniDB \
--q "{(fn =~ /\[anidb-(\d+)\]/)[0][1]}" \
-non-strict \
--lang "English" \
--log all \
--format "{plex}"
```

# Requesting new sub commands or features:

You may ask for a new feature in the github issues.


# Potential improvements

- fully featured ranaming of the series title as well as the episode number and name.
- Go templating engine in order to rename files according to all fetched meta data that is passed to the templating engine The resulting name must e.g. not contain any newline characters.
- optional flags that control the threshold values.

