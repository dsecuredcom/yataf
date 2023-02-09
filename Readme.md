# [y]et [a]nother [t]ool for [a]nalysis of [f]iles

Simple tool to analyze a files/urls content - it was primarily created to analyze the content of a javascript file
against a given set of regular expressions.
The main goal is to give you an idea if a files content might be of interest. This means that __yataf__ tries to find
secrets in the content as well as potential endpoints.

## Make

```
go build -o dist/yataf main.go
```

## Usage

Analyze a file using all regular expressions

```
./yataf -file some/file.js
```

Analyze an url using all regular expressions but limit the extracted content to 1000 characters (per regular expression)

```
./yataf -file https://an.url/file.js -char-limit 1000 -types all
```

Analyze a file using only regular expressions of type "credentials"

```
./yataf -file some/file.js -types credentials
```

## Types

* "all" (default)
* "credentials"
* "urls-paths"

## Good to know

* There is a hard coded list of negative entries.
* Tested @ Ubuntu 22.10 (should work on windows as well)
* Developed with Go 1.19

## WIP

As always this is a WIP, there are some todos, optimizations and so on!