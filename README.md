<div id="top">

<!-- HEADER STYLE: CLASSIC -->
<div align="center">

<img src="./dockenv-logo.png" width="30%" style="position: relative; top: 0; right: 0;" alt="Dockenv Logo"/>

# DOCKENV

<em></em>

<!-- BADGES -->
<img src="https://img.shields.io/github/license/arasdenizhan/dockenv?style=default&logo=mit&logoColor=white&color=0080ff" alt="license">
<img src="https://img.shields.io/github/last-commit/arasdenizhan/dockenv?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/arasdenizhan/dockenv?style=default&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/arasdenizhan/dockenv?style=default&color=0080ff" alt="repo-language-count">

<!-- default option, no dependency badges. -->

<!-- default option, no dependency badges. -->

</div>
<br>

---

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
  - [Project Index](#project-index)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Testing](#testing)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

**dockenv** is a lightweight security tool that analyzes Docker containers and images for exposed or risky environment variables.

Modern applications frequently store secrets such as API keys, database credentials, tokens, and private configuration values inside environment variables. While convenient, this practice can easily lead to accidental leaksö especially in shared images, misconfigured deployments, or CI/CD pipelines.

**dockenv** helps you catch these problems early.

It inspects Docker environments, extracts environment variables, and applies an intelligent rule-based analysis engine to detect sensitive data and assign a risk severity score. This allows developers and DevOps engineers to quickly identify security risks before they reach production.

The tool is designed to be fast, dependency-free, and CI-friendly, making it suitable for local development, pre-deployment validation, and automated security checks.

---

## Features
- 🔍 **Docker Environment Variable Security Scanning**
	- Scans Docker containers and images to detect sensitive or risky environment variables (e.g. secrets, tokens, passwords, API keys) 

- 🧠 **Rule-Based Risk Analysis Engine**
	- Built-in detection rules identify common secret patterns:
		- Credentials (PASSWORD, PASS, SECRET)
		- Tokens (TOKEN, JWT, AUTH)
		- API keys (API_KEY, ACCESS_KEY)
		- Private keys & sensitive configs

- 📊 **Risk Scoring System**
	- Each finding is evaluated and assigned a severity score so you can quickly prioritize real threats over noise.

- 🐳 **Supports Both Containers and Images**
	- Scan running containers
	- Scan local Docker images before deployment
	- Shift-left security checks in CI pipelines

- 🖥️ **Readable CLI Output**
	- Colored terminal output
	- Structured table format
	- Clear severity visualization (Low / Medium / High)

- ⚡ **Fast & Lightweight**
	- Written in Go, runs instantly without agents, sidecars, or external services.

- 🔌 **CI/CD Friendly**
	- Easily usable inside:
		- GitHub Actions
		- GitLab CI
		- Jenkins pipelines
		- Pre-deployment checks

- 🧩 **Modular Architecture**
	- Clean internal modules:
		- docker → container/image inspection
		- scanner → env extraction
		- analyzer → rules & scoring
		- output → formatting & reporting

---

## Project Structure

```sh
└── dockenv/
    ├── cmd
    │   ├── root.go
    │   └── scan.go
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── analyzer
    │   ├── docker
    │   ├── output
    │   └── scanner
    ├── main.go
    └── pkg
        └── model
```

### Project Index

<details open>
	<summary><b><code>DOCKENV/</code></b></summary>
	<!-- __root__ Submodule -->
	<details>
		<summary><b>__root__</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ __root__</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/go.sum'>go.sum</a></b></td>
					<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/main.go'>main.go</a></b></td>
					<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/go.mod'>go.mod</a></b></td>
					<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
				</tr>
			</table>
		</blockquote>
	</details>
	<!-- pkg Submodule -->
	<details>
		<summary><b>pkg</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ pkg</b></code>
			<!-- model Submodule -->
			<details>
				<summary><b>model</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ pkg.model</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/pkg/model/finding.go'>finding.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<!-- internal Submodule -->
	<details>
		<summary><b>internal</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ internal</b></code>
			<!-- scanner Submodule -->
			<details>
				<summary><b>scanner</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.scanner</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/scanner/scan.go'>scan.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- analyzer Submodule -->
			<details>
				<summary><b>analyzer</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.analyzer</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/analyzer/scorer.go'>scorer.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/analyzer/rules.go'>rules.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- output Submodule -->
			<details>
				<summary><b>output</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.output</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/output/table.go'>table.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/output/color.go'>color.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- docker Submodule -->
			<details>
				<summary><b>docker</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.docker</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/docker/image.go'>image.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/internal/docker/container.go'>container.go</a></b></td>
							<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<!-- cmd Submodule -->
	<details>
		<summary><b>cmd</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ cmd</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/cmd/root.go'>root.go</a></b></td>
					<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/arasdenizhan/dockenv/blob/master/cmd/scan.go'>scan.go</a></b></td>
					<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
				</tr>
			</table>
		</blockquote>
	</details>
</details>

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** Go
- **Package Manager:** Go modules
- **Container Runtime:** Docker

### Installation

Build dockenv from the source and intsall dependencies:

1. **Clone the repository:**

   ```sh
   ❯ git clone https://github.com/arasdenizhan/dockenv
   ```

2. **Navigate to the project directory:**

   ```sh
   ❯ cd dockenv
   ```

3. **Install the dependencies:**

<!-- SHIELDS BADGE CURRENTLY DISABLED -->

    <!-- [![docker][docker-shield]][docker-link] -->
    <!-- REFERENCE LINKS -->
    <!-- [docker-shield]: https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white -->
    <!-- [docker-link]: https://www.docker.com/ -->

    **Using [docker](https://www.docker.com/):**

    ```sh
    ❯ docker build -t arasdenizhan/dockenv .
    ```

<!-- SHIELDS BADGE CURRENTLY DISABLED -->

    <!-- [![go modules][go modules-shield]][go modules-link] -->
    <!-- REFERENCE LINKS -->
    <!-- [go modules-shield]: https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white -->
    <!-- [go modules-link]: https://golang.org/ -->

    **Using [go modules](https://golang.org/):**

    ```sh
    ❯ go build
    ```

### Usage

Run the project with:

**Using [docker](https://www.docker.com/):**

```sh
docker run -it {image_name}
```

**Using [go modules](https://golang.org/):**

```sh
go run {entrypoint}
```

### Testing

Dockenv uses the {**test_framework**} test framework. Run the test suite with:

**Using [go modules](https://golang.org/):**

```sh
go test ./...
```

---

## Roadmap

- [x] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

## Contributing

- **💬 [Join the Discussions](https://github.com/arasdenizhan/dockenv/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/arasdenizhan/dockenv/issues)**: Submit bugs found or log feature requests for the `dockenv` project.
- **💡 [Submit Pull Requests](https://github.com/arasdenizhan/dockenv/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/arasdenizhan/dockenv
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com/arasdenizhan/dockenv/graphs/contributors">
      <img src="https://contrib.rocks/image?repo=arasdenizhan/dockenv">
   </a>
</p>
</details>

---

## License

Dockenv is protected under the [MIT LICENSE](https://github.com/arasdenizhan/dockenv?tab=MIT-1-ov-file) License. For more details, refer to the [LICENSE](https://github.com/arasdenizhan/dockenv?tab=MIT-1-ov-file) file.

---
