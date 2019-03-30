# Route 53 Updater

Simple Route 53 updater that updates dns mappings with up to date public IP.

## Build

Built for raspberry pi ARM.

## Flow

* The cron job on Raspberry Pi will execute a binary which will determine the public IP of the home network and forward that IP to the AWS Lambda.
* The cron job and the lambda will share a secret key. which will be hashed along with a nonce