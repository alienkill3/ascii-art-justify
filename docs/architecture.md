### Comparing the Pipeline Approach with the Finite State Machine Approach

Both Pipeline and Finite State Machine (FSM) architectures are widely used for handling processes, especially in systems that require sequential operations or transformations. However, the choice between them depends on the complexity of the process, flexibility, and maintainability needs.
Here’s a comparison of the two approaches, including pros and cons of each, and a diagram of the pipeline approach.

### Pipeline Approach vs Finite State Machine (FSM) Approach

Pipeline Approach
The Pipeline approach treats data as flowing through a series of discrete stages or transformations. Each stage applies a specific transformation, and the data passes through these stages sequentially.
Example Use Case: Text transformation, data processing, etc., where each step (stage) performs a small transformation on the data.

### Finite State Machine (FSM) Approach

A Finite State Machine is a computational model with a limited number of states. The machine transitions from one state to another based on inputs and transitions, typically triggered by events or conditions. FSMs are often used for systems that need to handle different conditions or modes.

### Example Use Case: Systems requiring different modes or conditions, like a parser, authentication system, or process control.

### Comparison of Pipeline vs Finite State Machine (FSM) Approach

| **Criteria**                     | **Pipeline Approach**                                    | **Finite State Machine (FSM) Approach**                 |
|----------------------------------|----------------------------------------------------------|---------------------------------------------------------|
| **Use Case**                     | Linear data processing with clear, sequential stages.    | Systems with multiple states or modes that require specific transitions. |
| **Modularity**                   | High modularity. Each stage is independent and can be tested/modified separately. | Lower modularity as the system is more tightly coupled between states and transitions. |
| **Flexibility**                  | Flexible and easy to extend by adding/removing stages.    | Less flexible, as changing the state transition logic or adding states can be complex. |
| **Error Handling**               | Errors can be handled at each stage, but might require custom logic. | Centralized error handling can be tricky, as FSM relies on specific states and transitions. |
| **Complexity**                   | Simple to understand when stages are linear. Can get complex with many stages. | Can be complex to model, especially with many states or complex transitions. |
| **Scalability**                  | Easily scalable by adding more stages or operations in the pipeline. | Can become difficult to scale with too many states or transitions. |
| **Performance**                  | Potential performance issues due to overhead between stages, especially with large data. | Usually faster as it directly transitions between states without intermediate steps. |
| **Maintainability**              | High maintainability due to separation of concerns; easy to modify individual stages. | Can become hard to maintain as the number of states and transitions increases. |
| **Testability**                  | Easy to test each stage individually.                     | Testing can be harder because you need to verify the transitions and states, and their interactions. |
| **State Management**             | No explicit state management, as data flows through stages. | Explicit state management, making it ideal for cases where the system’s behavior depends on its state. |
| **Error Recovery**               | Error recovery can be handled within each stage.          | Error recovery requires defining fail states and potentially complex state transitions. |


### Decision Summury

- We choose: Pipeline approach
- Because: 
    1. Much Easier to Implement

    2. Linear data processing

    3. Easier Debugging

    4. Easy to Add Features

    5. Potential for Integration with Other Systems

    6. Following Industry Best Practices

