# Split Strings
**6 kyu**  

Завершите решение так, чтобы оно разбило строку на пары из двух символов. Если строка содержит нечетное количество символов, то она должна заменить отсутствующий второй символ последней пары символом подчеркивания ('_').

<details>
<summary>English (original)</summary>
Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
</details>  

Examples:
``` go
'abc' =>  ['ab', 'c_']
'abcdef' => ['ab', 'cd', 'ef']
```