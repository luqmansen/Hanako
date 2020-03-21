# Project Hanako
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/luqmansen/hanako/blob/master/docs/LICENSE.md)
[![Build Status](https://travis-ci.com/luqmansen/Hanako.svg?branch=master)](https://travis-ci.com/luqmansen/Hanako)
[![Lint Status](https://github.com/luqmansen/Hanako/workflows/Go/badge.svg)](https://github.com/luqmansen/Hanako/actions)
[![Commit activity](https://img.shields.io/github/commit-activity/w/luqmansen/hanako?foo=bar)](https://github.com/luqmansen/Hanako/pulse)
![last commit](https://img.shields.io/github/last-commit/luqmansen/hanako)

<p align="center">
    <img src="https://github.com/luqmansen/Hanako/raw/master/docs/assets/hanako-super-small.png"/>
</p>

## What is This ?
This is an anime database services based on [go-micro](https://github.com/micro/go-micro) framework  

## What's working now
- Anime service and Rest API is working now 

Other stuff still on development, fork this repo and help me 	（‐＾▽＾‐）

## Prerequisites
- [micro v2](https://github.com/micro/micro) 
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

Please follow specific instructions in each services for more info

## Development
##### Run testing
 via micro command line
````cgo
micro call hanako.<type(default=srv)>.<service_name> <ServiceName>.<Method_name> '{"json_param" : "value"}'

````
Or via micro web (Recommended) 
```cgo
micro web
```
Then call service via client page

##### Others
Please read readme on each service for specific instruction

## Deployment

- [Heroku](https://github.com/luqmansen/Hanako/wiki/Deployment) (Deprecated)
 
## Roadmap
Read the roadmap on the wiki [here](https://github.com/luqmansen/Hanako/wiki/Roadmap)

## Contributing
For now, just fork it and create a pull request, there will be a contribution guidelines soon!

## Authors

<a href="https://github.com/luqmansen/Hanako/graphs/contributors">
  <img src="https://contributors-img.firebaseapp.com/image?repo=luqmansen/Hanako" />
</a>


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/luqmansen/hanako/blob/master/docs/LICENSE.md) file for details
