#A simple revelapp in docker

#To build an image
docker build -t revel-app .

#To run image
docker run -it -v /home/prabesh/package/src/github.com/PrabeshNewBee/revel-app:/src/github.com/PrabeshNewBee/revel-app -p 9000:9000 revel-app R