# Linux Fundamentals
This is a collection of notes on Linux fundamentals. The notes are based on the [Linux Fundamentals](https://www.youtube.com/watch?v=EUu1E_YKGTw) course on Kubesimplify channel.

Table of Contents
=================

- [Linux Fundamentals](#linux-fundamentals)
- [Table of Contents](#table-of-contents)
  - [Linux Operating System](#linux-operating-system)
    - [Linux Distributions](#linux-distributions)
    - [Components of Linux](#components-of-linux)
  - [Linux File System](#linux-file-system)
  - [Linux Commands](#linux-commands)
  - [Packages](#packages)
  - [Create Files](#create-files)
  - [Permissions](#permissions)
  - [Redirects](#redirects)
  - [Services](#services)
  - [Git](#git)

## Linux Operating System

The Linux Operating System is essential for running and interacting with the hardware of a computer. Without an operating system we would have to interact with the hardware directly which is quite impossible. The operating system provides a layer of abstraction between the hardware and the user.

### Linux Distributions

Distributions are just different variation of the operating system.
There are various distributions of Linux. The most popular ones are Ubuntu, Debian, Red Hat, CentOS, Fedora, etc. These are just different flavours of Linux. Some of them have different package managers, different file systems, etc.

### Components of Linux

Linux Operating System is made up of three main components: the kernel, the shell and the file system.

## Linux File System

The Linux File System is a tree-like structure. The root directory is the topmost directory in the file system. It is represented by a forward slash `/`. The root directory contains other directories and files. The directories and files in the root directory are called children of the root directory.

> **Note:** Linux is just a series of files and sometimes those files interact in a certain way to do some special things. You may call it as a file-based operating system. It is just basically grouping of files. 
For example, the `/dev` directory contains files that represent devices. The `/dev/null` file is a special file that is used to discard output. The `/dev/zero` file is a special file that is used to generate zeros. The `/dev/random` file is a special file that is used to generate random numbers.

**The root directory contains the following directories:**
- `/bin` - contains binary files
- `/boot` - contains boot files used to boot the system.
- `/dev` - contains device configuration files like keyboard and mouse that allow your computer to interact with the peripherals or the hardware.
- `/etc` - contains system wide configuration files like network configuration files, user configuration files, etc.
- `/home` - contains home directories where each user can store their personal files or documents.
- `/lib` - contains library files
- `/media` - contains removable media like USB drives, CD-ROMs, etc., that allows you to interact with the removable media and checkout the files inside it.
- `/mnt` - contains mounted file systems  like network file systems, etc. It is a temporary mount point.
- `/opt` - contains optional software i.e., the software or programs thar we install manually. For example, if you have a piece of software that's built for Linux, that software will be distributed throughout multiple directories in your file system. The `/opt` directory is a place where you can put all of those files in one place. It is for programs that install bin and lib in one directory.
- `/proc` - contains process information. It represents the current state of kernel. 
- `/root` - contains root home directory. Root is the super user of the system. It is the administrator of the system.
- `/run` - contains temporary runtime files
- `/sbin` - contains system binaries
- `/snap` - contains snap packages. It is a package manager for Linux that allows you to install software from the internet.
- `/srv` - contains service data i.e., data for services provided by the system.  It contains site data like geolocation data, DNS data, etc.
- `/sys` - contains system information. It is a virtual file system that is used to interact with the kernel.
- `/tmp` - contains temporary files. It is a temporary location for running processes.
- `/usr` - contains user programs and data. It is a place where you can put all of your user binaries and files.
- `/var` - contains variable files like log files, cache files, etc.
- `/lost+found` - contains recovered files


The root directory also contains the following files:
- `/etc/fstab` - contains file system information
- `/etc/passwd` - contains user information
- `/etc/shadow` - contains encrypted passwords
- `/etc/hosts` - contains host names
- `/etc/resolv.conf` - contains DNS information
- `/etc/group` - contains group information
- `/etc/profile` - contains user profile information
- `/etc/issue` - contains system information
- `/etc/issue.net` - contains system information
- `/etc/mtab` - contains mounted file system information

## Linux Commands 

## Packages 

## Create Files

## Permissions 

## Redirects

## Services

## Git 
