# building-limit-height-plateau-api
This service contains one REST api endpoint. When valid json is posted to this api endpoint, then the service will find the coordinates for all of the features of `building_limits` and `plateau_heights`. These coordinates will save records of building_limits and plateau_heights to local mongodb collections.

## steps to run `proto/Makefile` as intended
1. run `make install`
2. run `make`

## steps to run locally and hit locally running API
1. `docker-compose up`
2. `curl -X POST -H "Content-Type: application/json" -d '{"your_JSON_body_here": "potato"}' http://localhost:9002/v1/buildingdata`

## What I would do differently next time.
1. Avoid gRPC. I like gRPC. It handles a bunch of validation automatically. However, getting gRPC to generate code that would accept two-or-more-dimensional took a lot of time to figure out. In addition, the solution for the 2+d json arrays isn't clean or very stable. I did learn a bunch of protoc tricks in the meantime though, so that was good.
2. Investigate what geographic polygons and intersections more thoroughly sooner. I thought that I had a good enough understanding to design a my own solution. It wasn't until later that I understood what the problem of polygons and intersections was actually was. Had I understood this sooner, I would have ditched gRPC sooner so that I could have easily injested geoJSON as a standardized format. There are probably REST validators for geoJSON already on github. I would have also learned more early that none of the available Golang tools for polygons were missing a key feature. I investigated https://github.com/paulmach/orb, https://github.com/golang/geo, and https://github.com/twpayne/go-geom only to learn that none of them support automatic calculations of the intersection/common-polygon from two given polygons. Had I known that sooner, then I would have investigated if there were some packages for this in a different langauge (such as JS or PHP).