# Wabeltools CLI

`wabeltools` is a command line interface for the Wabeltools service. This utility allows users to interact with the
Wabeltools service directly from their terminal.

## Available Platforms

We provide pre-compiled binaries for the following platforms:

- Linux (x86_64, ARM)
- macOS (x86_64, ARM)
- Windows (x86_64)

If your platform is not listed, you may be able to compile the code from source by cloning the repository and
running `go build`.

## Installation

You can install the `wabeltools` CLI on your system by downloading the appropriate binary for your platform from the
following links,
and then moving the binary to a directory in your system's PATH.

Please follow the instructions below for your platform:

### For Linux:

Get the binary:

```bash
curl -LO https://github.com/mx79/wabeltools-cli/raw/main/wabeltools-linux-amd64
```

Add permission to the binary and add it to your PATH:

```bash
chmod +x ./wabeltools-linux-amd64
sudo mv ./wabeltools-linux-amd64 /usr/local/bin/wabeltools
```

### For Mac:

Get the binary:

```bash
curl -LO https://github.com/mx79/wabeltools-cli/raw/main/wabeltools-darwin-amd64
```

Add permission to the binary and add it to your PATH:

```bash
chmod +x ./wabeltools-darwin-amd64
sudo mv ./wabeltools-darwin-amd64 /usr/local/bin/wabeltools
```

### For Windows

Download the appropriate binary for Windows from the following link:

https://github.com/mx79/wabeltools-cli/raw/main/wabeltools-windows-amd64.exe

Add the directory containing the downloaded binary to your system's PATH and you should be good to go.

## Usage

You can use the Wabeltools CLI by running wabeltools in your terminal, followed by the various commands and flags
available. For instance:

```bash
wabeltools init "your-api-key"
```

To get your remaining tokens of the day:

```bash
wabeltools tokens
```

To get the list of cost of the different services:

```bash
wabeltools costs
```

To get the list of available services:

```bash
wabeltools services
```

To get the list of available image processing methods:

```bash
wabeltools image 
```

For example, you can compress or resize one our multiple images locally or remotely with the following commands:

```bash
wabeltools image local --compress "path/to/image.jpg" --output="path/to/output"
wabeltools image remote --resize 0.5 "https://example.com/image.jpg"
wabeltools image local --resize 0.5 "path/to/image.jpg" "path/to/image2.jpg"
wabeltools image remote --compress --quality="80" "https://example.com/image.jpg" "https://example.com/image2.jpg" --output="path/to/output"
```

To get the list of available text processing methods:

```bash
wabeltools nlp 
```

For example, you can apply different text processing methods to a text with the following commands:

```bash
wabeltools nlp stopwords --lang="fr" "C'est un text en français" >> "path/to/output"
wabeltools nlp sentiment "I love you"
wabeltools nlp stemming "I want to stem this sentence" > "path/to/output"
wabeltools nlp segmenter "I want to segment this text"
wabeltools nlp ner "I want to extract entities from this text"
wabeltools nlp pos-tagging "I want to tag this text"
wabeltools nlp rake "I need the rake score of this sentence"
wabeltools nlp wer "I would like to get" "the word error rate of this two texts"
```

## License

Wabeltools CLI is licensed under the MIT License.

## Support

If you have any questions or issues regarding the Wabeltools CLI, please open an issue on our GitHub repository.
