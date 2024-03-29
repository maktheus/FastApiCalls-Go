# Fast API Caller

This project is a Go application that makes API calls in threads to improve performance.

## Prerequisites

- Go 1.15 or later

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository:

```
git clone https://github.com/maktheus/FastApiCalls-Go
```

Copy code

2. Navigate to the project directory:

```
cd FastApiCaller
```

Copy code

3. Build the project:

```
go build
```

Copy code

This will create an executable file for the project.

4. Run the executable file:
```
./[executable_filename]
```

You can also specify command-line flags when running the executable file. For example, to set the duration flag to 120 seconds, you can run the following command:

```
./[executable_filename] -d 120
```

add url flag


```
./[executable_filename]  -u https://www.google.com
```

change file type

```
./[executable_filename]  -f json
```

change file name

```
./[executable_filename]  -n myFile
```

## Type of files

```
./[executable_filename]  -f json
```

```
./[executable_filename]  -f csv
```

```
./[executable_filename]  -f xml
```

```
./[executable_filename]  -f txt
```




## Runing the executable file

Inside bin folder you will find the executable file. You can run it by typing the following command:

```
./events
```

## Running the tests

To run the tests for this project, run the following command:

```
cd tests
```

```
go test
```

## Built With

- [Go](https://golang.org/) - The programming language used

## Authors

- **Maktheus** - _Initial work_ - [Maktheus]( 

## Contributing

If you want to contribute to this project, please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
I hope this helps! Let me know if you have any questions.