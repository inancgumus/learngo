# Text Wrapper Challenge Guideline

In this project your goal is to mimic the soft text wrapping feature of text editors. For example, when there are 100 characters on a line and if the soft-wrapping is set to 40, an editor may cut the line that goes beyond 40 characters and display the rest of the line in the next line instead.

## EXAMPLE

Wrap the given text for 40 characters per line. For example, for the following input, the program should print the following output.

**INPUT:**

    Hello world, how is it going? It is ok.. The weather is beautiful.

**OUTPUT:**

    Hello world, how is it going? It is ok..
    The weather is beautiful.

## RULES

* The program should work with Unicode text. You can find a unicode text in [story.txt](story.txt) file.

* The program should not cut the words before they finish. Instead, it should put the whole word on the next line. 

For example, this is not OK:
 
    Hello world, how is it goi
    ng? It is o
    k. The weather is beautifu
    l.

## SOLUTION

* [Get the solution source code here](main.go).
