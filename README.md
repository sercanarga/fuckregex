# fuckregex

A user-friendly AI-powered tool for those who hate writing regular expressions! Powered by GPT-3.5, FuckRegex helps you generate regex patterns without the headache.

## Features

* Generate regex patterns with natural language input
* Interpret entered regex for better understanding
* Validate generated regex patterns
* Generate regex usage examples in popular programming languages

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

Ensure you have [golang](https://golang.org/dl/) (v1.20 or later) and [postgresql](https://www.postgresql.org/) database installed on your system:

### Installation

1. Clone the repository:

```bash
git clone https://github.com/sercanarga/fuckregex.git
````

2. Go to the project directory:

```bash
cd fuckregex
```

3. Modify `.env` file:
```bash
nano .env
```

4. Create a postgresql database using the sql file under the `resources` directory.
```bash
cat resources/*.sql
```

5. Install dependencies:

```bash
go mod tidy
```

6. Run the API service:

```bash
go run cmd/api/main.go
```

7. Run the web service:

```bash
go run main.go
```

8. Open your web browser and visit `http://localhost:8181` (replace `8181` with the port number specified in the configuration).

## Contributing

We welcome contributions from the community! You can open PRs and issues to help us improve the project.

## To-Do

- [ ] Improve security measures
- [ ] Implement request moderation for ChatGPT submissions
- [ ] Add the ability to interpret entered regex
- [ ] Introduce regex validation for generated patterns
- [ ] Generate regex usage examples in popular programming languages
- [ ] Enhance design and source code

## License

This project is licensed under the [MIT License](https://github.com/sercanarga/fuckregex/blob/main/LICENSE).
