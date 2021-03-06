## 模式定义
- Decouple an abstraction from its implementation so that the two can vary independently.
- 将抽象和实现解耦，使得它们可以独立地变化。
- 桥梁模式关注的是抽象和实现的分离，使得它们可以独立地发展；桥梁模式是结构型模式，侧重于软件结构。
- 策略模式关注的是对算法、规则的封装，使得算法可以独立于使用它的用户而变化；策略模式是行为型模式，侧重于对象行为。
- 设计模式其实就是一种编程思想，没有固定的结构。要区分不同的模式，要多从语义和用途的角度去判断。

## 应用场景
- 一个产品（或对象）有多种分类和多种组合，即两个（或多个）独立变化的维度，每个维度都希望独立进行扩展。
- 因为使用继承或因为多层继承导致系统类的个数急剧增加的系统，可以改用桥接模式来实现。

---
> [Bridge Design Pattern in Go (Golang)](https://golangbyexample.com/bridge-design-pattern-in-go/)