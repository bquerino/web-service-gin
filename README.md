# **Overkill - Hexagonal architecture with GoLang**

This project aims to use the hexagonal architecture way to do things with GoLang.

To create this project I used the official documentation of GoLang that can be found here: 
<https://golang.org/doc/tutorial/web-service-gin>

## **Layers**

In this section, I'll try to explain the layers of this project and how you can evolve from this point.

### Core Layer

This layer basically refers to a component who is agnostic about technology and contains all the business logic.
The core don't have responsibility or need to be aware of how the application can be served or where the data will be persisted.

### Driver adapters aka Primary actors

This layer is where we have everything who trigger the communication with the core.

Valid examples:

* HTTP interface(REST)
* Command line interface
* Application UI

### Driven adapters aka Secondary actors

This layer is where we expect the core to be the responsible to trigger the communication. 

Valid examples:

* Queue
* Databases
* Cache

### Ports

This layer basically is where we define the set of actions what our actors needs to implement.

## To do

* Implement tests
* Implement more adapters
* Update the documentation with images

## References

* Hexagonal architecture - <https://alistair.cockburn.us/hexagonal-architecture/>
* Alistair in the "Hexagone" 1/3 - <https://www.youtube.com/watch?v=th4AgBcrEHA&t=87s>
* Hexagonal architecture in Go - <https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3>
