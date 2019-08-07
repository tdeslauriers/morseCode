Below is the challenge that inspired my encoder/decoder.  
It didnt really follow it.   I wanted to make something that I could actually use like my base64 encoder/decoder.

To the point about ambiguity: if you insert a delimiter between letters (in my case a space) you get correct/ consistent decoding.

______

[2019-08-05] Challenge #380 [Easy] Smooshed Morse Code 1
For the purpose of this challenge, Morse code represents every letter as a sequence of 1-4 characters, each of which is either . (dot) or - (dash). The code for the letter a is .-, for b is -..., etc. The codes for each letter a through z are:

.- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --..
Normally, you would indicate where one letter ends and the next begins, for instance with a space between the letters' codes, but for this challenge, just smoosh all the coded letters together into a single string consisting of only dashes and dots.

Examples
smorse("sos") => "...---..."
smorse("daily") => "-...-...-..-.--"
smorse("programmer") => ".--..-.-----..-..-----..-."
smorse("bits") => "-.....-..."
smorse("three") => "-.....-..."
An obvious problem with this system is that decoding is ambiguous. For instance, both bits and three encode to the same string, so you can't tell which one you would decode to without more information.
