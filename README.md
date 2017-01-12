# tlist

Tlist is an in-memory time series list package for Golang.

[![](https://img.shields.io/circleci/token/2dc3aeee87f95b35fb9229f88dce56f01e6b4159/project/abcum/tlist/master.svg?style=flat-square)](https://circleci.com/gh/abcum/tlist) [![](https://img.shields.io/badge/status-alpha-ff00bb.svg?style=flat-square)](https://github.com/abcum/tlist) [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abcum/tlist) [![](https://goreportcard.com/badge/github.com/abcum/tlist?style=flat-square)](https://goreportcard.com/report/github.com/abcum/tlist) [![](https://img.shields.io/badge/license-Apache_License_2.0-00bfff.svg?style=flat-square)](https://github.com/abcum/tlist) 

#### Features

- In-memory doubly linked list
- Store values by version number
- Delete values by version number
- Find the initial and the latest version
- Ability to insert items at any position in the list
- Find exact versions or seek to the closest version
- Select items by version number or retrieve latest value
- Less efficient than a btree when seeking for a specific version: O(n) worst case
- Efficient when majority of selects are for the initial or latest version: O(1) worst case

#### Installation

```bash
go get github.com/abcum/tlist
```
