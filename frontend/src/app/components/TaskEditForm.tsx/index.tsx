"use client";

import { updateTodo } from "@/app/actions/todo";
import { Todo } from "@/repository/todo";
import { useActionState, useState } from "react";

interface TaskEditFormProps {
  todo: Todo;
}

const TaskEditForm = ({ todo }: TaskEditFormProps) => {
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description);
  const [dueDate, setDueDate] = useState(todo.dueDate.split("T")[0]);
  const [completed, setCompleted] = useState(todo.completed);

  const updateTodoWithId = updateTodo.bind(null, todo.id);
  const [state, formAction, isPending] = useActionState(updateTodoWithId, {
    error: "",
  });

  return (
    <div className="flex flex-col justify-center py-20">
      <h2 className="text-center text-2xl font-bold">タスクの編集</h2>
      {state.error && (
        <div className="text-red-500 w-full text-center mt-8">
          {state.error}
        </div>
      )}
      <form className="mx-auto px-16 py-8 w-full " action={formAction}>
        <div className="w-full grid grid-cols-2 gap-4">
          <div className="p-2">タスク名</div>
          <input
            className="p-2 w-full rounded border-2 border-gray-300 rounded-md ring-1 ring-inset border-0 shadow-sm p-2 m-2"
            type="text"
            name="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>
        <div className="w-full grid grid-cols-2 gap-4">
          <div className="p-2">タスク内容</div>
          <input
            className="p-2 w-full rounded border-2 border-gray-300 rounded-md ring-1 ring-inset border-0 shadow-sm p-2 m-2"
            type="text"
            name="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>
        <div className="w-full grid grid-cols-2 gap-4">
          <div className="p-2">締め切り日時</div>
          <input
            className="p-2 w-full rounded border-2 border-gray-300 rounded-md ring-1 ring-inset border-0 shadow-sm p-2 m-2"
            type="date"
            name="dueDate"
            value={dueDate}
            min="2000-01-01"
            max="2099-12-31"
            onChange={(e) => setDueDate(e.target.value)}
          />
        </div>
        <div className="w-full grid grid-cols-2 gap-4">
          <div className="p-2">タスク完了</div>
          <input
            className="p-2 m-2 size-6"
            type="checkbox"
            name="completed"
            checked={completed}
            onChange={(e) => setCompleted(e.target.checked)}
          />
        </div>
        <button
          type="submit"
          className="p-2 m-2 mt-8 w-full bg-blue-500 text-white rounded-md shadow-sm disabled:opacity-50 disabled:cursor-not-allowed"
          disabled={isPending}
        >
          更新
        </button>
      </form>
    </div>
  );
};

export default TaskEditForm;
