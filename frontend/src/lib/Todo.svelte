<script lang="ts">
  import type { TodoItem } from "./types";

  type Props = TodoItem & {
    onDelete: (id: number) => void;
  };

  const { id, title, description, onDelete }: Props = $props();
</script>

<div class="todo">
  <div class="todo-details">
    <p class="todo-title">{title}</p>
    <p class="todo-description">{description}</p>
  </div>
  <button class="delete-btn" onclick={() => onDelete(id)}>Delete</button>
</div>

<style>
  .todo {
    display: flex;
    justify-content: space-between;
    align-items: center;

    padding: 0px 20px;

    background-color: #303540;

    border-style: solid;
    border-color: #464e61;
    border-width: 1px;
    border-radius: 15px;

    margin: 10px 0px;
  }

  .todo-details {
    display: flex;
    flex-direction: column;
    align-items: start;

    padding: 15px 0px;
  }

  .todo-title {
    font-size: 28px;
    font-weight: bold;

    margin: 0px;
    margin-bottom: 5px;
  }

  .todo-description {
    display: flex;
    margin: 0px;

    font-size: 20px;
    color: #ddd;

    text-align: left;
  }

  .delete-btn {
    background-color: #ff4d4d;
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 8px;
    cursor: pointer;
    font-weight: bold;
    transition: background 0.2s;
  }

  .delete-btn:hover {
    background-color: #cc0000;
  }
</style>
