# sadExtractor
sadExtractor is a simple recon tool that extract all links from a web page. 


## Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go install -v github.com/SaDs3c/sadExtractor@latest
```

Otherwise [download a release for your platform](https://github.com/SaDs3c/sadExtractor/releases).
To make it easier to execute you can put the binary in your `$PATH`.

## Usage

Single URL: 
```
sadExtractor -d <domain> -o output.txt
```

Multiple URL in File:
```
sadExtractor -l urls.txt -o output.txt
```
