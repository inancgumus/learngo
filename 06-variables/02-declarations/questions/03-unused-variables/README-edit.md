# Questions: Unused Variables #

---

## 1. What happens when you run a program with an unused variable in the block scope? ##

(A) Nothing, it will run just fine.

(B) The compiler will take automatically remove it.

(C) The program won't compile.

---

**Solution**: (C)

That's right. The compiler won't build the program if there are unused variables in the block scope.

---

## 2. What happens when you run a program with an unused variable in the package scope? ##

(A) Nothing, it will run just fine.

(B) The compiler will automatically remove it.

(C) The program won't compile.

---

**Solution**: (A)

That's right. Go will still compile the program if its a variable declared in the package scope.

---

## 3. What can you do to prevent the compiler from producing an unused variable error? ##

(A) You can change the variable's name.

(B) You can use a blank-identifier to discard the variable.

(C) You can add a comment telling Go to discard unused variables.

---

**Solution**: (B)

Correct. Using a blank-identifier (`_`) tells Go to discard the variable and prevents the compiler from producing an unused variable error.