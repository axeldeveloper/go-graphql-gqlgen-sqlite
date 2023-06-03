# Project
- golang
- sqlite


# Teste

```js
    query {
        todos {
            id
            text
            user {
            id
            name
            }
        }
    }



    query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}


mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}

```