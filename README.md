# GoFin

GoFin is a command-line application designed to fetch and display stock data from the [Twelve Data API](https://twelvedata.com/). It presents this information in a clean, full-screen Text-based User Interface (TUI) for easy viewing.

## Features

- Fetches stock quotes for a predefined list of symbols.
- Displays stock data in a sortable and navigable table within a TUI.
- Responsive full-screen TUI built with Bubble Tea.
- Clear presentation of key stock metrics like price, change, volume, and more.

## Setup

Before running GoFin, you need to obtain an API key from [Twelve Data](https://twelvedata.com/).

1.  **Sign up** at [Twelve Data](https://twelvedata.com/) to get your free API key.
2.  **Set Environment Variable**:
    The application requires the `TWELVE_DATA_API_KEY` environment variable to be set. You can set it in your shell:
    ```bash
    export TWELVE_DATA_API_KEY="YOUR_API_KEY_HERE"
    ```
    Replace `"YOUR_API_KEY_HERE"` with the actual API key you obtained. To make this permanent, add this line to your shell's configuration file (e.g., `.zshrc`, `.bashrc`, or `.profile`).

## Prerequisites

- Go (version 1.23 or later recommended)
- Git (for cloning the repository)

## Installation & Running

1.  **Clone the repository**:

    ```bash
    git clone https://github.com/bm611/gofin.git
    # Or your appropriate repository URL
    cd gofin
    ```

2.  **Run the application**:
    Ensure your `TWELVE_DATA_API_KEY` environment variable is set. Then, from the project's root directory (`gofin`), run:
    ```bash
    go run main.go
    ```
    The application will start, fetch data for a default set of stock symbols (e.g., AAPL, MSFT, GOOGL, AMZN), and display it in the TUI.

## TUI Usage

Upon launching, GoFin presents a full-screen table displaying stock data.

**Keybindings:**

- **Navigate Table**:
  - `↑` or `k`: Move the selection up by one row.
  - `↓` or `j`: Move the selection down by one row.
- **Toggle Help**:
  - `h`: Show or hide the help text in the footer, which lists available keybindings.
- **Quit**:
  - `q` or `Ctrl+C`: Exit the application.

The table columns include:

- Symbol
- Name
- Price
- Change
- Change%
- Open
- High
- Low
- Volume
- Market (Status: Open/Closed)

## Dependencies

GoFin utilizes the following major Go libraries:

- [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - For the Text-based User Interface.
- [github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - For TUI styling.
- [github.com/spf13/cobra](https://github.com/spf13/cobra) - For command-line interface structure.
