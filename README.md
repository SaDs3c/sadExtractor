[![GitHub release](https://img.shields.io/badge/release-v1.0.0-brightgreen?style=flat-square)](https://github.com/SaDs3c/sadExtractor/releases/tag/V1.0.0)
[![GitHub stars](https://img.shields.io/github/stars/SaDs3c/sadExtractor?style=flat-square)](https://github.com/SaDs3c/sadExtractor/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/SaDs3c/sadExtractor?style=flat-square)](https://github.com/SaDs3c/sadExtractor/network)
[![GitHub issues](https://img.shields.io/github/issues/SaDs3c/sadExtractor?style=flat-square)](https://github.com/SaDs3c/sadExtractor/issues)
[![GitHub license](https://img.shields.io/github/license/SaDs3c/sadExtractor?style=flat-square)](https://github.com/SaDs3c/sadExtractor/blob/main/LICENSE)

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
