# yt-server-oauth
This script returns youtube credentials in base64 format. This was inspired by [google's auth code sample](https://developers.google.com/youtube/v3/code_samples/go).

This script removed all of the uneeded code for the client workflow. As well as no longer generating/reading from a local secrets file.

## Setup
1. To get access to the youtube api it requires an account under [Google Cloud Platform](https://cloud.google.com/)
2. Create a project
3. Under APIs & Services Enable the YouTube Data API v3 
4. Under YouTube Data API v3 go to credentials and click Create credentials
5. Choose OAuth clientID and for Application type pick Other
6. Save the presented client id as the environment variable YOUTUBE_CLIENT_ID 
7. Save the presented client secret as the environment variable YOUTUBE_SECRET

## Execution
You must have the following installed/configured for this script to run<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)

To run the script enter the following command.
```bash
docker-compose run ytauth
```