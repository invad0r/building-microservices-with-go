# building microservices with go

following nic allong with his series [youtube](https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)


## setup swagger
```sh
sudo apt install gnupg ca-certificates
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
echo "deb https://dl.bintray.com/go-swagger/goswagger-debian ubuntu main" | sudo tee /etc/apt/sources.list.d/goswagger.list
apt update 
apt install swagger
```
https://goswagger.io/install.html#debian-packages