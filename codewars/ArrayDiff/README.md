# Array.diff
**6 kyu**  

Ваша цель в этом ката - реализовать функцию разности, которая вычитает один список из другого и возвращает результат.

Он должен удалить все значения из списка a, которые присутствуют в списке b, сохраняя их порядок.

``` go
array_diff({1, 2}, {1}) == {2}
```
Если значение присутствует в b, все его вхождения должны быть удалены из другого:
``` go
array_diff({1, 2, 2, 2, 3}, {2}) == {1, 3}
```

<details>
<summary>English (original)</summary>
Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.

It should remove all values from list a, which are present in list b keeping their order.
``` go
array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
```
If a value is present in b, all of its occurrences must be removed from the other:
``` go
array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)
```
</details>  