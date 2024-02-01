# IP and Server Name Lookup Go Project

This is a Go (Golang) project that provides a command-line tool for retrieving information about IPs and Server Names associated with a specific host.

## 🚀 How to Use

### Prerequisites

Make sure you have Go installed on your machine. To install Go, visit [golang.org](https://golang.org/doc/install).

### 🔄 Cloning the Repository

```bash
git clone https://github.com/Renanllm/search-ips-and-sn.git
cd search-ips
```

### ⬇️ Downloading Dependencies

Use the following command to download the project dependencies:

```bash
go mod download
```

### 🛠️ Building the Project

To build the project, use the following command:

```bash
go build
```

### 🔍 Running IP Search

To search for IPs associated with a host, execute the following command:

```bash
./main ip --host {host}
```

Example:

```bash
./main ip --host amazon.com
```

### 🔍 Running Server Name Search

To search for Server Names associated with a host, execute the following command:

```bash
./main name --host {host}
```

Example:

```bash
./main name --host amazon.com.br
```

## 🤝 Contributions

Contributions are welcome! Feel free to open issues, propose improvements, or submit pull requests.

## 📝 License

This project is licensed under the [MIT License](LICENSE).
