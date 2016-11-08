MATH 221 FALL 2016

HOMEWORK 10 / LAB 10: Go language warm-up
Due: Tuesday, November 8th, 2016


For this assigmnent, you can follow the three examples of Go
programs I've placed in this folder:
   fib.go: computes all things Fibonacci (loops and recursion)
   pow.go: computes all things power (loops, recursion, modular arithmetic)
   search.go: sorts a Go array and uses it for binary search (arrays)
You'll notice that each of these Go programs lives in its own folder. This
is preferred by the Go system to have each program run on its own.  I've
included, in this folder, three empty folders "ex0", "ex1", and "ex2" where 
you can place your solutions to the exercises below (e.g. "ex1.go").

To run any of these, change to their folder
  cd fib
and then execute them with the "go run" command
  go run fib.go

For these and the exercise code that you write, you should only need to use 
the syntax of definitions, statements, and expressions that I demonstrated
in class today, and that I will demonstrate in class on Wednesday (in the 
case of arrays). The key components of the Go language that they each use 
include:

PROCEDURE and FUNCTION DEFINITIONS

These appear at the top level within your source code and are of the syntactic form

  func <proc-name> (<param-name> <type>, ...) {
      <block of statements>
  }

  func <func-name> (<param-name> <type>, ...) <return type> {
      <block of statements, including returns>
  }

and should include definition of a procedure called 'main':

  func main() {
      <statements that constitute your program>
  }

Please do decompose your Go programs into useful and meaningful helper
procedures and functions.  Don't just put everything in 'main'.


STATEMENTS

Your code block consists of a series of statements.  These could be 
variable declarations:

  var <name> <type>

or variable declarations with initialization:

  var <name> <type> = <expresssion>

Where <type> is typically "string", "int", or "bool".  They could 
instead be conditional statements, of the "if-then" form

  if <expression> {
      <statements>
  } 

or the "if-then-else" form:

  if <expression> {
      <statements>
  } else {
      <statements>
  }

You can "cascade" the statment like so:

  if <expression> {
      <statements>
  } else if <expression> {
      <statements>

  ...

  }


Your block statements could be variable assignment statements 
of the form

  <name> = <expression>

or a "return" statement that exits a procedure or a function.
If it's a function, then the form is

  return <expression>

and if it is a procedure (does not return a value) then the form
is simply

  return 

Go's looping statement, essentially the same as Python's "while",
is the "for" loop:

  for <expression> {
      <statements>
  }

And you can either call your own procedures like

  <name>(<expression>, ...)

or exported procedures from a package with

  <package>.<Name>(<expression>, ...)


INTEGER AND BOOLEAN OPERATIONS

For expressions, you can use the same arithmetic, comparison, and logical
operations that you know from other languages, just the Go variants:

   && ||     logical AND, OR for booleans
   == !=     equality comparison
   < > <= >= ordering comparison of integers
   + - * / % integer arithmetic (including div and mod)
   << >>     integer shift left, right
   & | ^     bitwise AND, OR, XOR of integers
   -         unary integer negation
   !         unary boolean negation

For strings, you can use 

   == !=     for string equality
   +         for string concatenation

COMMENTING YOUR CODE

The Go language allows end of line comments, like so

   // This is an end of line comment.

and block comments (maybe spanning several lines) like 
so

   /* This

      is a block
  
          comment.
    */

Please comment all your functions and any tricky code you write.
Give an overall comment about your code at the top of the file.


INPUT AND OUTPUT

Using the line near the top of your file:

  import ( "fmt" )

would allow you to print formatted input, which could either
be a fixed string:

  fmt.Printf("Hello!\n")

or an integer:

  fmt.Printf("The integer result is %d.\n", result)

or a string:

  fmt.Printf("I'm sorry, %s. I'm afraid I cannot 

do that.\n", name)

For input, we rely on the Scanf procedure in "fmt" and the "Atoi"
function in "strconv".  You can do the following:

  fmt.Printf("Please enter an integer.")
  fmt.Scanln(&s)
  var n int = 0
  n, _ = strconv.Atoi(s)

Code like This will make your Go programs interactive within Terminal.

----------------------------------------------------------------------

==> Exercise 0. As a warm-up, write a program that mimics my pow.go
and my fib.go and computes the factorial of an input integer. It
should compute it in three ways, each with its own function:

1. Iteratively, with a loop. Write a function 'factit' that uses a
loop to compute and return the factorial of an integer.

2. Iteratively using a tail recursive function. Write a one-parameter
function 'facttr' that relies on a two (or more) parameter recursive
helper function 'facttrh' that computes the factorial of a number,
passing along a product.

3. Recursively using normal recursion. Write a single-parameter recursive
function named 'factrec' that computes and returns the factorial.

Have your 'main' mimic mine for 'pow' and 'fib'. It should take an
integer as input, and output the factorial returned by each of the three
functions. 


==> Exercise 1. Write a program that requests an integer and prints
all the positive integers less than or equal to that integer whose bit
representation is a palindrome.  You'll want to borrow the code from
my examples for reading a string and converting it to an integer, and
for printing the decimal representation of an integer.  I want you to
build bit-checking code from scratch using modular arithmetic or using
the bit operations. That is, I don't want you to use built-in Go
packages for getting the binary representation.

I have two suggestions for writing this code:

a. Check the highest and lowest order bits.  Then the next highest and
next lowest order bits. Etc. You may want to write a function that 
finds the largest power of two that does not exceed the integer, in
order to do that.  You might, though it is not necessary, write a
helper function that returns a bit of an integer at a particular 
position.

b. Build two strings of the bits, one that lists the bits in order,
and the other the lists the bits in reverse order.  You can use the
string append operation +, for example like so

  s = s + "0"

to do this work.  Having built these two strings, compare them to see 
if they are equal.



==> Exercise 2.  Blackjack is a card game that can be played against a
"dealer".  A fifty-two card deck is shuffled at random.  The cards in
the deck are ranked 2 through 10, jack, queen, king, and ace, thirteen
total.  But there are four copies of each set of thirteen, each labelled
with an additional suit: hearts, diamonds, spades, and clubs.  That
makes a total of 4 x 13 = 52 cards. (E.g. the "ace of spades")

The dealer lays down cards, one at a time, building a "hand" for
itself and for you, the player.  It then lays down a card for you,
face up. That is, revealing the card's rank and suit.  It lays down
one card for itself, face down, that is, the rank and suit of the card
is not revealed.  It lays down a second for itself, face up.  It lays
down a second for you, face up.

Game play involves totalling the cards in your hand, and guessing the
total of the hand of the dealer.  The total contributed by a card is 
either the card's rank, if numbered 2 through 10, 10 if it is a jack,
queen, or king, and either 1 or 11 if it is an ace. Your goal is to 
have a higher total, but not one over 21.  When dealt an ace, a player
has the choice of making it count as 1 or as 11.  

Having set the initial totals with two cards apiece, play continues
with a series of rounds.  With each round, you can either "hit" or 
"stand". If you choose "hit", the dealer draws a card from the deck
and lays it face up on your hand.  You can continue hitting as much
as you like, and then stop with a "stand".  If, with any hit, you
go over 21, then you go "bust" and lose the game.

If you have not gone bust, then the dealer goes through a similar
round of hits. Having seen your cards, the dealer is forced to
follow a fixed strategy:  if its total is under 17, it must hit.
Otherwise it stands, or maybe has gone bust.

Write a Go program that simulates this gameplay. You can use the
(pseudo)random number generator package:

   import ( "math/rand" )

which allows you to write code like:

   r = rand.Intn(52)

to choose a random number between 0 and 51.  This will 
always produce the same sequence of integers, so you will 
want to use a clock-generated "seed" with the procedure
invocation

   rand.Seed(time.Now().UnixNano())

which means that you'll want to import the package "time"
as well.

To do the shuffling of the deck (you don't want cards to repeat),
you'll want to mimic my code in search.go to build an array of
the 52 cards.  Something like:

  var deck []int = make([]int,52)

and then a loop that sets each card in the deck.  Then you'll want 
to reorder the array randomly by a method of your choosing using 
array element swaps:

   deck[i],deck[j] = deck[j],deck[i] 

