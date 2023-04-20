# Data Access with Database/Sql package

## Prerequisites
- **[Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)**
- **Some programming experience.** 
    The code here is pretty simple, but it helps to know something about functions.

- **A tool to edit your code.** 
    ny text editor you have will work fine. Most text editors have good support for Go. The most popular are VSCode (free), GoLand (paid), and Vim (free).

- **A command terminal.** 
    Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.
    
- **Mysql Docker**
    Deploy a container for mysql docker with tag 8 (mysql version 8)

    quick start with command
    ```
    docker pull mysql:8
    docker run -p 3306:3306 -p 33060:33060 --name=mysql -e MYSQL_ROOT_PASSWORD=my-root-passwd -d mysql:8
    ```

    For more information, see the [docs](https://hub.docker.com/_/mysql)

## Tutorial

Go tutorials topic [Accessing a relational database](https://go.dev/doc/tutorial/database-access) : 
Introduces the basics of accessing a database using the standard library.