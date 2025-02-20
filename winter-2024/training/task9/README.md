# Анализ игрового поля

## Условие задачи

Игровое поле представляет собой белый прямоугольник $n \times m$, на котором изображены черные прямоугольные рамки. Толщина каждой рамки равна $1$, рамки не пересекаются и не касаются. Таким образом, для любых двух рамок $a$ и $b$ верно:

$\bullet$ либо $a$ вложена в $b$,

$\bullet$ либо $b$ вложена в $a$,

$\bullet$ либо $a$ не вложена в $b$ и одновременно $b$ не вложена в $a$.

Пример возможного поля изображён ниже. Белые символы обозначены точками ($ {.}  $), чёрные — звёздочками ($ {*} $).

${*********......}$
${*.......*..****}$
${*.***...*..*..*}$
${*.*.*...*..****}$
${*.***...*......}$
${*.......*......}$
${*...***.*......}$
${*...*.*.*.***..}$
${*...*.*.*.*.*..}$
${*...*.*.*.***..}$
${*...***.*......}$
${*.......*......}$
${*.......*......}$
${*********......}$

${...............}$
${.**************}$
${.*............*}$
${.*.*********..*}$
${.*.*.......*..*}$
${.*.*.*****.*..*}$
${.*.*.*...*.*..*}$
${.*.*.*****.*..*}$
${.*.*.......*..*}$
${.*.*********..*}$
${.*............*}$
${.**************}$

Для каждой рамки найдите количество рамок, в которые она вложена. Выведите получившиеся $r$ чисел в порядке неубывания, где $r$ — количество рамок на поле.

Например, для поля выше результат имеет вид: $0, 0, 0, 0, 1, 1, 1, 2$ (четыре рамки не вложены в какие-либо другие, три рамки вложены в одну, одна рамка вложена в две).

## Входные данные

В первой строке входных данных записано целое число $t$ ($1 <= t <= 1000$) — количество наборов входных данных.

Наборы входных данных в тесте являются независимыми. Друг на друга они никак не влияют.

Первая строка каждого набора входных данных содержит пару целых чисел $n, m$ ($3 <= n,m <= 2000$) — количество строк и столбцов на поле.

Далее следуют $n$ строк по $m$ символов в каждой строке. Каждый символ это либо $ {.}  $ (пустая белая клетка), либо $ {*}  $ (чёрная клетка, то есть часть рамки).

Гарантируется, что каждая рамка имеет прямоугольную форму, толщина каждой рамки равна $1$. Ширина и высота каждой рамки не меньше $3$. Никакие две рамки не пересекаются и не касаются по стороне или углу. На поле есть хотя бы одна рамка.

Гарантируется, что суммарный размер (площадь) всех полей в тесте не превосходят $4 \cdot 10^6$.

## Выходные данные

Для каждого набора входных данных выведите строку из целых чисел — глубины вложенностей рамок в отсортированном по неубыванию порядке.

## Пример теста 1

### Входные данные

```bash
3
26 15
*********......
*.......*..****
*.***...*..*..*
*.*.*...*..****
*.***...*......
*.......*......
*...***.*......
*...*.*.*.***..
*...*.*.*.*.*..
*...*.*.*.***..
*...***.*......
*.......*......
*.......*......
*********......
...............
.**************
.*............*
.*.*********..*
.*.*.......*..*
.*.*.*****.*..*
.*.*.*...*.*..*
.*.*.*****.*..*
.*.*.......*..*
.*.*********..*
.*............*
.**************
15 15
***************
*.............*
*.***********.*
*.*.........*.*
*.*.*******.*.*
*.*.*.....*.*.*
*.*.*.***.*.*.*
*.*.*.*.*.*.*.*
*.*.*.***.*.*.*
*.*.*.....*.*.*
*.*.*******.*.*
*.*.........*.*
*.***********.*
*.............*
***************
3 4
***.
*.*.
***.

```

### Выходные данные

```bash
0 0 0 0 1 1 1 2 
0 1 2 3 
0 

```
