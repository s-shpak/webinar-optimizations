# Вебинар "Некоторые техники оптимизации Go-кода"

## `strconv` vs `fmt`

Запустим бенчмарки в `cmd/strconv_vs_fmt`:

```bash
go test -bench=.
```

Что наблюдаем и почему?

Соберем информацию по выделению памяти:

```bash
go test -bench=. -benchmem
```

Теперь соберем профиль CPU:

```bash
go test -bench=. -cpuprofile cpu.out
```

И запустим pprof для его визуализации (документация: https://github.com/google/pprof/blob/main/doc/README.md#interpreting-the-callgraph):

```bash
go tool pprof cpu.out
```

Как объяснить полученный результат в контексте бенчмарок?

Соберем профиль памяти:

```bash
go test -bench=. -memprofile mem.out
```

И запустим pprof для его визуализации:

```bash
go tool pprof mem.out
```

Что наблюдаем?

Объяснение устройства интерфейсов: см. слайды.

Убедимся в том, что `FormatIntSprintf` действительно аллоцирует память при передаче `i` в `fmt.Sprintf`:

```bash
go build -gcflags '-S -N' main.go &> main.s
```

Определение `convT64` находится в исходном коде Go в `src/runtime/iface.go`.


**Вопрос**: разобрав как устроены интерфейсы в Go, посмотрите на 9 строку файла `cmd/iface_puzzle/decl.go`. Как вы думаете, что она делает?

### Несколько слов об Escape Analysis

Запустим бенчмарки с дополнительным флагом:

```bash
go test -gcflags '-m' -bench=.
```

Разница между стеком и кучей: см. слайды.

## Использование указателей как ключей `map`

См. пример в `cmd/map_string_keys`.

Как изменятся измерения, если поменять тип ключа в мапе на `int`?

См. определение структуры `string` в исходниках Go в `src/runtime/string.go`.


## Определение емкости слайса

Перед запуском бенчмарок посмотрите на код в `cmd/slices_cap/main.go` и ответьте:

- какая из версий будет работать быстрее?
- будет ли у них одинаковое потребление памяти?

Запустим бенчмарки в `cmd/slices_cap`:

```bash
go test -bench=. -benchmem
```

Разберем устройство слайсов: см. слайды и `src/runtime/slice.go`.

Понаблюдаем как меняется емкость слайсов:

```bash
go test ./... -run=TestPrintSliceCapacityChanges -v
```

## Bounds-checking elimination

Попробуем запустить тест `TestSumFirstElementsOfSlice` из `cmd/bce`:

```bash
go test  ./... -bench=this-bench-does-not-exist -run=TestSumFirstElementsOfSlice
```

Теперь отключим bounds checking:

```bash
go test -bench=this-bench-does-not-exist -gcflags='-B' ./... -run=TestSumFirstElementsOfSlice
```

Что наблюдаем?

Посмотрим на разницу в ассемблерном коде:

```bash
go build -gcflags '-S -N' main.go &> main.s

go build -gcflags '-S -N -B' main.go &> main-no-bc.s
```

Проверим, что версия без проверок действительно быстрее:

```bash
rm ./*.s

go test -run=this-test-does-not-exist -bench="^BenchmarkSumFirstElementsOfSlice$"

go test -run=this-test-does-not-exist -gcflags="-B" -bench="^BenchmarkSumFirstElementsOfSlice$"
```

Go позволяет узнать, на какой строке происходит bound check:

```bash
go test -gcflags '-d=ssa/check_bce/debug=1' -run=this-test-does-not-exist  -bench="^BenchmarkSumFirstElementsOfSlice$"
```

Как мы видим bound check происходит внутри цикла. Можно ли это оптимизировать?

Запустим бенчмарки для новой функции:

```bash
go test -gcflags="-d=ssa/check_bce/debug=1" -bench="^BenchmarkSumFirstElementsOfSliceBCE$" -run=this-test-does-not-exist  -count=5
```

Посмотрите на код новой функции в `main_bce.go`.

## `sync.Pool`

Документация к `sync.Pool` находится здесь: https://pkg.go.dev/sync#Pool

Код примера находится в `sync-pool`.

Запустим пример без использования `sync.Pool`:

```bash
go test -run TestNoPool -memprofile mem.out
```

и с `sync.Pool`:

```bash
go test -run TestWithPool -memprofile mem-pool.out
```