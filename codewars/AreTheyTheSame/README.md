# Are they the "same"?
**6 kyu**  

(*Что-то не так с тестами*)

Даны два массива `a` и `b` написана функция `comp(a, b)` (или `compSame(a, b)`), которая проверяет, имеют ли два массива "одинаковые" элементы с одинаковыми кратностями (кратность элемента - это количество раз, когда он появляется). "Одинаковый" здесь означает, что элементы в `b` являются элементами в `a` квадрате, независимо от порядка.

Примеры
Допустимые массивы
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
```
`comp(a, b)` возвращает `true`, потому что в `b` 121 - это квадрат из 11, 14641 - это квадрат из 121, 20736 - квадрат из 144, 361 - квадрат из 19, 25921 - квадрат из 161 и так далее. Это становится очевидным, если мы запишем `b` элементы в терминах квадратов:
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11] 
b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
```
Недопустимые массивы
Если, например, мы изменим первое число на что-то другое, `comp` больше не возвращается значение `true`:
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
```
`comp(a, b)` возвращает `false`, потому что в `b` 132 не является квадратом какого-либо числа `a`.
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
```
`comp(a,b)` возвращает `false`, потому что в `b` 36100 не является квадратом какого-либо числа `a`.

Примечания
`a` или `b` может быть `[]` or `{}` (все языки, кроме R, Shell).
`a` или `b` может быть `nil` или `null` или `None` или `nothing` (за исключением C ++, COBOL, Crystal, D, Dart, Elixir, Fortran, F #, Haskell, Nim, OCaml, Pascal, Perl, PowerShell, Prolog, PureScript, R, Racket, Rust, Shell, Swift).
Если `a` или `b` являются `nil` (или `null` или `None`, в зависимости от языка), проблема не имеет смысла, поэтому верните `false`.

Примечание для C
Два массива имеют одинаковый размер, `(> 0)` заданный в качестве параметра в функции `comp`.

<details>
<summary>English (original)</summary>

Given two arrays `a` and `b` write a function `comp(a, b)` (or `compSame(a, b)`) that checks whether the two arrays have the "same" elements, with the same multiplicities (the multiplicity of a member is the number of times it appears). "Same" means, here, that the elements in `b` are the elements in `a` squared, regardless of the order.

Examples
Valid arrays
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
```
`comp(a, b)` returns true because in `b` 121 is the square of 11, 14641 is the square of 121, 20736 the square of 144, 361 the square of 19, 25921 the square of 161, and so on. It gets obvious if we write `b`'s elements in terms of squares:
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11] 
b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
```
Invalid arrays
If, for example, we change the first number to something else, `comp` is not returning true anymore:
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
```
`comp(a,b)` returns false because in `b` 132 is not the square of any number of `a`.
``` go
a = [121, 144, 19, 161, 19, 144, 19, 11]  
b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
```
`comp(a,b)` returns false because in `b` 36100 is not the square of any number of `a`.

Remarks
`a` or `b` might be `[] or {}` (all languages except R, Shell).
`a` or `b` might be `nil` or `null` or `None` or `nothing` (except in C++, COBOL, Crystal, D, Dart, Elixir, Fortran, F#, Haskell, Nim, OCaml, Pascal, Perl, PowerShell, Prolog, PureScript, R, Racket, Rust, Shell, Swift).
If `a` or `b` are `nil` (or `null` or `None`, depending on the language), the problem doesn't make sense so return false.

Note for C
The two arrays have the same size `(> 0)` given as parameter in function `comp`.
</details>  
