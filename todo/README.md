# Purpose

This project is solely to be used as an example of a simple API.  I'll
be using it, in conjunction with another project, to demonstrate how
to wrap an existing API with a set of hypermedia wrappers.

Although I cannot imagine someone taking a serious interest in this
code, I just want to be explicit: you almost certainly don't want to
use this code for anything.

This example was at least partially inspired by
[Mike Amundsen's work along similar lines](https://github.com/LCHBook/todo-hyper).


# Operations

The following is a list of operations that should be supported:

  * ListAll - list the IDs of all todo items
  * ListActive - list the IDs of all uncompleted todo items
  * ListCompleted - list the IDs of all completed todo items
  * Get - Get title and status of item
  * Add - add a todo item
  * Edit - edit a todo item
  * Remove - remove a todo item
  * MarkCompleted - mark a given item as completed
  * ClearCompleted - mark a given item as uncompleted
