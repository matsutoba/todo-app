import { cookies } from "next/headers";

interface Todo {
  id: number;
  title: string;
  description: string;
  completed: boolean;
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
