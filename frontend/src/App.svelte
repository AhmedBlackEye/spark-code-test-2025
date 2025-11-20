<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  const API_URL = "http://localhost:8080";

  let todos: TodoItem[] = $state([]);
  let errorMessage = $state("");
  let newTodo: Omit<TodoItem, "id"> = $state({
    title: "",
    description: "",
  });

  async function fetchTodos() {
    errorMessage = "";
    try {
      const response = await fetch(API_URL);
      if (response.status !== 200) {
        errorMessage =
          "We couldn't load your list. Please try refreshing the page.";
        return;
      }

      todos = await response.json();
    } catch (e) {
      errorMessage = "Could not connect to server. Ensure it is running.";
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  async function addTodo(event: Event) {
    event.preventDefault();

    errorMessage = "";

    if (!newTodo.title.trim() || !newTodo.description.trim()) {
      errorMessage = "Please fill in both the title and description.";
      return;
    }

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
      } else {
        errorMessage =
          "Something went wrong adding the item. Please try again.";
      }
    } catch (e) {
      errorMessage = "Could not connect to server. Ensure it is running.";
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  async function deleteTodo(id: number) {
    errorMessage = "";

    try {
      const response = await fetch(`${API_URL}/?id=${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        await fetchTodos();
      } else if (response.status === 404) {
        errorMessage = "Task wasn't found.";
        await fetchTodos();
      } else {
        errorMessage = "Could not delete the item. Please try again.";
      }
    } catch (e) {
      errorMessage = "Could not delete item.";
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

  {#if errorMessage}
    <p class="error-display">{errorMessage}</p>
  {/if}
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
  .error-display {
    color: #ff4d4d;
    margin-top: 15px;
    font-weight: bold;
    font-size: 18px;
    background: rgba(207, 175, 175, 0.2);
    padding: 10px;
    border-radius: 8px;
    display: inline-block;
  }
</style>
