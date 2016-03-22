# Secure Environment

Stores environmental variables in a nacl/secretbox (XSalsa20 and Poly1305).

## Situations to use secenv:

  * *Untrusted peer users*
  * Trusted machine + network (out of our control?)
  * Trusted administrator(s) (root can read /proc/ and copy your files to reverse engineer)
  * If your program requires hard-to-type environmental variables at runtime (such as AWS or other API keys)
  * If you don't need to automatically reboot the server (you need to unlock the config every time it runs)

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
