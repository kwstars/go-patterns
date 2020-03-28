## 什么是策略模式
- Define a family of algorithms,encapsulate each one,and make them - interchangeable.Strategy lets the algorithm vary independently from the clients that use it.
- 定义一系列算法，将每个算法都封装起来，并且使它们之间可以相互替换。策略模式使算法可以独立于使用它的用户而变化。

## 应用场景
-  如果一个系统里面有许多类，它们之间的区别仅在于有不同的行为，那么可以使用策略模式动态地让一个对象在许多行为中选择一种。
-  一个系统需要动态地在几种算法中选择一种。
-  设计程序接口时希望部分内部实现由调用方自己实现。