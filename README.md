## 创建型: 解决对象的创建
 
**01 单例模式（Singleton）**

单例模式用来创建全局唯一的对象


**02 工厂模式(Factory）**


工厂模式用来创建不同但是相关类型的对象（继承同一父类或者接口的一组子类），由给定的参数来决定创建哪种类型的对象


**03 抽象工厂模式(Abstract Factory)**


**04 建造者模式(Builder)**

建造者模式是用来创建复杂对象，可以通过设置不同的可选参数，“定制化”地创建不同的对象


**05 原型模式（Prototype 不常用）**

原型模式针对创建成本比较大的对象，利用对已有对象进行复制的方式进行创建，以达到节省创建时间的目的


## 结构型: 解决类或对象的组合或组装

**06 代理模式(Proxy)**



**07 桥接模式(Bridge)**



**08 装饰者模式(Decorator)**



**09 适配器模式(Adapter)**



**10 门面模式（Facade 不常用）**



**11 组合模式（Composite 不常用）**



**12 享元模式（Flyweight 不常用）**



## 行为型: 解决类或对象之间的交互

**13 观察者模式(Observer)**



**14 模板模式(Template Method)**



**15 策略模式（Strategy）**



**16 职责链模式(Chain of Responsibility)**



**17 迭代器模式(Iterator)**



**18 状态模式(State)**



**19 访问者模式（Visitor 不常用）**



**20 备忘录模式（Memento 不常用）**



**21 命令模式（Command 不常用）**

![](http://www.plantuml.com/plantuml/png/X96nRjH048RxVOefan1VH6kTB9I8WWOa10-mNi_bbgntpEpum0Jw22b854WKDAIYIWiDRuCZ7eFOjsiT1a6rdfb_lfb_V2-8PcbhfupJ8S29i-W5pjKRvKaGX1gWLKpdq5zVtLny1T5d-8WhP1NR7lWDdI7VzW406AfhxIiusRvZxWxGqF2qsTkFSK29WhtOi3-R6gavt7owxDz-xjz_J12bh1VabJQOfi1vvE9BDAtWtaGhIOJg0HYbEVoylkelFwXachnX5FoTn6ZGRf3XSVpuax-xHb9DUsWZblpVraZI7FftN_fl5_Cz_eLIQcqh1AmRwQ36kaGEIiNVbXNfPYg-S91boNRA3RAPPJMBOR5uNKnzyxssrHIfAPrb1_y3xywU06fx81tWuTtHyR8ibeOSSSxKQNViMYmYO7aq5ehbKLaaxPEkmNn8m36kzSOIQpVMxdkZcz0wFT21yNcBGOOvsWELpz38uQYJDJAM7JWKiVuKEcgXqQnh56Jx2i6CzmsZ-8MLDLHsjK96BuaQ3k7sRH68mT7feDmf3Sv3spJ40gslI9SEWHfaFVpOi539oe_llteeczGEOI2Ax_-dwcLwZvWMMcbKbBH-OB62gIvgTIMj7tWOdKbSzGeT2GXP_F9u8UybCBcM3KG2Hdb2FcBQNLqnwKZ4TuL-3Sy3XCi-LJfxHBSsfVYR1NOIFZa1p1nRjJ47fi6i1G7wqP4KrcaFdc9K7P8TJnuusIL2cm1lhgsUKilckuVb5MC5QwnRbnPTPDOKNO-2MGys9So1ULWW254e_VAr_p2CF_PHs6gK5l8Uttutnc_tz2pdrz_kY6W8_GC0)

`当需要将发出请求的对象和执行请求的对象解耦的时候，使用命令模式`

将请求封装成对象，这可以让你使用不同的请求、队列或日志请求来参数化其他对象。命令模式也可以支持撤销操作。



**要点**

- 命令模式将发出请求的对象和执行请求的对象解耦
- 在被解耦的两者之间是通过命令对象进行沟通的。命令对象封装了接收者和一个或一组动作。
- 调用者通过调用命令对象的execute()发出请求，这会使得接收者的动作被调用。
- 调用者可以接受命令当作参数，甚至在运行时动态地进行。
- 命令可以支持撤销，做法是实现一个undo()方法来回到execute()被执行前的状态。
- 宏命令是命令的一种简单的延伸，允许调用多个命令。宏方法也可以支持撤销。
- 实际操作时，很常见使用“聪明”命令对象，也就是直接实现了请求，而不是将工作委托给接收者。
- 命令也可以用来实现日志和事物系统。


**22 解释器模式（Interpreter 不常用）**



**23 中介模式（Mediator 不常用）**

