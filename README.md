# Konnect - connect to thing!

Konnect is a tool for managing and connecting to remote hosts. By defining a list of hosts in a configuration file, Konnect allows you to connect and run commands on the defined hosts.

Installation
--------------
**Manual Download**

The easiest way to install Konnect is to head over to the [releases](https://github.com/exitshell/konnect/releases) tab and download the binary for your desired OS. As of now, MacOS and Linux are supported.

**Brew**

You can also download and install Konnect via brew.
1. `brew tap exitshell/konnect git@github.com:exitshell/konnect.git`
1. `brew install konnect`

Usage
--------------


Using Konnect is _really_ simple:
1. Define hosts in a `konnect.yml` config file.
1. Connect to a defined host.


A `konnect.yml` config file looks like this:

```
hosts:
  app:
    user: root
    host: 192.168.99.100
    port: 22
    key: /home/app/key
  database:
    user: admin
    host: 127.0.0.1
    port: 89
    key: ~/.ssh/id_rsa
```



### Isn't this like an ssh config file?
Yup, Konnect  is very much like an ssh config file.

However, Konnect aims to be more portable and configurable by providing additional functionality on a per-host / multi-host basis.

Examples
--------------

- **Create an empty konnect.yml config file**

	`konnect init <dir>`


- **View all defined hosts**

	`konnect ls`


- **Connect to a defined host**

	`konnect to <host>`


- **Display the SSH command for a host**

	`konnect args <host>`


- **Check the status of one or more hosts**

	`konnect status <host1> <host2>`

	`konnect status --all`


### Why use konnect?
...why not? 🤓


License
--------------
Konnect is released under the MIT license.

See [LICENSE](LICENSE) for the full text.
