# Go Web Server Project

This project implements a simple web server using Go.

## Getting Started

To setup a  webserver on your system, follow these simple steps:
1. **Clone the repository:**

   ```bash
   https://github.com/veera-damisetti/go-web-server.git
   cd go-web-server
2. **Build the Project**
   ```bash
   go build main.go
3. **Create a Service Unit File**
    ```bash
    sudo vi /etc/systemd/system/v-server.service

    [Unit]
    Description=Go Server by Veera
    After=network.target

    [Service]
    ExecStart=/bin/bash -c '<path to binary>'
    WorkingDirectory= <path to binary workdir>
    Restart=always
    User=root
    Group=root

    [Install]
    WantedBy=multi-user.target

4. **Reload systemd**
   ```bash
   sudo systemctl daemon-reload
5. **Enable and Start the Service**
    ```bash
    sudo systemctl enable v-server
    sudo systemctl start v-server


--
**Note:**
* By defualt webserver will listen on port 8443, you can change the port in main.go file 
