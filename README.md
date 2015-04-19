markov
======

Markov chain generator written in Go

Original code found at http://golang.org/doc/codewalk/markov/

To run:

```
cat samples/grimms.txt BibleNIV.txt | ./markov --words 100 --prefix 2
