# RocketFeed
A simple program to convert Gemini feeds to atom feeds.

## Installation
Either download one of the executables from the releases page or if you have Go installed you can clone the repository and run `go build` or `go install`.

## Usage
```
Usage: rocketfeed -b base-url gemfeed
  -b string
    	Base URL. This is required and should be where your gemfeed is located. Ex: gemini://example.org/gemlog/
  -n int
    	Number of most recent items to include in the atom feed. All items from the gemfeed are included by default.
  -o string
    	Where to save the converted atom feed. If not provided, prints to stdout.
  -t string
    	Feed title. Defaults to the first top level heading found in the gemfeed.
```
