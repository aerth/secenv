# Secure Environment

Stores environmental variables in a nacl/secretbox (XSalsa20 and Poly1305)

## Usage

Create a config file:

```
secenv

```

Delete config file:

```
secenv -d

```

Run something with the variables:

```
secenv test.sh

```
