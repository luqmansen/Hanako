# Project Hanako [![Build Status](https://travis-ci.com/luqmansen/Hanako.svg?branch=master)](https://travis-ci.com/luqmansen/Hanako)

![](docs/assets/hanako-small.png)


## What is This ?
This is an anime database services based on [go-micro](https://github.com/micro/go-micro) framework  

## What's working now
- Anime service on progress, some of endpoint working now

Other stuff still on development, fork this repo and help me :D

## Prerequisites
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

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
