# Brief Explanation
It's a little program that implements and showcases a very simple hand-shaking (key exchange) protocol based on tcp-ip

# Why?
I'm developing a game server for a text-based mmorpg with the help of another, very bright developer; and for the authentication part, we're gonna need some security.

# What?
A tcp-based protocol that exchanges public rsa-4096bit keys.

# How?
1- generate a 4096bit RSA key
1- send a JSON version of it to the peer
1- check if the peer has responded with "KEY OK"
  - if it has not, then stop the handshake
1- generate a random sequence of characters with length of ten
1- send the sequence to the peer
1- receive a buffer from the client
1- try to decipher the buffer using your private key
  - if you cannot, then stop the handshake
1- check if the decipher matches the sequence that was generated at step 4
  - if not, then stop the handshake
