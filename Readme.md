# 🧠 Stack vs Heap Allocation in Go (Escape Analysis)

Этот репозиторий демонстрирует, как **escape analysis** влияет на производительность Go-кода. Мы сравниваем два варианта:

- `escape()` — переменная уходит в **heap**
- `noEscape()` — переменная остаётся в **stack**

## 🔬 Цель

Понять:
- Почему escape приводит к аллокациям в куче
- Как это замедляет код
- Как компилятор Go анализирует "жизнь" переменных
- Как измерить производительность с помощью бенчмарков

## 🧪 Benchmark результаты

```bash
go test -bench=. -benchmem -benchtime=100000000x
```

![alt text](https://github.com/orayew2002/stack_or_heap/blob/main/assets/result.png)
