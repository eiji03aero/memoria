# Backend
## Components
### usecase
- Represents actual usecases that happens on the business domain.
- In other words, we shouldn't implement methods that do not seem to be fit for that granularity
    - example) Find a user that fits for certain usecase.

### domain/service
- Implements executable logic that has contains domain logic.

### infra/dao (domain/interfaces/repository)
- Implements any database retrieval / manipulation logic here.
    - Querying data
    - Create / Update / Delete data
- Any of our application code calls these methods to get to database.

## Remarks
### Implement query logic in dao or domain service?
- Sometimes this topic makes you wonder.
- There are two approaches:
    - 1: Try to abstract query builder part
        - This is to basically abstract the query building logic from dao, and move it to domain service.
        - So that the implementation will be like domain service to have the actual domain logic.
        - But abstracting is not a piece of cake
            - example) how do we use advanced rdbms features like join, with clause, group by in domain service?
    - 2: Implement logic in dao and call it a day
        - This is to give up the first approach and put the querying logic in dao.
- hmmm querying data is not a domain logic ... ?
- more like a applicatioin logic? like how to prepare for executing domain logic?
- well any way, I have decided to go with the #2 for now:
    - #1 will be very difficult (for abstracting quering logic), while its advantage is only the conceptual organization
    - #2 gives dev speed as well, plus it is more intuitive for dao to have that logic as that's where people look for db logic.
    - Plus, dao is supposed to implement interfaces defined in domain layer, so we can kind of call it like owned by domain layer

### Implement logic in usecase or domain service?
- This can be tricky as well, but let's have following as our rule:
    - When not sure, think if it's good fit for usecase first.
        - is it a good chunk of usecase for our application?
        - is it just an application logic that does not have domain layer logic?
    - if it fits, let's go with usecase
    - if it didn't, think if it's a fit for domain service
    - else, probably it is a time to introduce another coneptual module
