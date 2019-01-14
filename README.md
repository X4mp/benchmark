# Benchmark
This is a benchmark blockchain, where server providers and benchmark requesters create benchmark reports and sells them to customers that needs benchmark data.

## Install dependencies
We use the [dep dependency manager](https://github.com/golang/dep) in order to manage dependencies.  Make sure you have it installed on your computer, if you don't visit the link above and follow instructions to install it.  

Then, install the dependencies using this command:

~~~~
dep ensure
~~~~

## Command line application
To run the command line tool, simply execute this command:

~~~~
go run main.go --help
~~~~

This will show you the possible commands and an explanation for each.

### Generate config file
The config file contains an encrypted private key, protected by a password you give to the tool.  To generate a new config file, simply type this command:

~~~~
go run main.go generate --file ./credentials.xmn --pass MYPASS --rpass MYPASS
~~~~

Make sure to replace the password properly.  You can also replace the filename to whatever file you want.

### Spawn the blockchain
To spawn the blockchain, you need to tell the application where is your config file, and what is the password to decrypt your config file.  Here's the command:

~~~~
go run main.go spawn --dir ./db --pass MYPASS --file ./credentials.xmn
~~~~

The --dir parameter represents the folder where the blockchain will save its data.  The --pass represents the password to decrypt the config file and the --file parameter is the path to your encrypted file.
