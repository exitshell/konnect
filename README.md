# Konnect

*Simple remote host management*

<img src="https://img.shields.io/github/release/exitshell/konnect.svg" /> <img src="https://img.shields.io/github/license/exitshell/konnect.svg" />

Konnect is a tool for managing and connecting to remote hosts. By defining a list of hosts in a configuration file, Konnect allows you to connect and run commands on the defined hosts.

Installation
--------------
**Manual Download**

The easiest way to install Konnect is to head over to the [releases](https://github.com/exitshell/konnect/releases) tab and download the binary for your desired OS. As of now, MacOS and Linux are supported.

**Brew**

You can also download and install Konnect via brew.
1. `brew tap exitshell/konnect git@github.com:exitshell/homebrew-konnect.git`
1. `brew install konnect`

Usage
--------------

####  Defining hosts in the config file

The configuration file for konnect should be named `konnect.yml`.
In this `konnect.yml` config, we define two hosts: **app**, and **database**.

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

####  Connecting to hosts
In order to connect to a host you would run: `konnect to <host>`

- _example_: **`konnect to app`**

Alternatively, you can just run `konnect`, and it would start an interactive prompt for you to choose a host to connect to.

#### Defining tasks
Tasks are essentially bash commands which execute remotely on a host. We can define tasks in the `konnect.yml` file, and then run them on a specific host.

In this `konnect.yml` config, we define one host (**app**) and one task (**tailsys**).

```
hosts:
  app:
    user: root
    host: 192.168.99.100
    port: 22
    key: /home/app/key

tasks:
  tailsys: tail -f -n 100 /var/log/syslog
```

####  Running tasks
In order to connect to run a task on a host you would run: `konnect to <host> and <task>`

- _example_: **`konnect to app and tailsys`**


Commands
--------------

- **View all defined hosts**

	`konnect list`


- **Connect to a defined host**

	`konnect to <host>`


- **Connect to a host and run a task**

	`konnect to <host> and <task>`


- **Display the SSH command for a host**

	`konnect args <host>`


- **Create an empty konnect.yml config file**

	`konnect init <dir>`

- **Edit the konnect.yml config file**

	`konnect edit`


- **Check the status of one or more hosts**

	`konnect status <host1> <host2>`

	`konnect status --all`


Authors
--------------
Konnect was created by [Sandeep Jadoonanan](https://github.com/TunedMystic) 🤓

License
--------------
Konnect is released under the MIT license.

See [LICENSE](LICENSE) for the full text.
