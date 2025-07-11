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


## 💡 Почему escape замедляет программу?
Когда переменная **"уходит" в кучу (heap)**, происходит следующее:

1. 📦 **Аллокация в куче (heap allocation)**  
   - Go вызывает аллокатор памяти для размещения объекта в куче.  
   - Это **медленнее**, чем выделение памяти в стеке.  
   - Вызывается функция в рантайме Go, например `runtime.newobject`.

2. 🧹 **Работа сборщика мусора (GC)**  
   - Всё, что находится в куче, должно **отслеживаться сборщиком мусора**.  
   - GC периодически проходит по куче, чтобы найти и освободить больше неиспользуемые объекты.  
   - Это тратит дополнительное процессорное время и ухудшает предсказуемость задержек.

3. 📈 **Кэш-память CPU работает хуже**  
   - Переменные в стеке обычно расположены последовательно (локальность данных), что позволяет CPU кэшу (L1, L2) работать эффективнее.
   - Переменные в куче могут располагаться в разных местах памяти, что снижает эффективность кэширования.

⚠️ **Disclaimer**

Это только мои тесты и сбор информации для изучения работы Go escape analysis.  
Не воспринимайте результаты как единственный источник истины — в реальных приложениях поведение и производительность могут отличаться в зависимости от архитектуры, версии компилятора и окружения.
