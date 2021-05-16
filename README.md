# stock-notifier
Pushes stock price target notifications to Gotify

## Usage

### Define the required environment variables
```shell
# PUSH_SERVER_URL is the URL your Gotify instance is accepting connections on.
export PUSH_SERVER_URL="https://push.mydomain.io/message?token=atoken

# SYMBOLS is the symbols you want the stock-notifier to watch.
export SYMBOLS="ABC,DEF"

# <SYMBOL>_TARGETS is the price targets the stock-notifier should notify you for. 
# Use the '-' postfix for less than the target or the '+' postfix for more than the target.
export ABC_TARGETS="121+,80-" # Will notify you when the price for ABC reaches more than 121 or less than 80.
export DEF_TARGETS="42-,44+" # Will notify you when the price for DEF reaches more than 44 or less than 42.
```

### Build and run the stock-notifier
```shell
# Clone
git clone git@github.com:deifyed/stock-notifier.git && cd stock-notifier

# Build
go build main.go

# Run
./stock-notifier
```
