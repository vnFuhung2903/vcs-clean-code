# vcs-clean-code
## Table of content
- [Design Pattern](#design-pattern)
- [SOLID](#solid)

## Design Pattern
Reusable solutions for typical software design challenges are known as design patterns. Expert object-oriented software engineers use these best practices to write more structured, manageable, and scalable code. Design patterns provide a standard terminology and are specific to particular scenarios and problems. Design patterns are not finished code but templates or blueprints only.
### Creational Design Pattern
Creational Design Patterns focus on the process of object creation or problems related to object creation. They help in making a system independent of how its objects are created, composed and represented. There are 5 types of Creational Design Patterns:
- **Singleton**: Ensures that each class only has 1 instance, while provides a global access to the instance
- **Factory Method**: Provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created
- **Abstract Factory**: Produces families of related objects without specifying their concrete classes
- **Builder**: Constructs complex objects step by step, allows to produce different types and representations of an object using the same construction code
- **Prototype**: Copies existing objects without making the code dependent on their classes

### Structural Design Pattern
Structural Design Patterns solves problems related to how classes and objects are composed/assembled to form larger structures which are efficient and flexible in nature. Structural class patterns use inheritance to compose interfaces or implementations. There are 7 types of Structural Design Pattern:
- **Adapter**: Allows objects with incompatible interfaces to collaborate
- **Bridge**: Splits a large class or a set of closely related classes into two seperate hierarchies - abstraction and implementation - which can be developed independently of each other
- **Composite**: Composes objects into tree structures, then work with these trees as if they were individual objects
- **Decorator**: Attaches new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors
- **Facade**: Provides a simplified interface to a library, a framework, or any other complex set of classes
- **Flyweight**: Fits more objects into available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object
- **Proxy**: Provides a substitute or a placeholder for another project. A proxy controls access to the original object, allows to perform either before or after the request gets through to the original object

### Behavioral Design Pattern
Behavioral Patterns are concerned with algorithms and the assignment of responsibilities between objects. Behavioral patterns describe not just patterns of objects or classes but also the patterns of communication between them. These patterns characterize complex control flow that’s difficult to follow at run-time. There are 10 types of Behavioral Design Pattern:
- **Chain of Responsibility**: Passes requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain
- **Command**: Turns a request into a stand-alone object that contains all information about the request. This transformation will pass requests as a method arguments, delay or queue a request's execution, and support undoable operations
- **Iterator**: Traverses elements of a collection without exposing its underlying representation (list, track, tree, etc.)
- **Mediator**: Reduces chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via mediator object
- **Memento**: Saves and restores the previous state of an object without revealing the details of its implementation
- **Observer**: Defines a subscription mechanism to notify multiple objects about any events that happen to the objects they're observing
- **State**: Lets an object alter its behavior when its internal state changes. It appears as if the object changed its class
- **Strategy**: Defines a family of algorithms, put each of them into a separate class, and make their objects interchangeable
- **Template Method**: Defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of algorithm without changing its structure
- **Visitor**: Separates algorithms from the objects on which they operate

### Other Design Pattern:
- **Dependencies Injection**: A technique which makes the interactions between objects as minimal as possible through specific dependencies

## SOLID
The SOLID principles are a set of guidelines for writing high-quality, maintainable, and scalable software.
### Single Responsibility Principle (SRP)
The Single Responsibility Principle (SRP) states that a class should have only one reason to change, or in other words, it should have only one responsibility. This means that a class should have only one job to do, and it should do it well.\
If a class has too many responsibilities, it can become hard to understand, maintain, and modify. Changes to one responsibility can inadvertently affect another responsibility, leading to unintended consequences and bugs. By following SRP, we can create code that is more modular, easier to understand, and less prone to errors.

### Open-Closed Principle (OCP)
The Open-Closed Principle (OCP) states that software entities (classes, modules, functions, and so on) should be open for extension but closed for modification. This means that the behavior of a software entity can be extended without modifying its source code.\
The OCP is essential because it promotes software extensibility and maintainability. By allowing software entities to be extended without modification, developers can add new functionality without the risk of breaking existing code. This results in code that is easier to maintain, extend, and reuse.

### Liskov Substitution Principle (LSP)
The Liskov Substitution Principle (LSP) states that any instance of a derived class should be substitutable for an instance of its base class without affecting the correctness of the program. In other words, a derived class should behave like its base class in all contexts.\
The importance of the LSP lies in its ability to ensure that the behavior of a program remains consistent and predictable when substituting objects of different classes. Violating the LSP can lead to unexpected behavior, bugs, and maintainability issues.

### Interface Segregation Principle (ISP)
The Interface Segregation Principle (ISP) focuses on designing interfaces that are specific to their client's needs. It states that no client should be forced to depend on methods it does not use.\
The principle suggests that instead of creating a large interface that covers all the possible methods, it's better to create smaller, more focused interfaces for specific use cases. This approach results in interfaces that are more cohesive and less coupled.

### Dependency Inversion Principle (DIP)
The Dependency Inversion Principle (DIP) states that high-level modules should not depend on low-level modules, but both should depend on abstractions. Abstractions should not depend on details – details should depend on abstractions.\
This principle aims to reduce coupling between modules, increase modularity, and make the code easier to maintain, test, and extend.