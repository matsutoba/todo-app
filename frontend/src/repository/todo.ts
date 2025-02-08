import { cookies } from "next/headers";

export interface Todo {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  dueDate: string;
}

export const get = async (): Promise<Todo[]> => {
  const response = await fetch(`${process.env.API_URL}/todos`, {
    headers: {
      Authorization: `Bearer ${(await cookies()).get("token")?.value}`,
    },
    cache: "no-store",
  });
  const res = await response.json();
  return res.data;
};

export const getById = async (id: number): Promise<Todo> => {
  console.log("getById", id);
  const response = await fetch(`${process.env.API_URL}/todos/${id}`, {
    headers: {
      Authorization: `Bearer ${(await cookies()).get("token")?.value}`,
    },
    cache: "no-store",
  });
  const res = await response.json();
  return res.data;
};
