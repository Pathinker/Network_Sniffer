# Network Sniffer 

<p align="justify">    
Intercepts and analyzes data packets flowing through the Internet, it reads the data that goes to the network interface card and operating system.
</p>

## How to install

1.- Clone the repository.

2.- Execute the start.sh file, for this you will need first to give executable permissions to the file with``` chmod +x start.sh ``` and then type ```./start.sh```, this will load the configuration file.

3.- Change directory to the folder network ``` cd network ``` and perform the main Go file with ``` go run main.go ```.

> [!CAUTION]
> The Docker and Devcontainer deployment gives more compatibility however, Docker is limited and does not gather all the traffic that running it locally will, only showing TCP.
