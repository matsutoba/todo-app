"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export interface FormState {
  error: string;
}

export const updateTodo = async (
  id: number,
  state: FormState,
  todo: FormData
) => {
  const newTodo = {
    title: todo.get("title"),
    description: todo.get("description"),
    completed: Boolean(todo.get("completed")),
    dueDate: todo.get("dueDate"),
  };

  const response = await fetch(`${process.env.API_URL}/todos/${id}`, {
    method: "PUT",
    body: JSON.stringify(newTodo),
    headers: {
      Authorization: `Bearer ${(await cookies()).get("token")?.value}`,
    },
  });

  if (!response.ok) {
    return { ...state, error: "タスクの更新に失敗しました。" };
  }

  redirect("/");
};

export const deleteTodo = async (id: number) => {
  const response = await fetch(`${process.env.API_URL}/todos/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${(await cookies()).get("token")?.value}`,
    },
  });

  if (!response.ok) {
    return { error: "タスクの削除に失敗しました。" };
  }

  redirect("/");
};
