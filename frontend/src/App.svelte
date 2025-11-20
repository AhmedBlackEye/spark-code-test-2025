<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  const API_URL = "http://localhost:8080";

  let todos: TodoItem[] = $state([]);
  let newTodo: Omit<TodoItem, "id"> = $state({
    title: "",
    description: "",
  });

  async function fetchTodos() {
    try {
      const response = await fetch(API_URL);
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }

      todos = await response.json();
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  async function addTodo(event: Event) {
    event.preventDefault();

    try {
      const response = await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newTodo),
      });

      if (response.ok) {
        newTodo.title = "";
        newTodo.description = "";
        await fetchTodos();
      }
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  async function deleteTodo(id: number) {
    try {
      const response = await fetch(`${API_URL}/?id=${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        await fetchTodos();
      }
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <div class="todo-list">
    {#each todos as todo}
      <Todo {...todo} onDelete={deleteTodo} />
    {/each}
  </div>

  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form class="todo-list-form" onsubmit={addTodo}>
    <input placeholder="Title" bind:value={newTodo.title} name="title" />
    <input
      placeholder="Description"
      bind:value={newTodo.description}
      name="description"
    />
    <button>Add Todo</button>
  </form>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;

    text-align: center;
    font-size: 24px;

    min-height: 100vh;
    padding: 20px;
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .todo-list {
    margin: 50px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 100px;
  }

  .todo-list-form {
    margin-top: 10px;
  }
</style>
