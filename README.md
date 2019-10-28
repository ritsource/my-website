Visit the App - [https://ritwiksaha.com](https://ritwiksaha.com)

# About

A **Multi-container Docker** application, running two **Go** servers for API & template rendering, a **Nginx** server for handling the routing between services, and a **React** app (management-console) for managing content. And uses **MongoDB** as the database.

<img style="float: right;" src="https://gitlab.com/ritwik310/project-documents/raw/master/My-Website/My-Website-Microservices-Mockup-0.png"/>

# Description

The application contains a **Renderer** container, running a **Go server** that is responsible for rendering the correct HTML from server-side. Renderer requests content `raw markdown` from the **API** and then renders that markdown as HTML.

Here's a very simple diagram of a normal request flow.

<img style="float: right;" src="https://gitlab.com/ritwik310/project-documents/raw/master/My-Website/My-Website-Request-Flow-Mockup-0.png"/>

For the **API-Server**, other than all the **Routing**, **Authentication** (admin only) and **Static file serving**; the API serves documents, which is kind of static file serving, but from **remote hosting** and local **caching**.

### Caching Documents

The necessary files for **Content** (Blog and Project documents) are saved remotely on some generous free service file-hosting provider (**Github** for example) and linked in the database document. For the **first request** (doc request), the cilent **gets redirected** to the actual source (Github/Gitlab) and the file gets **cached** in the filesystem (this uses **volumes** in **Docker-environment**). The next requests doesn't have to go teh the actual source anymore, once its saved in the cache.

<img style="float: right;" src="https://gitlab.com/ritwik310/project-documents/raw/master/My-Website/My-Website-Doc-Caching-Mockup-0.png"/>

Other than that, as mentioned in the top, there's also a **React-app** for content management. And **Nginx-server** that handles routing between outside requests and different application services.

# Running Locally

### Clone the Project

First clone the project using
```shell
git clone git@github.com:ritcrap/my-website.git
```

Or by downloading zip from [here](https://github.com/ritcrap/my-website)

### Running using Docker

I you already have **Docker** and **Docker-compose** installed in your system just run this from root project directory

```shell
export DB_VOLUME_PATH=$HOME/Desktop/data/mongo-data-27017 # Some location to save database data
bash ./start.sh
```

### or Running without Docker

First you need to start a MongoDB server on `localhost:27017` or You can also use remote cluster `example.com/5050`

Run the following commands...

```shell
cd ./api
bash ./development.sh
```

```shell
cd ./renderer
bash ./development.sh
```

```shell
cd ./console
bash ./development.sh
```

If you want to use Admin-console you need these following **environment variables** in the ./api environment
```shell
export SESSION_KEY="$SOME_RANDOM_SECRET"

export GOOGLE_CLIENT_ID="" # GOOGLE_CLIENT_ID from https://console.cloud.google.com
export GOOGLE_CLIENT_SECRET="" # GOOGLE_CLIENT_SECRET from https://console.cloud.google.com

export ADMIN_EMAIL_A="" # Authorized admin email address No.1 (Connected to Google account)
export ADMIN_EMAIL_B=""# Authorized admin email address No.2 (Connected to Google account)
```


Eventually I don't think anyone would like to run a personal site in their local machine. 🤷   
Anyways, **Happy Hacking**

<!--View Source-code - [https://github.com/ritcrap/my-website](https://github.com/ritcrap/my-website)-->