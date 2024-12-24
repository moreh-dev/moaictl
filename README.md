# moaictl

moaictl is an administrative CLI tool for managing the **moai-accelerator-manager**.

## Installation

To set up and use moaictl, follow these steps:

- Build the binary:
 
    ```go build -o moaictl main.go```
 
- (Recommended) to set as follows: 

    ```alias moaictl='go run {path to project}/main.go'```

    Alternatively, you may use a shorter alias

    ```alias m='go run {path to project}/main.go```

## Usage

Need to export ENV_MOAICTL_ROOT to your project directory.

To explore available commands and usage options, run:

```moaictl -h```
