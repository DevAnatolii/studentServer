# Students

A sample RESTFull server and Android client for managing students results.</br>
Android application works on devices with JELLY_BEAN os and earlier.</br>

<a href="https://imgflip.com/gif/1wqi6l"><img src="https://i.imgflip.com/1wqi6l.gif" title="made at imgflip.com"/></a>

# How to launch
* First of all you need to install and launch MySql server with user root (password `0000`) on 3306 port locally
Create a database `university` and add `students` table there:
```
 CREATE TABLE students (
 id  VARCHAR(8) NOT NULL,
 name VARCHAR(64) NOT NULL,
 score INT NOT NULL,
 PRIMARY KEY (id)
 );
```
* Second step is to launch RESTFull server locally and obtain IP according to https://stackoverflow.com/a/4779992 </br>
* Next step is to install sample Android application provided in repository and connect your device to the same WiFi.</br>
* Finally you need to launch the Android application, input the server endpoint and tap on connect button.
Server endpoint should be following format: `http://{your_ip}:8080`
If connection is established successfully, both button and input field will be disabled otherwise an error message will
be shown and you will be able to correct the server endpoint.

# Different database implementation
If you would like to start the RESTFull server and do not affect database, you should launch the one with parameter `-storage=temporary`.
In this case the implementation of MySql database will be replaced to map implementation, that exists while server is running