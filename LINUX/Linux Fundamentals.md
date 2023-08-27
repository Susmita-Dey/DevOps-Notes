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

Following are some basic Linux commands:

1. ls: ls allows us to view the directory and it's contents
2. pwd: pwd is used to get our current directory where we are currently located.
3. cd : cd is used to change directories i.e., from one directory to another directory.
4. clear: clear is to clear the terminal screen and get to the top of the screen.
5. lsb_release: lsb_release is used to print distribution-specific information.
For example:
    ```bash
    ubuntu $ lsb_release -a
    No LSB modules are available.
    Distributor ID: Ubuntu
    Description:    Ubuntu 20.04.5 LTS
    Release:        20.04
    Codename:       focal
    ```
6. man: man is an interface to the system reference manuals.
7. cat: cat is used to print the contents present in a file. It is also used to concatenate 2 files.
For example:
    ```bash
    ubuntu $ cat /etc/os-release
    NAME="Ubuntu"
    VERSION="20.04.5 LTS (Focal Fossa)"
    ID=ubuntu
    ID_LIKE=debian
    PRETTY_NAME="Ubuntu 20.04.5 LTS"
    VERSION_ID="20.04"
    HOME_URL="https://www.ubuntu.com/"
    SUPPORT_URL="https://help.ubuntu.com/"
    BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
    PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
    VERSION_CODENAME=focal
    UBUNTU_CODENAME=focal
    ```
8. uname: uname is used to print system information.
    ```bash
    ubuntu $ uname -a
    Linux ubuntu 5.4.0-131-generic #147-Ubuntu SMP Fri Oct 14 17:07:22 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux
    ```
9. lscpu: lscpu is used to display information about CPU architechture.
10. lsmem: lsmem is used to display the memory information. It lists the ranges of available memory with their online status. 
    ```bash
    ubuntu $ lsmem
    RANGE                                 SIZE  STATE REMOVABLE BLOCK
    0x0000000000000000-0x000000007fffffff   2G online       yes  0-15

    Memory block size:       128M
    Total online memory:       2G
    Total offline memory:      0B
    ```
11. find: find is used to search for files in a directory hierarchy.
    ```bash
    ubuntu $ find / -name 'syslog'
    /run/systemd/journal/syslog
    /var/log/syslog
    /usr/local/go/src/log/syslog
    ```
12. less
13. more
14. grep 

### Creating a User
useradd - create a new user or update default new user information

```bash
ubuntu $ useradd -D
GROUP=100
HOME=/home
INACTIVE=-1
EXPIRE=
SHELL=/bin/sh
SKEL=/etc/skel
CREATE_MAIL_SPOOL=no
```
This is the default configuration of stuffs like shell type, home, group, etc., for all users.

By using `useradd -D -s /bin/bash`, you can change the default bash type to '/bin/bash' for all users.

For all other commands, please refer to [My Handwritten Notes of Intro to Linux & Terminal Commands](/LINUX/Handwritten%20Notes%20of%20Intro%20to%20Linux%20&%20Terminal%20Commands.pdf)

## Packages 

## Create Files

## Permissions 

## Redirects

## Services

## Git 
