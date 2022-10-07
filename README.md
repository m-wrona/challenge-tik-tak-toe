# challenge-tik-tak-toe
Coding challenge - tic-tac-toe game 

## Description 

Sample implementation of Tic-Tac-Toe game.
Implementation is based on [MinMax](https://www.neverstopbuilding.com/blog/minimax)
but with some simplifications & heuristic to minimize resource consumption.

#### Time spent

~4 hours of pure coding (perhaps a bit longer). Main piece was ready after ~2h. 
Rest of time was spent for polishing (project structure, adding additional tests, bug fixes etc.).
And polishing consumes always most of the time!
Analysis of possible algorithms & approaches, sketching API not included. 

## Project structure

Standard structure described [here](https://github.com/golang-standards/project-layout) has been chosen.

It allows to create many binaries (one with actual game, the other ones to test AI strategies etc.).

Moreover, introducing few packages allows better encapsulation between components and pieces of functionality. 

## Testing

Normally full pyramid of tests should be present but, since it's just a sample, unit tests are present for now.

Unit tests at this stage allow to test all what's required.

## Future improvements

* pretty printing
  * make sure that shell screen doesn't scroll down
  * print `X|O` instead of player IDs
* more AI strategies that can be compared 