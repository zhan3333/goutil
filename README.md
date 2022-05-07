# 用范型来写一些工具函数与方法

go 版本最低需要 go1.18.1

goland 建议版本 2022.1, 该版本优化了很多泛型的提示。低于该版本 IDE 很有可能有语法报错。

官方示例: https://go.dev/doc/tutorial/generics

## 功能

### 调用方式

提供两种方式来操作切片:

1. 函数方式

```
import util "github.com/zhan3333/goutil"

util.Containes([]int{1, 2, 3}, 2) // true
```

- Contains
- ContainsAny
- ContainsAll
- Unique
- Map
- Reduce
- Filter
- Reject
- First
- Last
- Empty
- Merge
- Reverse
- Random
- Shuffle
- CountIf
- Diff
- Push
- Pop
- Sum
- Equal
- Sort

2. 结构体方式

结构体方式大部分函数是原地操作，返回原对象以支持链式调用。

```
import util "github.com/zhan3333/goutil"

s := util.NewSlice([]int{1, 2, 3})

s.Contains(2) // true

s.Push(4).Contains(4) // true

s.Push(5).MergeSlice([]int{6}).ContainsAll(5, 6) // true
```

- Slice.Set
- Slice.Contains
- Slice.Slice
- Slice.Unique
- Slice.Reset
- Slice.Each
- Slice.Map
- Slice.Reduce
- Slice.Filter
- Slice.Reject
- Slice.First
- Slice.Last
- Slice.Empty
- Slice.Merge
- Slice.MergeSlice
- Slice.Reverse
- Slice.Random
- Slice.Shuffle
- Slice.Index
- Slice.Copy
- Slice.ContainsCount
- Slice.Len
- Slice.Push
- Slice.Pop
- Slice.Equal
- Slice.Pretty
- Slice.JSON
- Slice.JSONString
- Slice.Diff

