# File Search and Split Utility

A command-line utility designed for file management to perform various operations on text files, including searching for specific words, counting words and lines, splitting large files into manageable chunks, replacing words, and retrieving the most frequently occurring words.

## Installation

To use this utility, clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd <project-directory>
```

Then, build the Go application:

```bash
 go build -o fily main.go
```
## Usage

The utility behaves differently based on the specified type:

- **Search**: Searches for a specific word in the file and returns the lines containing that word. If the `-ln` flag is set, it will also include line numbers in the output.
  
- **Count**: Counts the total number of words and lines in the specified file and outputs the results.

- **Split**: Splits the specified file into smaller chunks based on the provided chunk size in MB. The split files are saved to the specified destination path.

- **Replace**: Replaces all occurrences of a specific word in the file with a new word. 

- **Top Words**: Retrieves the most frequently occurring words in the specified file. 


### Flags

- `-t`: Type of operation (`search`, `count`, `split`, or `replace`).
- `-fp`: File path to the file you want to process.
- `-w`: Word to search for (only used with `search` and `replace` types).
- `-r`: New word to replace the old word with (only used with `replace` type).
- `-dp`: Destination path for the split files (only used with `split` type).
- `-cs`: Chunk size in MB (only used with `split` type).
- `-ln`: Line number flag (set to `true` to include line numbers in search results).
- `-ow`: **Old word** to replace (only used with `replace` type).
- `-nw`: **New word** to replace the old word with (only used with `replace` type).
- `-k`: Number of top words to retrieve (only used with `topwords` type).



### Examples

#### Search Example

To search for the word "example" in a file located at `/path/to/file.txt`, use:

```bash
./fily -t search -fp /path/to/file.txt -w example -ln
```

#### Split Example

To split a file located at `/path/to/largefile.txt` into chunks of 5 MB and save them to `/path/to/output/`, use:

```bash
./fily -t split -fp /path/to/largefile.txt -dp /path/to/output/ -cs 5
```

#### Count Example

To count numbers of words and line in a file located at `/path/to/largefile.txt`, use:

```bash
./fily -t count -fp /path/to/largefile.txt 
```

#### Replace Example

To replace the word "oldword" with "newword" in a file located at `/path/to/file.txt`, use:

```bash
./fily -t replace -fp /path/to/file.txt -w oldword -r newword
```
#### Top Words Example

To retrieve the top 10 words from a file located at `/path/to/file.txt`, use:

```bash
./fily -t topwords -fp /path/to/file.txt -k 10
```