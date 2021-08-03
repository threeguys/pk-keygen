# pk-keygen
##### Simple command-line key-generator for Perkeep (https://perkeep.org)

DISCLAIMER: This repo is not affiliated with [Perkeep](https://perkeep.org) or the [Perkeep Repo](https://github.com/perkeep/perkeep) in any way, just a fan of the software.

Perkeep generates a secure key on startup if none is given. I fiddled with GnuPG for 
a bit and was unable to generate a key that perkeep would accept. Instead of digging into
it too deeply, I discovered the ease with which the perkeep source could be imported
and generate a key on the command line that way. All I'm really interested in doing is
to be able to generate keys ahead of time, and couple them with low-level config of
my own. There are a couple of issues around key generation / conformance, however I'm
too lazy to track those down. If this repo becomes obsolete with a couple of lines 
of bash script and GnuPG, please let me know!

## Building

Simply `go get` the project and build it.

```
go get github.com/threeguys/pk-keygen
go build github.com/threeguys/pk-keygen
```

## Usage

The program generates a `.gpg` file and `.json` file, containing the identity and
secret ring settings. You will need to take the settings in the JSON file and incorporate
them into your `server-config.json` in your perkeep installation.

```
$ pk-keygen -help
Usage of pk-keygen:
  -config string
    	Configuration information about key (default "secring-config.json")
  -secret string
    	Secret key ring file to generate (default "secring.gpg")
```

#### Example

Create a key with `my-secring.gpg` and `my-config.json` filenames

```
$ pk-keygen -config my-config.json -secret my-secring.gpg
Generated key to my-secring.gpg
Generated config to my-config.json
  Key Identifier: CA02F5397BAD0BA0
$ cat my-config.json
{
  "identity": "CA02F5397BAD0BA0",
  "identitySecretRing": "my-secring.gpg"
}
```

## Licensing
I'm distributing this under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0), given
that is how Perkeep is licensed. There's not really much code here, all the important stuff is
in the perkeep libs.

## Contributing
I don't really want to be maintaining this repo, I would much rather use GnuPG or other
common tools, however I will continue to keep it up-to-date with Perkeep releases, as long as
I am still using Perkeep in my personal and open source projects. Feel free to submit PRs or
Issues, however they will not be a top priority.
