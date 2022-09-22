# how to deploy to heroku
### create a heroku account
### install heroku cli
### create a heroku app
![Screenshot](1.png)
### create Mysql database
![Screenshot](2.png)
### set environment variables
![Screenshot](3.png)
###  add a remote to your git repo
```shell
heroku git:remote -a agmc-alterra-day8
```
### set stack to container
```shell
heroku stack:set container
```
###  push to heroku
```shell
git push heroku master
```
![Screenshot](4.png)
###  open the app in browser
![img.png](5.png)
### Test Postman
![img.png](6.png)
###  check the logs
```shell
heroku logs --tail
```
![img.png](7.png)