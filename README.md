## Programming Assignment: Hangman Console Game in Go

### Objective:

Develop a classic Hangman game using Go (Golang) for the console. The game's primary goal is to guess a hidden word by suggesting letters, one at a time. The game will only display the number of incorrect guesses made.

**Core Features:**

1.  **Word Selection:**
    *   The game should randomly select a word from a predefined list or a text file.
    *   Consider using different categories of words to add variety (e.g., animals, countries, food).

2.  **Guess Input:**
    *   Allow the user to input a single letter at a time.
    *   Validate input to ensure it's a letter and hasn't been guessed before.

3.  **Game Logic:**
    *   If the guessed letter is in the word, reveal its positions in the word.
    *   If the letter isn't in the word, increment the wrong guess counter.
    *   Display the current state of the word (e.g., "_ _ _ _ O") and the wrong guess count.

4.  **Winning/Losing Conditions:**
    *   If the user guesses all the letters in the word, they win!
    *   If the wrong guess count reaches a maximum limit (e.g., 10), they lose.

5.  **Console Display:**
    *   Use clear formatting to show the word's progress, wrong guesses, and any messages.
    *   Consider using colors or symbols to enhance the visual feedback.

**Example Output:**

```
Word: _ _ _ _ _
Wrong guesses: 0
Guess a letter: a

Word: _ a _ a _
Wrong guesses: 0
Take a guess: e

Word: _ a _ a _
Wrong guesses: 1
...
```

**Submission:**

Fork this repository and create a pull request to this repository with your changes.

Let [me](mailto:paal.gyula@rackforest.hu) know if you have any questions!
