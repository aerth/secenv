# Secure Environmental Variable Storage

Stores environmental variables in a nacl/secretbox (XSalsa20 and Poly1305) for later.

Note: While running the app using ecenv yourapp.sh, the variables are going to be unencrypted (plaintext) somewhere in /proc/ directory. Just like typing `KEY=123 run.sh` This would be possible to view if one is owner of the machine, while it is running. 

If you encounter any issues, report them on github.com/aerth/secenv/issues/new

Pull requests are welcome at github.

## Works with:

AWS, Mandrill, Sendgrid, Heroku, and on and on...

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
